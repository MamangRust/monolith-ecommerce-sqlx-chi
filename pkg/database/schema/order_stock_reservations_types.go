// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: order_stock_reservations.sql

package db

type DeleteOrderStockReservationParams struct {
	OrderID   int32 `db:"order_id" json:"order_id"`
	ProductID int32 `db:"product_id" json:"product_id"`
}

type MarkOrderStockReservationReleasedParams struct {
	OrderID   int32 `db:"order_id" json:"order_id"`
	ProductID int32 `db:"product_id" json:"product_id"`
}

type MarkOrderStockReservationReservedParams struct {
	OrderID   int32 `db:"order_id" json:"order_id"`
	ProductID int32 `db:"product_id" json:"product_id"`
}

type UpdateOrderStockReservationQuantityParams struct {
	OrderID   int32 `db:"order_id" json:"order_id"`
	ProductID int32 `db:"product_id" json:"product_id"`
	Quantity  int32 `db:"quantity" json:"quantity"`
}

type UpsertOrderStockReservationParams struct {
	OrderID   int32 `db:"order_id" json:"order_id"`
	ProductID int32 `db:"product_id" json:"product_id"`
	Quantity  int32 `db:"quantity" json:"quantity"`
}
