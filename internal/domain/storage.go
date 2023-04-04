package domain

import "context"

type RecipeStorage interface {
	Create(ctx context.Context, recipe *Recipe) error
	GetList(ctx context.Context) ([]Recipe, error)
	GetByID(ctx context.Context, id uint) (Recipe, error)
	UpdateByID(ctx context.Context, id uint, values map[string]interface{}) (Recipe, error)
	DeleteByID(ctx context.Context, id uint) error
}
