package db

import (
	"context"
	"fmt"

	"github.com/pgvector/pgvector-go"
)

// similarityThreshold is the max cosine distance for two tickets to be
// considered duplicates. Lower = stricter match. Tune this during testing.
const similarityThreshold = 0.15

// DuplicateMatch describes an existing ticket found within the
// similarity threshold of a newly submitted embedding.
type DuplicateMatch struct {
	ID       string
	Title    string
	Distance float64
}

// FindDuplicate searches for an existing ticket whose embedding is within
// similarityThreshold of the given embedding, using pgvector's cosine
// distance operator (<=>). Returns nil, nil if no close match is found.
func FindDuplicate(ctx context.Context, embedding []float32) (*DuplicateMatch, error) {
	vec := pgvector.NewVector(embedding)

	row := Pool.QueryRow(ctx, `
		SELECT id, title, embedding <=> $1 AS distance
		FROM tickets
		WHERE embedding IS NOT NULL
		ORDER BY distance ASC
		LIMIT 1
	`, vec)

	var match DuplicateMatch
	if err := row.Scan(&match.ID, &match.Title, &match.Distance); err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, fmt.Errorf("dedup query failed: %w", err)
	}

	if match.Distance > similarityThreshold {
		return nil, nil
	}

	return &match, nil
}
