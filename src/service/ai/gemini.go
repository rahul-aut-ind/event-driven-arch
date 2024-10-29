package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"hello-world-lambda/src/service/internal/config"
)

const (
	geminiAPIURL = "https://generativelanguage.googleapis.com/v1beta/"
	modelsPath   = "models/gemini-1.5-flash:generateContent"
)

type (
	GeminiHandler interface {
		GetRecommendationsFromGemini(encodedImage string) (*RecommendationResp, error)
	}

	GeminiAI struct {
		apiKey string
	}

	UsageMetadata struct {
		PromptTokenCount     int `json:"promptTokenCount"`
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
	}

	SafetyRating struct {
		Category    string `json:"category"`
		Probability string `json:"probability"`
	}

	ResponsePart struct {
		Text string `json:"text"`
	}

	ResponseContent struct {
		Parts []ResponsePart `json:"parts"`
		Role  string         `json:"role"`
	}

	Candidate struct {
		Content       ResponseContent `json:"content"`
		FinishReason  string          `json:"finishReason"`
		Index         int             `json:"index"`
		SafetyRatings []SafetyRating  `json:"safetyRatings"`
	}

	GeminiResponse struct {
		Candidates    []Candidate   `json:"candidates"`
		UsageMetadata UsageMetadata `json:"usageMetadata"`
	}

	InlineData struct {
		MimeType string `json:"mime_type,omitempty"`
		Data     string `json:"data,omitempty"`
	}

	RequestPart struct {
		Text       *string     `json:"text,omitempty"`
		InlineData *InlineData `json:"inline_data,omitempty"`
	}

	RequestContent struct {
		Parts []RequestPart `json:"parts"`
	}

	GeminiRequest struct {
		Contents []RequestContent `json:"contents"`
	}

	GeminiSuggestion struct {
		Status                             string `json:"status"`
		Food                               string `json:"food"`
		EstimatedSize                      string `json:"estimated_size"`
		NutritionalFactsEstimatePerServing struct {
			Calories      string `json:"calories"`
			Protein       string `json:"protein"`
			Carbohydrates string `json:"carbohydrates"`
			Fats          string `json:"fats"`
			Fiber         string `json:"fiber"`
			Sodium        string `json:"sodium"`
		} `json:"nutritional_facts_estimate_per_serving"`
		Advice struct {
			FitnessGoal          string `json:"fitness_goal"`
			PortionSuggestion    string `json:"portion_suggestion"`
			AlternateIngredients string `json:"alternate_ingredients"`
			LifeHack             string `json:"life_hack"`
		} `json:"advice"`
		HealthScore               int    `json:"health_score"`
		KeyRecommendation         string `json:"key_recommendation"`
		RecommendationExplanation string `json:"recommendation_explanation"`
	}
)

func NewGemini(env *config.Env) *GeminiAI {
	return &GeminiAI{
		apiKey: env.GeminiAPIKey,
	}
}

