package models

type CategoryPrimarKey struct {
	Id string `json:"category_id"`
}

type CreateCategory struct {
	CategoryName string `json:"category_name"`
	ParentId     string `json:"parent_id"`
}

type Category struct {
	Id           string `json:"category_id"`
	ParentId     string `json:"parent_id"`
	CategoryName string `json:"category_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type UpdateCategory struct {
	Id           string `json:"category_id"`
	ParentId     string `json:"parent_id"`
	CategoryName string `json:"category_name"`
}

type GetListCategoryRequest struct {
	Limit  int32
	Offset int32
}

type GetListCategoryResponse struct {
	Count      int32       `json:"count"`
	Categories []*Category `json:"categories"`
}
