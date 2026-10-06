package seeder

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

const seedCreateSlider = `INSERT INTO
    sliders (name, image)
VALUES ($1, $2)
RETURNING
    slider_id,
    name,
    image,
    created_at,
    updated_at`

type sliderSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewSliderSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *sliderSeeder {
	return &sliderSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *sliderSeeder) Seed() error {
	sliders := []db.CreateSliderParams{
		{Name: "Promo Akhir Tahun", Image: "slider1.jpg"},
		{Name: "Diskon Elektronik", Image: "slider2.jpg"},
		{Name: "Flash Sale Mingguan", Image: "slider3.jpg"},
		{Name: "Produk Terbaru", Image: "slider4.jpg"},
		{Name: "Promo Kesehatan", Image: "slider5.jpg"},
		{Name: "Belanja Hemat", Image: "slider6.jpg"},
		{Name: "Gaming Gear Diskon", Image: "slider7.jpg"},
		{Name: "Gratis Ongkir", Image: "slider8.jpg"},
		{Name: "Ramadhan Sale", Image: "slider9.jpg"},
		{Name: "Perlengkapan Bayi", Image: "slider10.jpg"},
	}

	for _, slider := range sliders {
		var created db.CreateSliderRow
		if err := r.db.GetContext(r.ctx, &created, seedCreateSlider,
			slider.Name,
			slider.Image,
		); err != nil {
			r.logger.Error("failed to seed slider", zap.Error(err))
			return err
		}
	}

	r.logger.Info("slider successfully seeded")

	return nil
}
