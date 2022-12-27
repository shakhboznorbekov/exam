package models

import "gorm.io/plugin/soft_delete"

type OrdersPrimarKey struct {
	Id string `json:"orders_id"`
}

type CreateOrders struct {
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
}

type Orders struct {
	Id          string                `json:"orders_id"`
	Description string                `json:"description"`
	ProductId   string                `json:"product_id"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `json:"gorm:"softDelete:milli""`
}

type UpdateOrders struct {
	Id          string `json:"category_id"`
	Description string `json:"description"`
	ProductId   string `json:"product_id"`
}

type GetListOrdersRequest struct {
	Limit  int32
	Offset int32
}

type GetListOrdersResponse struct {
	Count int32     `json:"count"`
	Order []*Orders `json:"orderss"`
}

type Ords struct {
	Count int32    `json:"count"`
	Order []*Order `json:"orders"`
}

type ChildOrder struct {
	OrderId     string      `json:"order_id"`
	Description string      `json:"description"`
	Product     []*Products `json:"product"`
}

type Products struct {
	ProductId string        `json:"product_id"`
	Name      string        `json:"name"`
	Category  []*Categories `json:"category"`
}

type Categories struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"name"`
	ParentId     string `json:"parent_id"`
}

type Order struct {
	OrderId     string       `json:"order_id"`
	Description string       `json:"description"`
	Product     OrderProduct `json:"product"`
}

type OrderProduct struct {
	Id       string          `json:"order_id"`
	Name     string          `json:"description"`
	Category ProductCategory `json:"product"`
}

type ProductCategory struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}
