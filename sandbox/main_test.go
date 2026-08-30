package sandbox

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

type MockPublisher struct {
	t *testing.T
}

func (m *MockPublisher) Publish(event Event) {
	m.t.Logf("Mock SSE Event Fired | Phase: %s | Message: %s\n", event.Phase, event.Message)
}

func TestEndToEndSandboxAndRCA(t *testing.T) {
	// Skip this test in normal CI unless we explicitly want to run it, 
	// because it requires Docker and a Gemini API key
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("Skipping test because GEMINI_API_KEY is not set")
	}

	ctx := context.Background()
	dm, err := NewDockerManager()
	if err != nil {
		t.Fatalf("Failed to create Docker manager: %v", err)
	}

	// 1. Run the sandbox on dummy_repo
	repoPath, err := filepath.Abs("dummy_repo")
	if err != nil {
		t.Fatalf("Failed to get abs path: %v", err)
	}
	
	publisher := &MockPublisher{t: t}
	ticketID := "test-ticket-123"

	report, err := dm.RunDummyRepo(ctx, ticketID, publisher, repoPath)
	if err != nil {
		t.Fatalf("Sandbox execution failed: %v", err)
	}

	// It should fail with an exit code > 0 because of the intentional bug
	if report.ExitCode == 0 {
		t.Errorf("Expected non-zero exit code due to intentional bug, got 0")
	}

	// 2. Load the source code to send to Gemini
	codeBytes, err := os.ReadFile(filepath.Join(repoPath, "calculator.py"))
	if err != nil {
		t.Fatalf("Failed to read calculator.py: %v", err)
	}

	// 3. Run RCA
	ai, err := NewAIAnalyzer()
	if err != nil {
		t.Fatalf("Failed to create AI analyzer: %v", err)
	}

	rca, err := ai.AnalyzeCrash(ctx, ticketID, publisher, report, string(codeBytes), "The calculator app crashed when trying to divide by zero.")
	if err != nil {
		t.Fatalf("RCA analysis failed: %v", err)
	}

	if rca.RootCause == "" {
		t.Errorf("Expected RootCause in response, got empty")
	}
	t.Logf("RCA Output: %+v", rca)
}
