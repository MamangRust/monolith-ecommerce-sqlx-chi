package repository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

const getOrderStockReservationsByOrder = `SELECT reservation_id, order_id, product_id, quantity, status, created_at, updated_at
FROM order_stock_reservations
WHERE order_id = $1
ORDER BY product_id`

const upsertOrderStockReservation = `INSERT INTO order_stock_reservations (order_id, product_id, quantity, status)
VALUES ($1, $2, $3, 'reserved')
ON CONFLICT (order_id, product_id)
DO UPDATE SET
    quantity = order_stock_reservations.quantity + EXCLUDED.quantity,
    status = 'reserved',
    updated_at = CURRENT_TIMESTAMP
RETURNING reservation_id, order_id, product_id, quantity, status, created_at, updated_at`

const updateOrderStockReservationQuantity = `UPDATE order_stock_reservations
SET quantity = $3, status = 'reserved', updated_at = CURRENT_TIMESTAMP
WHERE order_id = $1 AND product_id = $2
RETURNING reservation_id, order_id, product_id, quantity, status, created_at, updated_at`

const markOrderStockReservationReleased = `UPDATE order_stock_reservations
SET status = 'released', updated_at = CURRENT_TIMESTAMP
WHERE order_id = $1 AND product_id = $2 AND status = 'reserved'
RETURNING reservation_id, order_id, product_id, quantity, status, created_at, updated_at`

const markOrderStockReservationReserved = `UPDATE order_stock_reservations
SET status = 'reserved', updated_at = CURRENT_TIMESTAMP
WHERE order_id = $1 AND product_id = $2 AND status = 'released'
RETURNING reservation_id, order_id, product_id, quantity, status, created_at, updated_at`

const getReservedStockReservationsForTrashedOrders = `SELECT r.reservation_id, r.order_id, r.product_id, r.quantity, r.status, r.created_at, r.updated_at
FROM order_stock_reservations r
JOIN orders o ON o.order_id = r.order_id
WHERE o.deleted_at IS NOT NULL AND r.status = 'reserved'
ORDER BY r.order_id, r.product_id`

const getReleasedStockReservationsForTrashedOrders = `SELECT r.reservation_id, r.order_id, r.product_id, r.quantity, r.status, r.created_at, r.updated_at
FROM order_stock_reservations r
JOIN orders o ON o.order_id = r.order_id
WHERE o.deleted_at IS NOT NULL AND r.status = 'released'
ORDER BY r.order_id, r.product_id`

const getReleasedReservationsForActiveOrders = `SELECT r.reservation_id, r.order_id, r.product_id, r.quantity, r.status, r.created_at, r.updated_at
FROM order_stock_reservations r
JOIN orders o ON o.order_id = r.order_id
WHERE o.deleted_at IS NULL AND r.status = 'released'
ORDER BY r.order_id, r.product_id`

const deleteOrderStockReservations = `DELETE FROM order_stock_reservations WHERE order_id = $1`

const deleteOrderStockReservation = `DELETE FROM order_stock_reservations WHERE order_id = $1 AND product_id = $2`

const deleteAllOrderStockReservations = `DELETE FROM order_stock_reservations
WHERE order_id IN (SELECT order_id FROM orders WHERE deleted_at IS NOT NULL)`

const deleteOldReleasedReservations = `DELETE FROM order_stock_reservations r
USING orders o
WHERE r.order_id = o.order_id
  AND r.status = 'released'
  AND o.deleted_at IS NOT NULL
  AND r.updated_at < $1`

const deleteOldProductStockAdjustments = `DELETE FROM product_stock_adjustments
WHERE created_at < $1`

type StockReservationRepository interface {
	GetByOrder(ctx context.Context, orderID int) ([]*db.OrderStockReservation, error)
	Upsert(ctx context.Context, orderID, productID, quantity int) (*db.OrderStockReservation, error)
	UpdateQuantity(ctx context.Context, orderID, productID, quantity int) (*db.OrderStockReservation, error)
	Release(ctx context.Context, orderID, productID int) (*db.OrderStockReservation, error)
	Reserve(ctx context.Context, orderID, productID int) (*db.OrderStockReservation, error)
	GetReservedForTrashedOrders(ctx context.Context) ([]*db.OrderStockReservation, error)
	GetReleasedForTrashedOrders(ctx context.Context) ([]*db.OrderStockReservation, error)
	DeleteByOrder(ctx context.Context, orderID int) error
	DeleteByOrderProduct(ctx context.Context, orderID, productID int) error
	DeleteAllForTrashedOrders(ctx context.Context) error

	// GetReleasedForActiveOrders returns reservations marked released while their
	// order is still active — a drift pattern repaired by durable reconciliation.
	GetReleasedForActiveOrders(ctx context.Context) ([]*db.OrderStockReservation, error)

	// DeleteOldReleasedReservations removes released reservations for trashed orders
	// whose updated_at predates the cutoff. It returns the number of rows removed.
	DeleteOldReleasedReservations(ctx context.Context, cutoff time.Time) (int64, error)

	// DeleteOldProductStockAdjustments purges idempotency ledger rows older than the
	// cutoff. It returns the number of rows removed.
	DeleteOldProductStockAdjustments(ctx context.Context, cutoff time.Time) (int64, error)
}

