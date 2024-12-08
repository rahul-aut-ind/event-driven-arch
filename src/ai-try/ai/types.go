package ai

type (
	AnalysisResp struct {
		ID                        string           `json:"id"`
		ImageURL                  string           `json:"fullSizeUrl"`
		EstimatedSize             string           `json:"estimatedSize"`
		HealthScore               int              `json:"healthScore"`
		Title                     string           `json:"title"`
		KeyRecommendation         string           `json:"keyRecommendation"`
		RecommendationExplanation string           `json:"recommendationExplanation"`
		FitnessGoal               string           `json:"fitnessGoal"`
		PortionSuggestion         string           `json:"portionSuggestion"`
		AlternateIngredients      string           `json:"alternateIngredients"`
		LifeHack                  string           `json:"lifeHack"`
		Status                    string           `json:"status"`
		NutritionalFacts          NutritionalFacts `json:"nutritionalFacts"`
	}

	NutritionalFacts struct {
		Calories      string `json:"calories"`
		Protein       string `json:"protein"`
		Carbohydrates string `json:"carbohydrates"`
		Fats          string `json:"fats"`
		Fiber         string `json:"fiber"`
		Sodium        string `json:"sodium"`
	}
)
