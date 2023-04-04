package app

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/johnnyzhao/recipe-api/internal/domain"
	"gorm.io/gorm"
)

type Api struct {
	recipeRepo domain.RecipeStorage
}

func NewApi(repo domain.RecipeStorage) *Api {
	return &Api{recipeRepo: repo}
}

func (a *Api) CreateRecipe(c *gin.Context) {
	var payload RecipePayload
	//TODO, validate length of payload fields
	if err := c.BindJSON(&payload); err != nil {
		a.handleBadRequest(c, "invalid payload")
		return
	}
	required, ok := payload.ValidateRequired()
	if !ok {
		a.handleBadRequest(c, "Recipe creation failed!", required...)
		return
	}
	recipe := domain.Recipe{
		Title:       *payload.Title,
		MakingTime:  *payload.MakingTime,
		Serves:      *payload.Serves,
		Ingredients: *payload.Ingredients,
		Cost:        *payload.Cost,
	}
	if err := a.recipeRepo.Create(c, &recipe); err != nil {
		a.handleInternalError(c)
		return
	}

	c.JSON(http.StatusOK, CreateItemResponse{
		Message: "Recipe successfully created!",
		Recipe:  TimestampedItem{Recipe: recipe, CreatedAt: recipe.CreatedAt, UpdatedAt: recipe.UpdatedAt},
	})
}

func (a *Api) ListRecipes(c *gin.Context) {
	recipes, err := a.recipeRepo.GetList(c)
	if err != nil {
		a.handleInternalError(c)
		return
	}
	c.JSON(http.StatusOK, ListResponse{Recipes: recipes})
}

func (a *Api) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		a.handleBadRequest(c, "invalid ID")
		return
	}

	recipe, err := a.recipeRepo.GetByID(c, uint(id))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		a.handleNotFound(c)
		return
	case err != nil:
		a.handleInternalError(c)
		return
	}
	message := "Recipe details by id"
	c.JSON(http.StatusOK, ListResponse{
		Message: &message,
		Recipes: []domain.Recipe{recipe},
	})
}

func (a *Api) PatchByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		a.handleBadRequest(c, "invalid ID")
		return
	}

	//TODO, validate length of payload fields
	var payload RecipePayload
	if err := c.BindJSON(&payload); err != nil {
		a.handleBadRequest(c, "invalid payload")
		return
	}

	recipe, err := a.recipeRepo.GetByID(c, uint(id))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		a.handleNotFound(c)
		return
	case err != nil:
		a.handleInternalError(c)
		return
	}

	recipe, err = a.recipeRepo.UpdateByID(c, uint(id), payload.ToUpdateValues())
	if err != nil {
		a.handleInternalError(c)
		return
	}

	recipe, err = a.recipeRepo.GetByID(c, uint(id))

	if err != nil {
		a.handleInternalError(c)
		return
	}

	c.JSON(http.StatusOK, ItemResponse{
		Message: "Recipe successfully updated!",
		Recipe:  recipe,
	})
}

func (a *Api) DeleteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		a.handleBadRequest(c, "invalid ID")
		return
	}

	_, err = a.recipeRepo.GetByID(c, uint(id))
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		a.handleNotFound(c)
		return
	case err != nil:
		a.handleInternalError(c)
		return
	}

	if err = a.recipeRepo.DeleteByID(c, uint(id)); err != nil {
		a.handleInternalError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe successfully removed!"})
}

func (a *Api) handleInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Message: "Internal Server Error",
	})
}

func (a *Api) handleBadRequest(c *gin.Context, message string, required ...string) {
	if len(required) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: message,
		})
		return
	}
	requiredStr := strings.Join(required, ", ")
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Message:  message,
		Required: &requiredStr,
	})
}

func (a *Api) handleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Message: "No recipe found",
	})
}
