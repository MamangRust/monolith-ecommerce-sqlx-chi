package seeder

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

const seedCreateReview = `INSERT INTO
    reviews (
        user_id,
        product_id,
        name,
        comment,
        rating
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING
    review_id,
    user_id,
    product_id,
    name,
    comment,
    rating,
    created_at,
    updated_at`

const seedCreateReviewDetail = `INSERT INTO
    review_details (review_id, type, url, caption)
VALUES ($1, $2, $3, $4)
RETURNING
    review_detail_id,
    review_id,
    type,
    url,
    caption,
    created_at,
    updated_at`

type reviewSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewReviewSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *reviewSeeder {
	return &reviewSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *reviewSeeder) Seed() error {
	reviews := []db.CreateReviewParams{
		{UserID: 1, ProductID: 1, Name: "John", Comment: "Produk bagus!", Rating: 5},
		{UserID: 2, ProductID: 2, Name: "Anna", Comment: "Sangat puas dengan kualitasnya.", Rating: 4},
		{UserID: 3, ProductID: 3, Name: "Budi", Comment: "Cukup oke untuk harga segini.", Rating: 3},
		{UserID: 4, ProductID: 4, Name: "Siti", Comment: "Pengiriman cepat dan aman.", Rating: 4},
		{UserID: 5, ProductID: 5, Name: "Rina", Comment: "Tidak sesuai ekspektasi.", Rating: 2},
		{UserID: 6, ProductID: 6, Name: "Agus", Comment: "Top, pasti beli lagi!", Rating: 5},
		{UserID: 7, ProductID: 7, Name: "Dian", Comment: "Cocok untuk hadiah.", Rating: 4},
		{UserID: 8, ProductID: 8, Name: "Made", Comment: "Kualitas standar saja.", Rating: 3},
	}

	for _, review := range reviews {
		var createdReview db.CreateReviewRow
		if err := r.db.GetContext(r.ctx, &createdReview, seedCreateReview,
			review.UserID,
			review.ProductID,
			review.Name,
			review.Comment,
			review.Rating,
		); err != nil {
			r.logger.Error("failed to create review", zap.Error(err))
			return err
		}

		var createdDetail db.CreateReviewDetailRow
		if err := r.db.GetContext(r.ctx, &createdDetail, seedCreateReviewDetail,
			createdReview.ReviewID,
			"photo",
			"https://example.com/review_"+review.Name+".jpg",
			toStringPtr("Foto review oleh "+review.Name),
		); err != nil {
			r.logger.Error("failed to create review detail", zap.Error(err))
			return err
		}
	}

	r.logger.Info("review & review detail successfully seeded")

	return nil
}
