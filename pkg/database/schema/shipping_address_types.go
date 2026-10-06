// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: shipping_address.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateShippingAddressParams struct {
	OrderID        int32   `db:"order_id" json:"order_id"`
	Alamat         string  `db:"alamat" json:"alamat"`
	Provinsi       string  `db:"provinsi" json:"provinsi"`
	Negara         string  `db:"negara" json:"negara"`
	Kota           string  `db:"kota" json:"kota"`
	Courier        string  `db:"courier" json:"courier"`
	ShippingMethod string  `db:"shipping_method" json:"shipping_method"`
	ShippingCost   float64 `db:"shipping_cost" json:"shipping_cost"`
}

type CreateShippingAddressRow struct {
	ShippingAddressID int32            `db:"shipping_address_id" json:"shipping_address_id"`
	OrderID           int32            `db:"order_id" json:"order_id"`
	Alamat            string           `db:"alamat" json:"alamat"`
	Provinsi          string           `db:"provinsi" json:"provinsi"`
	Negara            string           `db:"negara" json:"negara"`
	Kota              string           `db:"kota" json:"kota"`
	Courier           string           `db:"courier" json:"courier"`
	ShippingMethod    string           `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64          `db:"shipping_cost" json:"shipping_cost"`
	CreatedAt         pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt         pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetShippingAddressParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetShippingAddressRow struct {
	ShippingAddressID int32            `db:"shipping_address_id" json:"shipping_address_id"`
	OrderID           int32            `db:"order_id" json:"order_id"`
	Alamat            string           `db:"alamat" json:"alamat"`
	Provinsi          string           `db:"provinsi" json:"provinsi"`
	Negara            string           `db:"negara" json:"negara"`
	Kota              string           `db:"kota" json:"kota"`
	Courier           string           `db:"courier" json:"courier"`
	ShippingMethod    string           `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64          `db:"shipping_cost" json:"shipping_cost"`
	CreatedAt         pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt         pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount        int64            `db:"total_count" json:"total_count"`
}

type GetShippingAddressActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetShippingAddressActiveRow struct {
	ShippingAddressID int32            `db:"shipping_address_id" json:"shipping_address_id"`
	OrderID           int32            `db:"order_id" json:"order_id"`
	Alamat            string           `db:"alamat" json:"alamat"`
	Provinsi          string           `db:"provinsi" json:"provinsi"`
	Negara            string           `db:"negara" json:"negara"`
	Kota              string           `db:"kota" json:"kota"`
	Courier           string           `db:"courier" json:"courier"`
	ShippingMethod    string           `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64          `db:"shipping_cost" json:"shipping_cost"`
	CreatedAt         pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt         pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt         pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount        int64            `db:"total_count" json:"total_count"`
}

type GetShippingAddressByOrderIDRow struct {
	ShippingAddressID int32            `db:"shipping_address_id" json:"shipping_address_id"`
	OrderID           int32            `db:"order_id" json:"order_id"`
	Alamat            string           `db:"alamat" json:"alamat"`
	Provinsi          string           `db:"provinsi" json:"provinsi"`
	Negara            string           `db:"negara" json:"negara"`
	Kota              string           `db:"kota" json:"kota"`
	Courier           string           `db:"courier" json:"courier"`
	ShippingMethod    string           `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64          `db:"shipping_cost" json:"shipping_cost"`
	CreatedAt         pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt         pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetShippingAddressTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetShippingAddressTrashedRow struct {
	ShippingAddressID int32            `db:"shipping_address_id" json:"shipping_address_id"`
	OrderID           int32            `db:"order_id" json:"order_id"`
	Alamat            string           `db:"alamat" json:"alamat"`
	Provinsi          string           `db:"provinsi" json:"provinsi"`
	Negara            string           `db:"negara" json:"negara"`
	Kota              string           `db:"kota" json:"kota"`
	Courier           string           `db:"courier" json:"courier"`
	ShippingMethod    string           `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64          `db:"shipping_cost" json:"shipping_cost"`
	CreatedAt         pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt         pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt         pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount        int64            `db:"total_count" json:"total_count"`
}

type GetShippingByIDRow struct {
	ShippingAddressID int32            `db:"shipping_address_id" json:"shipping_address_id"`
	OrderID           int32            `db:"order_id" json:"order_id"`
	Alamat            string           `db:"alamat" json:"alamat"`
	Provinsi          string           `db:"provinsi" json:"provinsi"`
	Negara            string           `db:"negara" json:"negara"`
	Kota              string           `db:"kota" json:"kota"`
	Courier           string           `db:"courier" json:"courier"`
	ShippingMethod    string           `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64          `db:"shipping_cost" json:"shipping_cost"`
	CreatedAt         pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt         pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateShippingAddressParams struct {
	ShippingAddressID int32   `db:"shipping_address_id" json:"shipping_address_id"`
	Alamat            string  `db:"alamat" json:"alamat"`
	Provinsi          string  `db:"provinsi" json:"provinsi"`
	Negara            string  `db:"negara" json:"negara"`
	Kota              string  `db:"kota" json:"kota"`
	Courier           string  `db:"courier" json:"courier"`
	ShippingMethod    string  `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64 `db:"shipping_cost" json:"shipping_cost"`
}

type UpdateShippingAddressRow struct {
	ShippingAddressID int32            `db:"shipping_address_id" json:"shipping_address_id"`
	OrderID           int32            `db:"order_id" json:"order_id"`
	Alamat            string           `db:"alamat" json:"alamat"`
	Provinsi          string           `db:"provinsi" json:"provinsi"`
	Negara            string           `db:"negara" json:"negara"`
	Kota              string           `db:"kota" json:"kota"`
	Courier           string           `db:"courier" json:"courier"`
	ShippingMethod    string           `db:"shipping_method" json:"shipping_method"`
	ShippingCost      float64          `db:"shipping_cost" json:"shipping_cost"`
	CreatedAt         pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt         pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