func (v *GeminiAI) callGeminiAPI(encodedImage string) (*GeminiResponse, error) {

	encodedImage = "/9j/4Q/+RXhpZgAATU0AKgAAAAgABgESAAMAAAABAAEAAAEaAAUAAAABAAAAVgEbAAUAAAABAAAAXgEoAAMAAAABAAIAAAITAAMAAAABAAEAAIdpAAQAAAABAAAAZgAAAAAAAABIAAAAAQAAAEgAAAABAAeQAAAHAAAABDAyMjGRAQAHAAAABAECAwCgAAAHAAAABDAxMDCgAQADAAAAAQABAACgAgAEAAAAAQAAAligAwAEAAAAAQAAAyCkBgADAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA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"

	prompt := `Act as a german speaking dietary coach, use the german 'Du', if you address the user. Identify the food in the image, estimate the size and the range in nutritional facts. Give very specific advice based on what you know about the user with motivating words. Also, provide a health score from 1 to 10 (1 being least healthy, 10 being most healthy),  \
Additionally, provide a concise, high-value actionable recommendation that summarizes all aspects, and an explanation for this recommendation. If you detect non-food in the picture or detect something harmful or poisonous, give a warning in the key_recommendation and set the status flag to exception. Answer in the following JSON format in ENGLISH language: \
    {\
        status: String with exception if it is a non-food or harmful item, ok otherwise,\
        food: String describing the main food items identified in the photo,\
        estimated_size: String with approximate size of the food item,\
        nutritional_facts_estimate_per_serving: {\
            calories: String (estimated range in calories per serving of the estimated size. Also state the size of the serving you assumed, as short as possible),\
            protein: String (estimated range in  grams),\
            carbohydrates: String (estimated range in grams),\
            fats: String (estimated range in grams),\
            fiber: String (estimated range in grams),\
            sodium: String (estimated range in milligrams)\
        },\
        advice: {\
            fitness_goal: String with tailored advice for fitness goals,\
            portion_suggestion: String suggesting appropriate portion size,\
            alternate_ingredients: String with healthier alternative suggestions,\
            life_hack: String with a life hack on how to improve the food\
        },\
        health_score: Number that reflects compliance with the criteria: healthy & fit, good for weightloss, low sugar, no alcohol, no ready-made, highly processed food, balanced plate (1/3 vegetables, 1/3 protein, 1/3 carbs) or light plate (2/3 vegetables, 1/3 protein). Please always give a number here, 0 for not applicable cases,\
        key_recommendation: String with a concise, high-value actionable recommendation. A warning, if status is set to exception,\
        recommendation_explanation: String with a more detailed explanation of the key recommendation\
    }`

	var request GeminiRequest
	request.Contents = []RequestContent{
		{
			Parts: []RequestPart{
				{
					Text: &prompt,
				},
				{
					InlineData: &InlineData{
						MimeType: "image/jpeg",
						Data:     encodedImage,
					},
				},
			},
		},
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request body")
	}

	// Parse the URL
	reqURL, err := url.Parse(geminiAPIURL + modelsPath)
	if err != nil {
		return nil, fmt.Errorf("could not add api key")
	}

	// Add query parameters
	query := reqURL.Query()
	query.Set("key", v.apiKey)
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	// fmt.Printf("GeminiAPI request %+v", req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to Gemini API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status code: %d, response body: %s", resp.StatusCode, string(bodyBytes))
	}

	var response GeminiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil

}

func (v *GeminiAI) GetRecommendationsFromGemini(encodedImage string) (*RecommendationResp, error) {
	response, err := v.callGeminiAPI(encodedImage)
	if err != nil {
		return nil, err
	}

	if len(response.Candidates) == 0 {
		return nil, fmt.Errorf("no suggestions from gemini API")
	}

	// fmt.Printf("\n******* Gemini response: %+v", response)

	result := response.Candidates[0].Content.Parts[0].Text
	// fmt.Printf("\nResponse from Gemini %T %+v\n", result, result)

	result = strings.ReplaceAll(result, "```json", "")
	result = strings.ReplaceAll(result, "```", "")

	fmt.Printf("\nCLEANED Response from Gemini %T %+v\n", result, result)

	var geminiRec GeminiSuggestion
	err = json.Unmarshal([]byte(result), &geminiRec)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall the suggestions from response")
	}

	return &RecommendationResp{
		HealthScore:               geminiRec.HealthScore,
		Title:                     geminiRec.Food,
		KeyRecommendation:         geminiRec.KeyRecommendation,
		RecommendationExplanation: geminiRec.RecommendationExplanation,
		FitnessGoal:               geminiRec.Advice.FitnessGoal,
		PortionSuggestion:         geminiRec.Advice.PortionSuggestion,
		AlternateIngredients:      geminiRec.Advice.AlternateIngredients,
		LifeHack:                  geminiRec.Advice.LifeHack,
		EstimatedSize:             geminiRec.EstimatedSize,
		Status:                    geminiRec.Status,
		NutritionalFacts: NutritionalFacts{
			Calories:      geminiRec.NutritionalFactsEstimatePerServing.Calories,
			Protein:       geminiRec.NutritionalFactsEstimatePerServing.Protein,
			Carbohydrates: geminiRec.NutritionalFactsEstimatePerServing.Carbohydrates,
			Fats:          geminiRec.NutritionalFactsEstimatePerServing.Fats,
			Fiber:         geminiRec.NutritionalFactsEstimatePerServing.Fiber,
			Sodium:        geminiRec.NutritionalFactsEstimatePerServing.Sodium,
		},
	}, nil

}
