package catalog

import "context"

var(
	ErrNotFound = errors.New(:Entity not found)
)

type Repository interface {
	Close()
	PutProduct(ctx context.Context, p Product)
	GetProductByID(ctx context.Context, id string)(*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64)([]Product, error)
	ListProductsWithIDs(ctx context.Context, ids []string)([]product,error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64)([]product, error)
}

type elasticRepository struct {
	client *elastic.Client
}

type productDocument struct{
	Name string `json:"name"`
	Description string `json:"descriptiom"`
	Price string `json:"price"`
}