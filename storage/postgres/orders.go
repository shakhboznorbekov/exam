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

type OrdersRepo struct {
	db *pgxpool.Pool
}

func NewOrdersRepo(db *pgxpool.Pool) *OrdersRepo {
	return &OrdersRepo{
		db: db,
	}
}

func (f *OrdersRepo) Create(ctx context.Context, orders *models.CreateOrders) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO orders(
			orders_id,
			description, 
			poduct_id, 
			updated_at
		) VALUES ( $1, $2 , $3, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		orders.Description,
		orders.ProductId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *OrdersRepo) GetByPKey(ctx context.Context, pkey *models.OrdersPrimarKey) (*models.Order, error) {

	var (
		id          sql.NullString
		description sql.NullString
		productId   sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
		respCg      = models.ProductCategory{}
		respProduct = models.OrderProduct{}
	)

	query := `
		SELECT
			orders_id,
			description,
			product_id, 
			created_at,
			updated_at
		FROM
			orders
		WHERE orders_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&description,
			&productId,
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	var categoryId string
	err = f.db.QueryRow(ctx,
		"select product_id, product_name, category_id from products where product_id = $1",
		productId).Scan(&respProduct.Id, &respProduct.Name, &categoryId)
	if err != nil {
		return nil, err
	}

	err = f.db.QueryRow(ctx,
		"select category_id, category_name, parent_id from category where category_id = $1",
		categoryId).Scan(&respCg.Id, &respCg.Name, &respCg.ParentId)
	if err != nil {
		return nil, err
	}
	respProduct.Category = respCg

	return &models.Order{
		OrderId:     id.String,
		Description: description.String,
		Product:     respProduct,
	}, nil
}

func (f *OrdersRepo) GetList(ctx context.Context, req *models.GetListOrdersRequest) (*models.Ords, error) {

	var (
		respMain    = models.Ords{}
		respCg      = models.ProductCategory{}
		respProduct = models.OrderProduct{}
		offset      = " OFFSET 0"
		limit       = " LIMIT 20"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
	SELECT
	orders_id,
	description,
	product_id, 
	FROM
	orders WHERE deleted_at = 0;
	`
	query += offset + limit

	rows, err := f.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		res := &models.Order{}
		prId := ""
		err := rows.Scan(
			&res.OrderId,
			&res.Description,
			&prId,
		)

		if err != nil {
			return nil, err
		}

		var categoryId string
		err = f.db.QueryRow(ctx,
			"select product_id, product_name, category_id from products where product_id = $1",
			prId).Scan(&respProduct.Id, &respProduct.Name, &categoryId)
		if err != nil {
			return nil, err
		}

		err = f.db.QueryRow(ctx,
			"select category_id, category_name, parent_id from category where category_id = $1",
			categoryId).Scan(&respCg.Id, &respCg.Name, &respCg.ParentId)
		if err != nil {
			return nil, err
		}
		respProduct.Category = respCg
		res.Product = respProduct
		respMain.Order = append(respMain.Order, res)
	}
	return &respMain, err
}

func (f *OrdersRepo) Update(ctx context.Context, req *models.UpdateOrders) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			description = :description,
			product_id = :product_id, 
			updated_at = now()
		WHERE orders_id = :orders_id
	`

	params = map[string]interface{}{
		"orders_id":   req.Id,
		"description": req.Description,
		"product_id":  req.ProductId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *OrdersRepo) Delete(ctx context.Context, req *models.OrdersPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM category WHERE orders_id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
