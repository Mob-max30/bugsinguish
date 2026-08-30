package models

import "time"

// TicketStatus represents the lifecycle stage of a ticket.
type TicketStatus string

const (
	StatusNew             TicketStatus = "new"
	StatusTriaging        TicketStatus = "triaging"
	StatusSandboxRunning  TicketStatus = "sandbox_running"
	StatusDiagnosed       TicketStatus = "diagnosed"
	StatusResolved        TicketStatus = "resolved"
)

// Diagnosis holds the AI root-cause-analysis output produced after the
// sandbox reproduces the bug (populated by Ujwal's RCA pipeline).
type Diagnosis struct {
	RootCause   string `json:"root_cause"`
	Explanation string `json:"explanation"`
	File        string `json:"file"`
}

// Ticket is the core record for a single bug report.
type Ticket struct {
	ID            string       `json:"id"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	StackTrace    string       `json:"stack_trace"`
	RepoBranchURL string       `json:"repo_branch_url"`
	Severity      string       `json:"severity"`
	Status        TicketStatus `json:"status"`
	Embedding     []float32    `json:"embedding,omitempty"`
	Diagnosis     *Diagnosis   `json:"diagnosis,omitempty"`
	Diff          string       `json:"diff,omitempty"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}
