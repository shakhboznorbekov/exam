package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"crud/models"
	"crud/pkg/helper"
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (f *CategoryRepo) Create(ctx context.Context, category *models.CreateCategory) (string, error) {

	var (
		id     = uuid.New().String()
		query  string
		nullId sql.NullString
	)

	query = `
		INSERT INTO category(
			category_id,
			parent_id,
			category_name, 
			updated_at
		) VALUES ( $1, $2 , $3, now())
	`

	if category.ParentId == "" {
		_, err := f.db.Exec(ctx, query,
			id,
			nullId,
			category.CategoryName,
		)

		if err != nil {
			return "", err
		}
	} else {

		_, err := f.db.Exec(ctx, query,
			id,
			category.ParentId,
			category.CategoryName,
		)

		if err != nil {
			return "", err
		}

	}

	return id, nil
}

func (f *CategoryRepo) GetByPKey(ctx context.Context, pkey *models.CategoryPrimarKey) (*models.Category, error) {

	var (
		id           sql.NullString
		parentId     sql.NullString
		categoryName sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	query := `
		SELECT
			category_id,
			parent_id,
			category_name, 
			created_at,
			updated_at
		FROM
			category
		WHERE category_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&parentId,
			&categoryName,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:           id.String,
		ParentId:     parentId.String,
		CategoryName: categoryName.String,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}, nil
}

func (f *CategoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	var (
		resp   = models.GetListCategoryResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 2"
	)

	query := `
		SELECT
			COUNT(*) OVER(),
			category_id,
			parent_id,
			category_name,  
			updated_at
		FROM
			category
		WHERE 
			deleted_at is null	
	`

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id           sql.NullString
			parentId     sql.NullString
			categoryName sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&parentId,
			&categoryName,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &models.Category{
			Id:           id.String,
			ParentId:     parentId.String,
			CategoryName: categoryName.String,
			UpdatedAt:    updatedAt.String,
		})

	}

	return &resp, err
}

func (f *CategoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			category
		SET
			parent_id = :parent_id,
			category_name = :category_name, 
			updated_at = now()
		WHERE category_id = :category_id
	`

	params = map[string]interface{}{
		"category_id":   req.Id,
		"parent_id":     req.ParentId,
		"category_name": req.CategoryName,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *CategoryRepo) Delete(ctx context.Context, req *models.CategoryPrimarKey) error {

	query := `
		UPDATE
			category
		SET
			deleted_at = now()
		WHERE category_id = :$1
	`
	_, err := f.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return err
}
