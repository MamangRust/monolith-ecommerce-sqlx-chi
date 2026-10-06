package seeder

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

const seedCreateOrder = `INSERT INTO
    orders (
        merchant_id,
        user_id,
        total_price
    )
VALUES ($1, $2, $3)
RETURNING
    order_id,
    user_id,
    merchant_id,
    total_price,
    created_at,
    updated_at`

const seedCreateOrderItem = `INSERT INTO
    order_items (
        order_id,
        product_id,
        quantity,
        price
    )
VALUES ($1, $2, $3, $4)
RETURNING
    order_item_id,
    order_id,
    product_id,
    quantity,
    price,
    created_at,
    updated_at`

type orderSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewOrderSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *orderSeeder {
	return &orderSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *orderSeeder) Seed() error {
	for i := 1; i <= 8; i++ {
		var order db.CreateOrderRow
		if err := r.db.GetContext(r.ctx, &order, seedCreateOrder,
			int32(i),
			int32(i),
			int32(10000*i),
		); err != nil {
			r.logger.Error("failed to create order", zap.Error(err))
			return err
		}

		var orderItem db.CreateOrderItemRow
		if err := r.db.GetContext(r.ctx, &orderItem, seedCreateOrderItem,
			order.OrderID,
			int32(i),
			int32(i),
			int32(10000),
		); err != nil {
			r.logger.Error("failed to create order item", zap.Error(err))
			return err
		}
	}

	r.logger.Info("order & order-item successfully seeded")

	return nil
}
