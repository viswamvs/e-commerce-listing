package product

import (
	"e-commerce-listing/daos/product"
	"e-commerce-listing/database/models"
	"e-commerce-listing/dtos"
	"e-commerce-listing/utils/context"
	"log"
	"time"

	"go.uber.org/zap"
)

type IProductService interface {
	GetProduct(ctx *context.Context, id string) (*dtos.Product, error)
	GetProducts(ctx *context.Context) (*dtos.Products, error)
	SaveProducts(ctx *context.Context, ms *dtos.Product) error
	DeleteProduct(ctx *context.Context) error
}

type ProductService struct {
	db product.IProduct
}

func NewProductService() IProductService {
	return &ProductService{
		db: product.NewProduct(),
	}
}

func (t *ProductService) GetProduct(ctx *context.Context, id string) (*dtos.Product, error) {

	val, err := t.db.Get(ctx, id)
	if err != nil {
		log.Println("unable to fetch product by id", zap.Error(err), zap.Any("id", id))
		return nil, err
	}

	product := convertModelToDto(val)

	return product, nil
}

func (t *ProductService) GetProducts(ctx *context.Context) (*dtos.Products, error) {

	products := &dtos.Products{}

	vals, err := t.db.GetAll(ctx)
	if err != nil {
		log.Println("unable to fetch all products", zap.Error(err))
		return nil, err
	}

	result := []*dtos.Product{}

	for _, v := range vals {
		product := convertModelToDto(v)
		result = append(result, product)
	}

	products.Products = result

	return products, nil
}

func (t *ProductService) SaveProducts(ctx *context.Context, ms *dtos.Product) error {

	if ctx.Param("id") != "" {
		ms.UpdatedAt = time.Now().UTC()
	} else {
		ms.CreatedAt = time.Now().UTC()
	}

	val := convertDtoToModel(ms)

	err := t.db.Upsert(ctx, val)
	if err != nil {
		log.Println("unable to save product", zap.Error(err), zap.Any("id", val.Id))
		return err
	}

	return nil
}

func (t *ProductService) DeleteProduct(ctx *context.Context) error {

	err := t.db.Delete(ctx, ctx.Param("id"))
	if err != nil {
		log.Println("unable to delete product", zap.Error(err), zap.Any("id", ctx.Param("id")))
		return err
	}
	return nil

}

func convertModelToDto(p *models.Product) *dtos.Product {
	return &dtos.Product{
		Id:               p.Id,
		Name:             p.Name,
		Description:      p.Description,
		Price:            p.Price,
		Quantity:         p.Quantity,
		Rating:           p.Rating,
		Review:           p.Review,
		TotalNoOfRatings: p.TotalNoOfRatings,
		TotalNoOfReviews: p.TotalNoOfReviews,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
		DeletedAt:        p.DeletedAt,
	}
}

func convertDtoToModel(p *dtos.Product) *models.Product {

	return &models.Product{
		Id:               p.Id,
		Name:             p.Name,
		Description:      p.Description,
		Price:            p.Price,
		Quantity:         p.Quantity,
		Rating:           p.Rating,
		Review:           p.Review,
		TotalNoOfRatings: p.TotalNoOfRatings,
		TotalNoOfReviews: p.TotalNoOfReviews,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
		DeletedAt:        p.DeletedAt,
	}
}
