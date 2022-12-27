package models

type ProductPrimarKey struct {
	Id string `json:"product_id"`
}

type CreateProduct struct {
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	CategoryId  string `json:"category_id"`
}
type Product struct {
	Id          string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	CategoryId  string `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type UpdateProduct struct {
	Id          string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	CategoryId  string `json:"category_id"`
}

type GetListProductRequest struct {
	Limit      int32
	Offset     int32
	CategoryId string
}

type GetListProductResponse struct {
	Count    int32      `json:"count"`
	Products []*Product `json:"products"`
}
