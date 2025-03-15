package product

import (
	"e-commerce-listing/database/models"
	"e-commerce-listing/utils/context"
	"log"

	"go.uber.org/zap"
)

type IProduct interface {
	Upsert(ctx *context.Context, m ...*models.Product) error
	Get(ctx *context.Context, id string) (*models.Product, error)
	GetFor(ctx *context.Context, ids []string, name string) ([]*models.Product, error)
	GetAll(ctx *context.Context) ([]*models.Product, error)
	Delete(ctx *context.Context, id string) error
}

type Product struct {
}

func NewProduct() IProduct {
	return &Product{}
}

func (t *Product) getTable() string {
	return "products"
}

func (t *Product) Upsert(ctx *context.Context, m ...*models.Product) error {
	return ctx.DB.Debug().WithContext(ctx.Request.Context()).Table(t.getTable()).Save(m).Error
}

func (t *Product) Get(ctx *context.Context, id string) (*models.Product, error) {
	var result models.Product

	err := ctx.DB.Debug().WithContext(ctx.Request.Context()).Table(t.getTable()).First(&result, "id = ?", id).Error
	if err != nil {
		log.Println("unable to product details", zap.Error(err))
		return nil, err
	}

	return &result, err
}

func (t *Product) GetFor(ctx *context.Context, ids []string, name string) ([]*models.Product, error) {
	var result []*models.Product

	tx := ctx.DB.Debug().WithContext(ctx.Request.Context()).Table(t.getTable())

	if name != "" {
		tx.Where("name ilike ?", "%"+name+"%")
	}

	if len(ids) > 0 {
		tx.Where("id IN ?", ids)
	}

	err := tx.Find(result).Error
	if err != nil {
		log.Println("unable to fetch products", zap.Error(err))
	}

	return result, nil
}

func (t *Product) GetAll(ctx *context.Context) ([]*models.Product, error) {
	var result []*models.Product

	err := ctx.DB.Debug().WithContext(ctx.Request.Context()).Table(t.getTable()).Find(&result).Error
	if err != nil {
		log.Println("unable to fetch the products", zap.Error(err))
		return nil, err
	}

	return result, err
}

func (t *Product) Delete(ctx *context.Context, id string) error {
	var result models.Product

	err := ctx.DB.WithContext(ctx.Request.Context()).Table(t.getTable()).Delete(&result, "id = ?", id).Error
	if err != nil {
		log.Println("unable to delete product", zap.Error(err))
		return err
	}

	return err
}
