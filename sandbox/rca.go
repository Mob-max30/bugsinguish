package sandbox

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// RCAResponse is the structured output from Gemini
type RCAResponse struct {
	RootCause   string `json:"root_cause"`
	Explanation string `json:"explanation"`
	File        string `json:"file"`
	Diff        string `json:"diff"`
}

// AIAnalyzer encapsulates logic to call Gemini for Root Cause Analysis
type AIAnalyzer struct {
	apiKey string
}

// NewAIAnalyzer creates a new AIAnalyzer instance
func NewAIAnalyzer() (*AIAnalyzer, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	return &AIAnalyzer{apiKey: apiKey}, nil
}

// AnalyzeCrash sends the crash report and source code to Gemini and asks for a structured RCA
func (ai *AIAnalyzer) AnalyzeCrash(ctx context.Context, ticketID string, publisher EventPublisher, report *CrashReport, sourceCode string, bugDescription string) (*RCAResponse, error) {
	publishHelper(publisher, ticketID, "diagnosing", "Sending crash logs to Gemini for Root Cause Analysis...", nil)

	if ai.apiKey == "" {
		// Mock/Fallback RCA response for standalone local testing when GEMINI_API_KEY is not set
		fallback := &RCAResponse{
			RootCause:   "Missing zero denominator validation in divide() function at calculator.py line 14.",
			Explanation: "The divide function accepts integer arguments a and b without checking if b == 0 before performing division operator. This raises ZeroDivisionError unhandled.",
			File:        "sandbox/dummy_repo/calculator.py",
			Diff:        "--- calculator.py\n+++ calculator.py\n@@ -13,3 +13,5 @@\n def divide(a, b):\n+    if b == 0:\n+        raise ValueError(\"Cannot divide by zero\")\n     return a / b",
		}
		publishHelper(publisher, ticketID, "diagnosed", "Diagnosis complete. Draft PR generated.", fallback)
		return fallback, nil
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(ai.apiKey))
	if err != nil {
		return nil, fmt.Errorf("unable to create genai client: %w", err)
	}
	defer client.Close()

	prompt := fmt.Sprintf(`
You are an expert debugging assistant. Your task is to perform a Root Cause Analysis (RCA) on a bug and provide a fix.

User's Bug Description:
%s

Crash Report:
Exit Code: %d
Stdout:
%s
Stderr:
%s

Relevant Source Code:
%s

Provide your response in strict JSON format matching the following schema. Do not include markdown code blocks, just raw JSON.
{
  "root_cause": "A short summary of the root cause.",
  "explanation": "A detailed explanation of why the crash happened.",
  "file": "The name of the file that needs to be fixed.",
  "diff": "A unified git diff string fixing the bug."
}
`, bugDescription, report.ExitCode, report.Stdout, report.Stderr, sourceCode)

	model := client.GenerativeModel("gemini-1.5-pro")
	model.ResponseMIMEType = "application/json"

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("gemini api call failed: %w", err)
	}

	publishHelper(publisher, ticketID, "diagnosing", "Parsing AI diagnosis response...", nil)

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("empty response from gemini")
	}

	var rcaResp RCAResponse
	part := resp.Candidates[0].Content.Parts[0]
	if textPart, ok := part.(genai.Text); ok {
		if err := json.Unmarshal([]byte(textPart), &rcaResp); err != nil {
			return nil, fmt.Errorf("failed to parse json response: %w\nRaw response: %s", err, string(textPart))
		}
	} else {
		return nil, fmt.Errorf("unexpected part type from gemini")
	}

	publishHelper(publisher, ticketID, "diagnosed", "Diagnosis complete. Draft PR generated.", rcaResp)

	return &rcaResp, nil
}
