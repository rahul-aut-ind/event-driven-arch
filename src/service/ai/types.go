package ai

import "time"

type (
	ImageProcessor struct {
		IsDeleted      bool          `json:"isDeleted" validate:"required"`
		UserID         string        `json:"userId" validate:"required"`
		PhotoID        string        `json:"photoId" validate:"required"`
		Path           string        `json:"path" validate:"required"`
		ContextName    *string       `json:"contextName,omitempty"`
		ContextID      *string       `json:"contextId,omitempty"`
		OpenApiMessage *AnalysisResp `json:"openApiMessage,omitempty"`
		GeminiMessage  *AnalysisResp `json:"geminiMessage,omitempty"`
		TakenAt        time.Time     `json:"takenAt" validate:"required"`
		UploadedAt     time.Time     `json:"uploadedAt" validate:"required"`
		UpdatedAt      time.Time     `json:"updatedAt" validate:"required"`
		EncodedImage   string        `json:"encodedImage" validate:"required"`
	}

	ImageProcessorResp struct {
		ID          string `json:"id"`
		FullSizeURL string `json:"fullSizeUrl"`
		Metadata
	}

	Metadata struct {
		TakenAt time.Time        `json:"takenAt" validate:"required"`
		Context *MetadataContext `json:"context,omitempty"`
	}

	MetadataContext struct {
		ID   *string `json:"id,omitempty"`
		Name string  `json:"name" validate:"required"`
	}

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
