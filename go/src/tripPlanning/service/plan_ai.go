package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tripPlanning/constants"
)

// TravelPlannerService handles travel planning logic
type TravelPlannerService struct {
	openAIKey string
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// NewTravelPlannerService creates a new instance of TravelPlannerService
func NewTravelPlannerService(openAIKey string) *TravelPlannerService {
	return &TravelPlannerService{
		openAIKey: constants.Openai_key,
	}
}

// AiGeneratedPlan generates a plan using ChatGPT API
func (s *TravelPlannerService) AiGeneratedPlan(city, startDay, endDay string) (string, error) {
	// Set up OpenAI API client
	apiEndpoint := constants.OpenaiEndpoint

	fmt.Println("successful connect to gpt")

	// Construct the input prompt for ChatGPT
	promptMessage := fmt.Sprintf("Give me a trip plan to travel to %s from %s to %s with the most famous places of interest. Start with \"Here is AI-generated trip advising:\"", city, startDay, endDay)

	// Construct the request payload
	requestPayload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "You are a helpful assistant capable of generating travel plans.",
			},
			{
				"role":    "user",
				"content": promptMessage,
			},
		},
		"max_tokens": 200,
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return "", fmt.Errorf("error encoding JSON: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set API key header
	req.Header.Set("Authorization", "Bearer "+constants.Openai_key)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal JSON response into OpenAIResponse struct
	var response OpenAIResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Check if response contains content
	if len(response.Choices) > 0 && len(response.Choices[0].Message.Content) > 0 {
		// Return the content of the first choice
		return response.Choices[0].Message.Content, nil
	} else {
		return "", fmt.Errorf("no content found in response, length of response.choices is %d", len(response.Choices))
	}
}
