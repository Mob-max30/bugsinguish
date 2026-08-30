package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool is the shared Postgres connection pool for the app.
var Pool *pgxpool.Pool

// Connect opens a connection pool to the Neon Postgres database at the
// given connection string (e.g. from the DATABASE_URL env var), enables
// the pgvector extension, and runs the tickets table migration.
func Connect(ctx context.Context, connString string) error {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	Pool = pool

	if err := enablePgvector(ctx); err != nil {
		return err
	}

	if err := migrateTickets(ctx); err != nil {
		return err
	}

	log.Println("connected to Neon Postgres, pgvector enabled, tickets table ready")
	return nil
}

// enablePgvector turns on the pgvector extension if it isn't already.
func enablePgvector(ctx context.Context) error {
	_, err := Pool.Exec(ctx, `CREATE EXTENSION IF NOT EXISTS vector;`)
	if err != nil {
		return fmt.Errorf("unable to enable pgvector extension: %w", err)
	}
	return nil
}

// migrateTickets creates the tickets table if it doesn't exist yet.
// Embedding uses a 768-dim vector (adjust to match your embedding model's
// output size). Diagnosis is stored as JSONB since it's a nested object.
func migrateTickets(ctx context.Context) error {
	const schema = `
	CREATE TABLE IF NOT EXISTS tickets (
		id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title            TEXT NOT NULL,
		description      TEXT NOT NULL,
		stack_trace      TEXT,
		repo_branch_url  TEXT,
		severity         TEXT NOT NULL,
		status           TEXT NOT NULL DEFAULT 'new',
		embedding        vector(768),
		diagnosis        JSONB,
		diff             TEXT,
		created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
		updated_at       TIMESTAMPTZ NOT NULL DEFAULT now()
	);
	`
	if _, err := Pool.Exec(ctx, schema); err != nil {
		return fmt.Errorf("unable to migrate tickets table: %w", err)
	}
	return nil
}
