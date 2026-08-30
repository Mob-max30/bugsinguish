package sse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Event is the shape published to the team channel and sent to every
// connected client. "phase" matches the ticket lifecycle / sandbox
// progress steps (e.g. "cloning", "sandbox_running", "diagnosed").
type Event struct {
	TicketID string      `json:"ticket_id"`
	Phase    string      `json:"phase"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Time     time.Time   `json:"time"`
}

// broker holds all currently-connected clients and fans out published
// events to each of them.
type broker struct {
	mu      sync.Mutex
	clients map[chan Event]bool
}

var b = &broker{
	clients: make(map[chan Event]bool),
}

// Publish sends an event to every connected SSE client. Safe to call from
// anywhere in the backend (ticket handlers, sandbox manager, etc.).
func Publish(evt Event) {
	if evt.Time.IsZero() {
		evt.Time = time.Now().UTC()
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	for ch := range b.clients {
		select {
		case ch <- evt:
		default:
			// Client is too slow / buffer full — drop the event for it
			// rather than blocking the publisher.
		}
	}
}

// StreamHandler handles GET /api/stream — upgrades the connection to SSE
// and pushes every published Event to this client until it disconnects.
func StreamHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch := make(chan Event, 16)

	b.mu.Lock()
	b.clients[ch] = true
	b.mu.Unlock()

	defer func() {
		b.mu.Lock()
		delete(b.clients, ch)
		b.mu.Unlock()
		close(ch)
	}()

	for {
		select {
		case evt := <-ch:
			payload, err := json.Marshal(evt)
			if err != nil {
				continue
			}
			fmt.Fprintf(w, "data: %s\n\n", payload)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
