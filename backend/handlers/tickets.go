package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	"bugsinguish-backend/db"
	"bugsinguish-backend/embeddings"
	"bugsinguish-backend/models"
)

// CreateTicket handles POST /tickets — generates an embedding for the new
// report, checks pgvector for a near-duplicate, and either short-circuits
// with the existing ticket (409) or persists a new one.
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var t models.Ticket
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if t.Status == "" {
		t.Status = models.StatusNew
	}

	// Generate an embedding from the title + description and check for a
	// near-duplicate before inserting. If embedding generation fails, we
	// log it via the error return and fall through to a normal create
	// rather than blocking ticket intake on an AI-service hiccup.
	embedding, embErr := embeddings.GenerateEmbedding(r.Context(), t.Title+" "+t.Description)
	if embErr == nil {
		dup, dupErr := db.FindDuplicate(r.Context(), embedding)
		if dupErr == nil && dup != nil {
			writeJSON(w, http.StatusConflict, map[string]interface{}{
				"duplicate_of": dup.ID,
				"title":        dup.Title,
				"distance":     dup.Distance,
			})
			return
		}
		t.Embedding = embedding
	}

	row := db.Pool.QueryRow(r.Context(), `
		INSERT INTO tickets (title, description, stack_trace, repo_branch_url, severity, status, embedding)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`, t.Title, t.Description, t.StackTrace, t.RepoBranchURL, t.Severity, t.Status, embedding)

	if err := row.Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt); err != nil {
		http.Error(w, "failed to create ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

// ListTickets handles GET /tickets — returns all tickets.
func ListTickets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Pool.Query(r.Context(), `
		SELECT id, title, description, stack_trace, repo_branch_url, severity, status, diff, created_at, updated_at
		FROM tickets ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, "failed to list tickets: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tickets := []models.Ticket{}
	for rows.Next() {
		var t models.Ticket
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.StackTrace, &t.RepoBranchURL,
			&t.Severity, &t.Status, &t.Diff, &t.CreatedAt, &t.UpdatedAt); err != nil {
			http.Error(w, "failed to scan ticket: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tickets = append(tickets, t)
	}

	writeJSON(w, http.StatusOK, tickets)
}

// GetTicket handles GET /tickets/:id — returns a single ticket.
func GetTicket(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var t models.Ticket
	row := db.Pool.QueryRow(r.Context(), `
		SELECT id, title, description, stack_trace, repo_branch_url, severity, status, diagnosis, diff, created_at, updated_at
		FROM tickets WHERE id = $1
	`, id)

	if err := row.Scan(&t.ID, &t.Title, &t.Description, &t.StackTrace, &t.RepoBranchURL,
		&t.Severity, &t.Status, &t.Diagnosis, &t.Diff, &t.CreatedAt, &t.UpdatedAt); err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "ticket not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// UpdateTicket handles PATCH /tickets/:id — updates a ticket's status
// (and optionally diagnosis/diff once the RCA pipeline is wired in).
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var body struct {
		Status *models.TicketStatus `json:"status,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if body.Status == nil {
		http.Error(w, "status field is required", http.StatusBadRequest)
		return
	}

	tag, err := db.Pool.Exec(r.Context(), `
		UPDATE tickets SET status = $1, updated_at = $2 WHERE id = $3
	`, *body.Status, time.Now().UTC(), id)

	if err != nil {
		http.Error(w, "failed to update ticket: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if tag.RowsAffected() == 0 {
		http.Error(w, "ticket not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"id": id, "status": string(*body.Status)})
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
