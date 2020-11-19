package nlp

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

type NLPResp struct {
	Text   string `json:"text"`
	Sentiment float32 `json:"sentiment"`
	Magnitude float32 `json:"magnitude"`
}

func GetSentiment(c *fiber.Ctx) error {
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}

	// Sets the text to analyze.
	text := c.Query("text")
	if text == "" {
		// no text dont query
		return nil
	}
	fmt.Printf(text)

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	Score := sentiment.DocumentSentiment.Score
	Magnitude := sentiment.DocumentSentiment.Magnitude
	if sentiment.DocumentSentiment.Score >= 0 {
		fmt.Println("Sentiment: positive")
	} else {
		fmt.Println("Sentiment: negative")
	}
	// See https://cloud.google.com/natural-language/docs/reference/rest/v1/Sentiment
	return c.JSON(NLPResp{Text: text, Sentiment: Score, Magnitude: Magnitude})
}
