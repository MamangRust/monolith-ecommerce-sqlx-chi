package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/convert"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	"github.com/MamangRust/monolith-ecommerce-shared/errors/product_errors"
	"github.com/jmoiron/sqlx"
)

const getProducts = `SELECT
    product_id,
    merchant_id,
    category_id,
    name,
    description,
    price,
    count_in_stock,
    brand,
    weight,
    rating,
    slug_product,
    image_product,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM products as p
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR p.name ILIKE '%' || $1 || '%'
        OR p.description ILIKE '%' || $1 || '%'
        OR p.brand ILIKE '%' || $1 || '%'
        OR p.slug_product ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getProductsActive = `SELECT
    product_id,
    merchant_id,
    category_id,
    name,
    description,
    price,
    count_in_stock,
    brand,
    weight,
    rating,
    slug_product,
    image_product,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM products as p
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR p.name ILIKE '%' || $1 || '%'
        OR p.description ILIKE '%' || $1 || '%'
        OR p.brand ILIKE '%' || $1 || '%'
        OR p.slug_product ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getProductsTrashed = `SELECT
    product_id,
    merchant_id,
    category_id,
    name,
    description,
    price,
    count_in_stock,
    brand,
    weight,
    rating,
    slug_product,
    image_product,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM products as p
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR p.name ILIKE '%' || $1 || '%'
        OR p.description ILIKE '%' || $1 || '%'
        OR p.brand ILIKE '%' || $1 || '%'
        OR p.slug_product ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getProductsByMerchant = `WITH
    filtered_products AS (
        SELECT
            p.product_id,
            p.merchant_id,
            p.category_id,
            p.weight,
            p.rating,
            p.slug_product,
            p.name,
            p.description,
            p.price,
            p.count_in_stock,
            p.brand,
            p.image_product,
            p.created_at,
            p.updated_at,
            c.name AS category_name
        FROM products p
            JOIN categories c ON p.category_id = c.category_id
        WHERE
            p.deleted_at IS NULL
            AND p.merchant_id = $1
            AND (
                p.name ILIKE '%' || COALESCE($2, '') || '%'
                OR p.description ILIKE '%' || COALESCE($2, '') || '%'
                OR $2 IS NULL
            )
            AND (
                c.category_id = NULLIF($3, 0)
                OR NULLIF($3, 0) IS NULL
            )
            AND (
                p.price >= COALESCE(NULLIF($4, 0), 0)
                AND p.price <= COALESCE(NULLIF($5, 0), 999999999)
            )
    )
SELECT (
        SELECT COUNT(*)
        FROM filtered_products
    ) AS total_count, fp.product_id, fp.merchant_id, fp.category_id, fp.weight, fp.rating, fp.slug_product, fp.name, fp.description, fp.price, fp.count_in_stock, fp.brand, fp.image_product, fp.created_at, fp.updated_at, fp.category_name
FROM filtered_products fp
ORDER BY fp.created_at DESC
LIMIT $6
OFFSET
    $7`

const getProductsByCategoryName = `WITH
    filtered_products AS (
        SELECT
            p.product_id,
            p.merchant_id,
            p.category_id,
            p.weight,
            p.rating,
            p.slug_product,
            p.name,
            p.description,
            p.price,
            p.count_in_stock,
            p.brand,
            p.image_product,
            p.created_at,
            p.updated_at,
            c.name AS category_name
        FROM products p
            JOIN categories c ON p.category_id = c.category_id
        WHERE
            p.deleted_at IS NULL
            AND c.name = $1
            AND (
                $2::TEXT IS NULL
                OR p.name ILIKE '%' || $2::TEXT || '%'
                OR p.description ILIKE '%' || $2::TEXT || '%'
            )
            AND (
                (
                    $3::INTEGER IS NULL
                    OR p.price >= $3::INTEGER
                )
                AND (
                    $4::INTEGER IS NULL
                    OR p.price <= $4::INTEGER
                )
            )
    )
