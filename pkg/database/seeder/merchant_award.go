package seeder

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

const seedCreateMerchantCertificationOrAward = `INSERT INTO
    merchant_certifications_and_awards (
        merchant_id,
        title,
        description,
        issued_by,
        issue_date,
        expiry_date,
        certificate_url
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    merchant_certification_id,
    merchant_id,
    title,
    description,
    issued_by,
    issue_date,
    expiry_date,
    certificate_url,
    created_at,
    updated_at`

type merchantAwardSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewMerchantAwardSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *merchantAwardSeeder {
	return &merchantAwardSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *merchantAwardSeeder) Seed() error {
	awards := []db.CreateMerchantCertificationOrAwardParams{
		{
			MerchantID:     1,
			Title:          "ISO 9001 Certified",
			Description:    toStringPtr("Manajemen mutu bersertifikat"),
			IssuedBy:       toStringPtr("ISO Organization"),
			IssueDate:      toDate(time.Date(2020, time.January, 15, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC)),
			CertificateUrl: toStringPtr("https://example.com/iso9001-cert.pdf"),
		},
		{
			MerchantID:     2,
			Title:          "Top UMKM 2023",
			Description:    toStringPtr("Penghargaan untuk UMKM terbaik tahun 2023"),
			IssuedBy:       toStringPtr("Kementerian Koperasi"),
			IssueDate:      toDate(time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Time{}),
			CertificateUrl: toStringPtr("https://example.com/umkm-award-2023.pdf"),
		},
		{
			MerchantID:     3,
			Title:          "Halal Certified",
			Description:    toStringPtr("Sertifikasi halal dari MUI"),
			IssuedBy:       toStringPtr("Majelis Ulama Indonesia"),
			IssueDate:      toDate(time.Date(2021, time.March, 12, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Date(2024, time.March, 12, 0, 0, 0, 0, time.UTC)),
			CertificateUrl: toStringPtr("https://example.com/halal-cert.pdf"),
		},
		{
			MerchantID:     4,
			Title:          "Best Food Product 2022",
			Description:    toStringPtr("Penghargaan untuk produk makanan terbaik tahun 2022"),
			IssuedBy:       toStringPtr("Asosiasi Kuliner Indonesia"),
			IssueDate:      toDate(time.Date(2022, time.November, 5, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Time{}),
			CertificateUrl: toStringPtr("https://example.com/best-food-2022.pdf"),
		},
		{
			MerchantID:     5,
			Title:          "Eco-Friendly Business",
			Description:    toStringPtr("Sertifikasi bisnis ramah lingkungan"),
			IssuedBy:       toStringPtr("Green Business Council"),
			IssueDate:      toDate(time.Date(2023, time.April, 22, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Date(2026, time.April, 22, 0, 0, 0, 0, time.UTC)),
			CertificateUrl: toStringPtr("https://example.com/eco-friendly-cert.pdf"),
		},
		{
			MerchantID:     6,
			Title:          "Top Seller 2023",
			Description:    toStringPtr("Penjual terbaik platform e-commerce tahun 2023"),
			IssuedBy:       toStringPtr("Tokopedia"),
			IssueDate:      toDate(time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Time{}),
			CertificateUrl: toStringPtr("https://example.com/top-seller-2023.pdf"),
		},
		{
			MerchantID:     7,
			Title:          "BPOM Certified",
			Description:    toStringPtr("Sertifikasi produk dari Badan Pengawas Obat dan Makanan"),
			IssuedBy:       toStringPtr("Badan POM RI"),
			IssueDate:      toDate(time.Date(2022, time.August, 3, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Date(2025, time.August, 3, 0, 0, 0, 0, time.UTC)),
			CertificateUrl: toStringPtr("https://example.com/bpom-cert.pdf"),
		},
		{
			MerchantID:     8,
			Title:          "Creativepreneur Award",
			Description:    toStringPtr("Penghargaan untuk wirausaha kreatif"),
			IssuedBy:       toStringPtr("Kementerian Pariwisata dan Ekonomi Kreatif"),
			IssueDate:      toDate(time.Date(2023, time.December, 15, 0, 0, 0, 0, time.UTC)),
			ExpiryDate:     toDate(time.Time{}),
			CertificateUrl: toStringPtr("https://example.com/creativepreneur-award.pdf"),
		},
	}

	for _, award := range awards {
		var created db.MerchantCertificationsAndAward
		if err := r.db.GetContext(r.ctx, &created, seedCreateMerchantCertificationOrAward,
			award.MerchantID,
			award.Title,
			award.Description,
			award.IssuedBy,
			award.IssueDate,
			award.ExpiryDate,
			award.CertificateUrl,
		); err != nil {
			r.logger.Error("failed to seed merchant award", zap.Error(err))
			return err
		}
	}

	r.logger.Info("merchant awards seeded successfully")

	return nil
}
