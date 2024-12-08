package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hello-world-lambda/src/ai-try/internal/config"
	"io"
	"net/http"
)

const (
	openAIAPIURL = "https://api.openai.com/v1/chat/completions"
	modelName    = "gpt-4o"
	maxTokens    = 500
	responseType = "json_object"
)

type (
	Message struct {
		Role    string    `json:"role"`
		Content []Content `json:"content"`
	}

	Content struct {
		Type     string    `json:"type"`
		Text     string    `json:"text,omitempty"`
		ImageURL *ImageURL `json:"image_url,omitempty"`
	}

	ImageURL struct {
		URL string `json:"url"`
	}

	Request struct {
		Model          string    `json:"model"`
		Messages       []Message `json:"messages"`
		MaxTokens      int       `json:"max_tokens"`
		Temperature    float64   `json:"temperature"`
		ResponseFormat struct {
			Type string `json:"type"`
		} `json:"response_format"`
	}

	OpenAIResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	VisionSuggestion struct {
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

	OpenAIHandler interface {
		GetRecommendationsFromOpenAI(encodedImage string) (*AnalysisResp, error)
	}

	VisionAI struct {
		apiKey string
	}
)

// New creates a new instance of VisionAI
func NewOpenAI(env *config.Env) *VisionAI {
	return &VisionAI{
		apiKey: env.VisionAPIKey,
	}
}

// nolint:funlen // expected to be long
func (v *VisionAI) callOpenAIAPI(encodedImage string) (*OpenAIResponse, error) {

	prompt := `Act as a german speaking dietary coach, use the german "Du", if you address the user. Identify the food in the image, estimate the size and the range in nutritional facts. Give very specific advice based on what you know about the user with motivating words. Also, provide a health score from 1 to 10 (1 being least healthy, 10 being most healthy),  Additionally, provide a concise, high-value actionable recommendation that summarizes all aspects, and an explanation for this recommendation. If you detect non-food in the picture or detect something harmful or poisonous, give a warning in the key_recommendation and set the is_exception flag to true.

Answer in the following JSON format in ENGLISH language:
    {
        "status": "String with 'exception' if it is a non-food or harmful item, 'ok' otherwise),
        "food": "String describing the main food items identified in the photo",
        "estimated_size": "String with approximate size of the food item",
        "nutritional_facts_estimate_per_serving": {
            "calories": "String (estimated range in calories per serving of the estimated size. Also state the size of the serving you assumed, as short as possible)",
            "protein": "String (estimated range in  grams)",
            "carbohydrates": "String (estimated range in grams)",
            "fats": "String (estimated range in grams)",
            "fiber": "String (estimated range in grams)",
            "sodium": "String (estimated range in milligrams)"
        },
        "advice": {
            "fitness_goal": "String with tailored advice for fitness goals",
            "portion_suggestion": "String suggesting appropriate portion size",
            "alternate_ingredients": "String with healthier alternative suggestions",
            "life_hack": "String with a life hack on how to improve the food"
        },
        "health_score": "Number that reflects compliance with the criteria: healthy & fit, good for weightloss, low sugar, no alcohol, no ready-made, highly processed food, balanced plate (1/3 vegetables, 1/3 protein, 1/3 carbs) or light plate (2/3 vegetables, 1/3 protein). Please always give a number here, 0 for not applicable cases",
        "key_recommendation": "String with a concise, high-value actionable recommendation. A warning, if is_exception is set to true",
        "recommendation_explanation": "String with a more detailed explanation of the key recommendation"
    }
		`

	request := Request{
		Model: modelName,
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: prompt,
					},
					{
						Type: "image_url",
						ImageURL: &ImageURL{
							URL: encodedImage,
						},
					},
				},
			},
		},
		MaxTokens:   maxTokens,
		Temperature: 0.0,
		ResponseFormat: struct {
			Type string `json:"type"`
		}{
			Type: responseType,
		},
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request body")
	}

	req, err := http.NewRequest("POST", openAIAPIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	// fmt.Printf("VisionAPI request %+v", req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", v.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to OpenAI Vision API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status code: %d, response body: %s", resp.StatusCode, string(bodyBytes))
	}

	var response OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil

}

func (v *VisionAI) GetRecommendationsFromOpenAI(encodedImage string) (*AnalysisResp, error) {

	response, err := v.callOpenAIAPI(encodedImage)
	if err != nil {
		return nil, err
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no suggestions from vision API")
	}

	// fmt.Printf("\n******* OpenAI VisonAPI response: %+v", response)

	result := response.Choices[0].Message.Content
	fmt.Printf("\nResponse from OpenAI %T %+v\n", result, result)
	var visionRec VisionSuggestion

	err = json.Unmarshal([]byte(result), &visionRec)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall the suggestions from response")
	}

	return &AnalysisResp{
		HealthScore:               visionRec.HealthScore,
		Title:                     visionRec.Food,
		KeyRecommendation:         visionRec.KeyRecommendation,
		RecommendationExplanation: visionRec.RecommendationExplanation,
		FitnessGoal:               visionRec.Advice.FitnessGoal,
		PortionSuggestion:         visionRec.Advice.PortionSuggestion,
		AlternateIngredients:      visionRec.Advice.AlternateIngredients,
		LifeHack:                  visionRec.Advice.LifeHack,
		EstimatedSize:             visionRec.EstimatedSize,
		Status:                    visionRec.Status,
		NutritionalFacts: NutritionalFacts{
			Calories:      visionRec.NutritionalFactsEstimatePerServing.Calories,
			Protein:       visionRec.NutritionalFactsEstimatePerServing.Protein,
			Carbohydrates: visionRec.NutritionalFactsEstimatePerServing.Carbohydrates,
			Fats:          visionRec.NutritionalFactsEstimatePerServing.Fats,
			Fiber:         visionRec.NutritionalFactsEstimatePerServing.Fiber,
			Sodium:        visionRec.NutritionalFactsEstimatePerServing.Sodium,
		},
	}, nil

}
