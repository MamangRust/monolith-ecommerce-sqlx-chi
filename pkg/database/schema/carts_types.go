// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: carts.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateCartParams struct {
	UserID    int32  `db:"user_id" json:"user_id"`
	ProductID int32  `db:"product_id" json:"product_id"`
	Name      string `db:"name" json:"name"`
	Price     int32  `db:"price" json:"price"`
	Image     string `db:"image" json:"image"`
	Quantity  int32  `db:"quantity" json:"quantity"`
	Weight    int32  `db:"weight" json:"weight"`
}

type CreateCartRow struct {
	CartID    int32            `db:"cart_id" json:"cart_id"`
	UserID    int32            `db:"user_id" json:"user_id"`
	ProductID int32            `db:"product_id" json:"product_id"`
	Name      string           `db:"name" json:"name"`
	Price     int32            `db:"price" json:"price"`
	Image     string           `db:"image" json:"image"`
	Quantity  int32            `db:"quantity" json:"quantity"`
	Weight    int32            `db:"weight" json:"weight"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type DeleteAllCartByUserIdParams struct {
	Column1 []int32 `db:"column1" json:"column_1"`
	UserID  int32   `db:"user_id" json:"user_id"`
}

type DeleteCartByIdAndUserIdParams struct {
	CartID int32 `db:"cart_id" json:"cart_id"`
	UserID int32 `db:"user_id" json:"user_id"`
}

type GetCartsParams struct {
	UserID  int32  `db:"user_id" json:"user_id"`
	Column2 string `db:"column2" json:"column_2"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetCartsRow struct {
	CartID     int32            `db:"cart_id" json:"cart_id"`
	UserID     int32            `db:"user_id" json:"user_id"`
	ProductID  int32            `db:"product_id" json:"product_id"`
	Name       string           `db:"name" json:"name"`
	Price      int32            `db:"price" json:"price"`
	Image      string           `db:"image" json:"image"`
	Quantity   int32            `db:"quantity" json:"quantity"`
	Weight     int32            `db:"weight" json:"weight"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}
