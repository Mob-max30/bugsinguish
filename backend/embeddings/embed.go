package embeddings

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

const (
	embeddingModel = "gemini-embedding-001"
	// Must match the vector(768) column in backend/db/neon.go's migration.
	outputDimensionality = int32(768)
)

// GenerateEmbedding calls the Google Gen AI SDK's embedding endpoint for a
// single string and returns the resulting vector as []float32, ready to
// store in the tickets.embedding column.
func GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create genai client: %w", err)
	}

	contents := []*genai.Content{
		{Parts: []*genai.Part{{Text: text}}},
	}

	dim := outputDimensionality
	config := &genai.EmbedContentConfig{
		OutputDimensionality: &dim,
	}

	resp, err := client.Models.EmbedContent(ctx, embeddingModel, contents, config)
	if err != nil {
		return nil, fmt.Errorf("embedding request failed: %w", err)
	}
	if len(resp.Embeddings) == 0 {
		return nil, fmt.Errorf("embedding response contained no embeddings")
	}

	return resp.Embeddings[0].Values, nil
}
