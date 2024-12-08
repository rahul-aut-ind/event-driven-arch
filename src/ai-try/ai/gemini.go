package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hello-world-lambda/src/ai-try/internal/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	geminiAPIURL = "https://generativelanguage.googleapis.com/v1beta/"
	modelsPath   = "models/gemini-1.5-flash:generateContent"
)

type (
	GeminiHandler interface {
		GetRecommendationsFromGemini(encodedImage string) (*AnalysisResp, error)
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

func (v *GeminiAI) GetRecommendationsFromGemini(encodedImage string) (*AnalysisResp, error) {
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

	return &AnalysisResp{
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
