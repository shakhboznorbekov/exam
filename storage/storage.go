package storage

import (
	"context"

	"crud/models"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Product() ProductRepoI
	Order() OrdersRepoI
}

type ProductRepoI interface {
	Create(ctx context.Context, req *models.CreateProduct) (string, error)
	GetByPKey(ctx context.Context, req *models.ProductPrimarKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimarKey) error
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategory) (string, error)
	GetByPKey(ctx context.Context, req *models.CategoryPrimarKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.Category, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.CategoryPrimarKey) error
}

type OrdersRepoI interface {
	Create(ctx context.Context, req *models.CreateOrders) (string, error)
	GetByPKey(ctx context.Context, req *models.OrdersPrimarKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrdersRequest) (*models.Ords, error)
	Update(ctx context.Context, req *models.UpdateOrders) (int64, error)
	Delete(ctx context.Context, req *models.OrdersPrimarKey) error
}
