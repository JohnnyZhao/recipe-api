package repo

import (
	"context"

	"github.com/johnnyzhao/retail-ai-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeRepo struct {
	db *gorm.DB
}

func NewRecipeRepo(db *gorm.DB) *RecipeRepo {
	return &RecipeRepo{db: db}
}

func (r *RecipeRepo) Create(ctx context.Context, recipe *domain.Recipe) error {
	return r.db.WithContext(ctx).Create(recipe).Error
}

func (r *RecipeRepo) GetList(ctx context.Context) ([]domain.Recipe, error) {
	var recipes []domain.Recipe

	if err := r.db.WithContext(ctx).Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *RecipeRepo) GetByID(ctx context.Context, id uint) (domain.Recipe, error) {
	var recipe domain.Recipe

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&recipe).Error; err != nil {
		return domain.Recipe{}, err
	}
	return recipe, nil
}

func (r *RecipeRepo) UpdateByID(ctx context.Context, id uint, values map[string]interface{}) (domain.Recipe, error) {
	var recipe domain.Recipe
	if err := r.db.WithContext(ctx).Model(&recipe).
		Where("id = ?", id).
		Updates(values).Error; err != nil {
		return domain.Recipe{}, err
	}
	return recipe, nil
}

func (r *RecipeRepo) DeleteByID(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Recipe{}, id).Error
}

func (r *RecipeRepo) Migrate() error {
	return r.db.AutoMigrate(&domain.Recipe{})
}
