package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MrAndiw/instagram-roasting/app/helper"
)

// ContentPart struct represents a part of the content
type ContentPart struct {
	Text string `json:"text"`
}

// Content struct represents the content object in the request body
type Content struct {
	Role  string        `json:"role"`
	Parts []ContentPart `json:"parts"`
}

// GenerateContentRequest struct represents the request body for generateContent
type GenerateContentRequest struct {
	Contents []Content `json:"contents"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type CitationSource struct {
	StartIndex int    `json:"startIndex"`
	EndIndex   int    `json:"endIndex"`
	URI        string `json:"uri"`
	License    string `json:"license"`
}

type CitationMetadata struct {
	CitationSources []CitationSource `json:"citationSources"`
}

type Candidate struct {
	Content          Content          `json:"content"`
	FinishReason     string           `json:"finishReason"`
	Index            int              `json:"index"`
	SafetyRatings    []SafetyRating   `json:"safetyRatings"`
	CitationMetadata CitationMetadata `json:"citationMetadata"`
}

type Response struct {
	Candidates    []Candidate    `json:"candidates"`
	UsageMetadata map[string]int `json:"usageMetadata"`
}

func GetGemini(text string, config helper.Config) (string, error) {
	// Replace with your actual API key
	apiKey := config.ApiKeyGemini

	// Define the model name
	modelName := "gemini-pro"

	// Build the request URL
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/%s:generateContent", modelName)

	// Create the request body
	requestBody := GenerateContentRequest{
		Contents: []Content{
			{
				Role: "user",
				Parts: []ContentPart{
					{Text: text},
				},
			},
		},
	}

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("[GetGemini] Error marshalling request body: %s", err)
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a POST request with the specified URL, headers, and body
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return "", fmt.Errorf("[GetGemini] Error creating request: %s", err)
	}

	// Set the authorization header with your API key
	req.Header.Set("x-goog-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Do the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("[GetGemini] Error making request: %s", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("[GetGemini] StatusCode: %d Error: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("[GetGemini] Error reading response body: %s", err)
	}

	// Unmarshall the response body to a struct
	var data Response
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return "", fmt.Errorf("[GetGemini] Error unmarshalling response body: %s", err)
	}

	// Access the generated content from the first candidate
	generatedContent := data.Candidates[0].Content.Parts[0].Text

	return generatedContent, nil
}
