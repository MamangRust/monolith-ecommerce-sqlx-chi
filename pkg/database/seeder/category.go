package seeder

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

const seedCreateCategory = `INSERT INTO
    categories (
        name,
        description,
        slug_category,
        image_category
    )
VALUES ($1, $2, $3, $4)
RETURNING
    category_id,
    name,
    description,
    slug_category,
    image_category,
    created_at,
    updated_at`

type categorySeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewCategorySeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *categorySeeder {
	return &categorySeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *categorySeeder) Seed() error {
	categories := []db.CreateCategoryParams{
		{
			Name:          "Elektronik",
			Description:   toStringPtr("Produk elektronik seperti smartphone, laptop, dan aksesori elektronik lainnya."),
			SlugCategory:  toStringPtr("elektronik"),
			ImageCategory: toStringPtr("elektronik.jpg"),
		},
		{
			Name:          "Kesehatan & Kecantikan",
			Description:   toStringPtr("Produk perawatan tubuh, skincare, dan suplemen kesehatan."),
			SlugCategory:  toStringPtr("kesehatan-kecantikan"),
			ImageCategory: toStringPtr("kesehatan.jpg"),
		},
		{
			Name:          "Peralatan Rumah Tangga",
			Description:   toStringPtr("Peralatan dapur, perlengkapan rumah, dan furnitur."),
			SlugCategory:  toStringPtr("peralatan-rumah"),
			ImageCategory: toStringPtr("rumah.jpg"),
		},
		{
			Name:          "Ibu & Bayi",
			Description:   toStringPtr("Produk khusus untuk ibu hamil, menyusui, dan bayi."),
			SlugCategory:  toStringPtr("ibu-bayi"),
			ImageCategory: toStringPtr("ibu-bayi.jpg"),
		},
		{
			Name:          "Olahraga & Outdoor",
			Description:   toStringPtr("Perlengkapan olahraga, fitness, dan kegiatan luar ruangan."),
			SlugCategory:  toStringPtr("olahraga-outdoor"),
			ImageCategory: toStringPtr("olahraga.jpg"),
		},
		{
			Name:          "Makanan & Minuman",
			Description:   toStringPtr("Makanan ringan, minuman, bahan makanan segar dan kemasan."),
			SlugCategory:  toStringPtr("makanan-minuman"),
			ImageCategory: toStringPtr("makanan.jpg"),
		},
		{
			Name:          "Gaming & Console",
			Description:   toStringPtr("Konsol game, aksesori, dan game terbaru dari berbagai platform."),
			SlugCategory:  toStringPtr("gaming-console"),
			ImageCategory: toStringPtr("gaming.jpg"),
		},
		{
			Name:          "Perlengkapan Otomotif",
			Description:   toStringPtr("Aksesori mobil dan motor, oli, serta sparepart kendaraan."),
			SlugCategory:  toStringPtr("otomotif"),
			ImageCategory: toStringPtr("otomotif.jpg"),
		},
	}

	for _, category := range categories {
		var created db.CreateCategoryRow
		if err := r.db.GetContext(r.ctx, &created, seedCreateCategory,
			category.Name,
			category.Description,
			category.SlugCategory,
			category.ImageCategory,
		); err != nil {
			r.logger.Error("Failed to insert category", zap.Error(err))
			return err
		}
	}

	r.logger.Info("Successfully seeded 10 categories")
	return nil
}