SELECT (
        SELECT COUNT(*)
        FROM filtered_products
    ) AS total_count, fp.product_id, fp.merchant_id, fp.category_id, fp.weight, fp.rating, fp.slug_product, fp.name, fp.description, fp.price, fp.count_in_stock, fp.brand, fp.image_product, fp.created_at, fp.updated_at, fp.category_name
FROM filtered_products fp
ORDER BY fp.created_at DESC
LIMIT $5
OFFSET
    $6`

const getProductByID = `SELECT
    product_id,
    merchant_id,
    category_id,
    name,
    description,
    price,
    count_in_stock,
    brand,
    weight,
    rating,
    slug_product,
    image_product,
    created_at,
    updated_at
FROM products
WHERE
    product_id = $1
    AND deleted_at IS NULL`

const getProductByIdTrashed = `SELECT product_id, merchant_id, category_id, name, description, price, count_in_stock, brand, weight, rating, slug_product, image_product, created_at, updated_at, deleted_at FROM products WHERE product_id = $1`

type productQueryRepository struct {
	db *sqlx.DB
}

func NewProductQueryRepository(db *sqlx.DB) *productQueryRepository {
	return &productQueryRepository{
		db: db,
	}
}

func (r *productQueryRepository) FindAll(ctx context.Context, req *requests.FindAllProduct) ([]*db.GetProductsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var products []*db.GetProductsRow
	err := r.db.SelectContext(ctx, &products, getProducts,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, product_errors.ErrFindAllProducts.WithInternal(err)
	}

	return products, nil
}

func (r *productQueryRepository) FindActive(ctx context.Context, req *requests.FindAllProduct) ([]*db.GetProductsActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var products []*db.GetProductsActiveRow
	err := r.db.SelectContext(ctx, &products, getProductsActive,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, product_errors.ErrFindActiveProducts.WithInternal(err)
	}

	return products, nil
}

func (r *productQueryRepository) FindTrashed(ctx context.Context, req *requests.FindAllProduct) ([]*db.GetProductsTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var products []*db.GetProductsTrashedRow
	err := r.db.SelectContext(ctx, &products, getProductsTrashed,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, product_errors.ErrFindTrashedProducts.WithInternal(err)
	}

	return products, nil
}

func (r *productQueryRepository) FindByMerchant(ctx context.Context, req *requests.FindAllProductByMerchant) ([]*db.GetProductsByMerchantRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var products []*db.GetProductsByMerchantRow
	err := r.db.SelectContext(ctx, &products, getProductsByMerchant,
		int32(req.MerchantID),
		convert.StringPtr(req.Search),
		int32(req.CategoryID),
		int32(IntPtrToInt(req.MinPrice)),
		int32(IntPtrToInt(req.MaxPrice)),
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, product_errors.ErrFindProductsByMerchant.WithInternal(err)
	}

	return products, nil
}

func (r *productQueryRepository) FindByCategory(ctx context.Context, req *requests.FindAllProductByCategory) ([]*db.GetProductsByCategoryNameRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var products []*db.GetProductsByCategoryNameRow
	err := r.db.SelectContext(ctx, &products, getProductsByCategoryName,
		req.CategoryName,
		req.Search,
		int32(IntPtrToInt(req.MinPrice)),
		int32(IntPtrToInt(req.MaxPrice)),
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, product_errors.ErrFindProductsByCategory.WithInternal(err)
	}

	return products, nil
}

func (r *productQueryRepository) FindByID(ctx context.Context, product_id int) (*db.GetProductByIDRow, error) {
	var product db.GetProductByIDRow
	err := r.db.GetContext(ctx, &product, getProductByID, int32(product_id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, product_errors.ErrProductNotFound.WithInternal(err)
		}
		return nil, product_errors.ErrProductInternal.WithInternal(err)
	}

	return &product, nil
}

func (r *productQueryRepository) FindByIDTrashed(ctx context.Context, product_id int) (*db.Product, error) {
	var product db.Product
	err := r.db.GetContext(ctx, &product, getProductByIdTrashed, int32(product_id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, product_errors.ErrProductNotFound.WithInternal(err)
		}
		return nil, product_errors.ErrProductInternal.WithInternal(err)
	}

	return &product, nil
}
