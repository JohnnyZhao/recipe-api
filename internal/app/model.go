package app

import (
	"time"

	"github.com/johnnyzhao/retail-ai-api/internal/domain"
)

type ListResponse struct {
	Recipes []domain.Recipe `json:"recipes"`
}

type ItemResponse struct {
	Message string        `json:"message"`
	Recipe  domain.Recipe `json:"recipe"`
}

type TimestampedItem struct {
	domain.Recipe
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateItemResponse struct {
	Message string          `json:"message"`
	Recipe  TimestampedItem `json:"recipe"`
}

type ErrorResponse struct {
	Message  string  `json:"message"`
	Required *string `json:"required,omitempty"`
}

type RecipePayload struct {
	Title       *string `json:"title,omitempty"`
	MakingTime  *string `json:"making_time,omitempty"`
	Serves      *string `json:"serves,omitempty"`
	Ingredients *string `json:"ingredients,omitempty"`
	Cost        *int    `json:"cost,omitempty"`
}

func (p *RecipePayload) ToUpdateValues() map[string]interface{} {
	result := make(map[string]interface{})
	if p.Title != nil {
		result["title"] = *p.Title
	}
	if p.MakingTime != nil {
		result["making_time"] = *p.MakingTime
	}
	if p.Serves != nil {
		result["serves"] = *p.Serves
	}
	if p.Ingredients != nil {
		result["ingredients"] = *p.Ingredients
	}
	if p.Cost != nil {
		result["cost"] = *p.Cost
	}
	return result
}

func (p *RecipePayload) ValidateRequired() ([]string, bool) {
	result := make([]string, 0)
	if p.Title == nil {
		result = append(result, "title")
	}
	if p.MakingTime == nil {
		result = append(result, "making_time")
	}
	if p.Serves == nil {
		result = append(result, "serves")
	}
	if p.Ingredients == nil {
		result = append(result, "ingredients")
	}
	if p.Cost == nil {
		result = append(result, "cost")
	}
	if len(result) > 0 {
		return result, false
	}
	return result, true
}