type stockReservationRepository struct {
	db *sqlx.DB
}

func NewStockReservationRepository(db *sqlx.DB) StockReservationRepository {
	return &stockReservationRepository{db: db}
}

func (r *stockReservationRepository) GetByOrder(ctx context.Context, orderID int) ([]*db.OrderStockReservation, error) {
	var reservations []*db.OrderStockReservation
	err := r.db.SelectContext(ctx, &reservations, getOrderStockReservationsByOrder, int32(orderID))
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *stockReservationRepository) Upsert(ctx context.Context, orderID, productID, quantity int) (*db.OrderStockReservation, error) {
	var reservation db.OrderStockReservation
	err := r.db.GetContext(ctx, &reservation, upsertOrderStockReservation,
		int32(orderID),
		int32(productID),
		int32(quantity),
	)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *stockReservationRepository) UpdateQuantity(ctx context.Context, orderID, productID, quantity int) (*db.OrderStockReservation, error) {
	var reservation db.OrderStockReservation
	err := r.db.GetContext(ctx, &reservation, updateOrderStockReservationQuantity,
		int32(orderID),
		int32(productID),
		int32(quantity),
	)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *stockReservationRepository) Release(ctx context.Context, orderID, productID int) (*db.OrderStockReservation, error) {
	var reservation db.OrderStockReservation
	err := r.db.GetContext(ctx, &reservation, markOrderStockReservationReleased,
		int32(orderID),
		int32(productID),
	)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *stockReservationRepository) Reserve(ctx context.Context, orderID, productID int) (*db.OrderStockReservation, error) {
	var reservation db.OrderStockReservation
	err := r.db.GetContext(ctx, &reservation, markOrderStockReservationReserved,
		int32(orderID),
		int32(productID),
	)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *stockReservationRepository) GetReservedForTrashedOrders(ctx context.Context) ([]*db.OrderStockReservation, error) {
	var reservations []*db.OrderStockReservation
	err := r.db.SelectContext(ctx, &reservations, getReservedStockReservationsForTrashedOrders)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *stockReservationRepository) GetReleasedForTrashedOrders(ctx context.Context) ([]*db.OrderStockReservation, error) {
	var reservations []*db.OrderStockReservation
	err := r.db.SelectContext(ctx, &reservations, getReleasedStockReservationsForTrashedOrders)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *stockReservationRepository) DeleteByOrder(ctx context.Context, orderID int) error {
	if _, err := r.db.ExecContext(ctx, deleteOrderStockReservations, int32(orderID)); err != nil {
		return err
	}
	return nil
}

func (r *stockReservationRepository) DeleteByOrderProduct(ctx context.Context, orderID, productID int) error {
	if _, err := r.db.ExecContext(ctx, deleteOrderStockReservation,
		int32(orderID),
		int32(productID),
	); err != nil {
		return err
	}
	return nil
}

func (r *stockReservationRepository) DeleteAllForTrashedOrders(ctx context.Context) error {
	if _, err := r.db.ExecContext(ctx, deleteAllOrderStockReservations); err != nil {
		return err
	}
	return nil
}

func (r *stockReservationRepository) GetReleasedForActiveOrders(ctx context.Context) ([]*db.OrderStockReservation, error) {
	var reservations []*db.OrderStockReservation
	err := r.db.SelectContext(ctx, &reservations, getReleasedReservationsForActiveOrders)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *stockReservationRepository) DeleteOldReleasedReservations(ctx context.Context, cutoff time.Time) (int64, error) {
	res, err := r.db.ExecContext(ctx, deleteOldReleasedReservations, pgtype.Timestamp{Time: cutoff, Valid: true})
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (r *stockReservationRepository) DeleteOldProductStockAdjustments(ctx context.Context, cutoff time.Time) (int64, error) {
	res, err := r.db.ExecContext(ctx, deleteOldProductStockAdjustments, pgtype.Timestamp{Time: cutoff, Valid: true})
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
