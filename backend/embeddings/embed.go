package embeddings

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const embeddingModel = "text-embedding-004"

// GenerateEmbedding calls the Google Gen AI SDK's embedding endpoint for a
// single string and returns the resulting vector as []float32 (768 dimensions),
// ready to store in the tickets.embedding column.
func GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		// Return a zero-vector if GEMINI_API_KEY is unset so local dev and testing
		// can run without requiring a live Gemini API key.
		return make([]float32, 768), nil
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("unable to create genai client: %w", err)
	}
	defer client.Close()

	em := client.EmbeddingModel(embeddingModel)
	res, err := em.EmbedContent(ctx, genai.Text(text))
	if err != nil {
		return nil, fmt.Errorf("embedding request failed: %w", err)
	}
	if res.Embedding == nil || len(res.Embedding.Values) == 0 {
		return nil, fmt.Errorf("embedding response contained no values")
	}

	return res.Embedding.Values, nil
}
