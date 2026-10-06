package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	"github.com/MamangRust/monolith-ecommerce-shared/errors/slider_errors"
	"github.com/jmoiron/sqlx"
)

const createSlider = `INSERT INTO
    sliders (name, image)
VALUES ($1, $2)
RETURNING
    slider_id,
    name,
    image,
    created_at,
    updated_at`

const updateSlider = `UPDATE sliders
SET
    name = $2,
    image = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE
    slider_id = $1
    AND deleted_at IS NULL
RETURNING
    slider_id,
    name,
    image,
    created_at,
    updated_at`

const trashSlider = `UPDATE sliders
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    slider_id = $1
    AND deleted_at IS NULL
RETURNING
    slider_id,
    name,
    image,
    created_at,
    updated_at,
    deleted_at`

const restoreSlider = `UPDATE sliders
SET
    deleted_at = NULL
WHERE
    slider_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    slider_id,
    name,
    image,
    created_at,
    updated_at,
    deleted_at`

const deleteSliderPermanently = `DELETE FROM sliders WHERE slider_id = $1 AND deleted_at IS NOT NULL`

const restoreAllSliders = `UPDATE sliders SET deleted_at = NULL WHERE deleted_at IS NOT NULL`

const deleteAllPermanentSliders = `DELETE FROM sliders WHERE deleted_at IS NOT NULL`

type sliderCommandRepository struct {
	db *sqlx.DB
}

func NewSliderCommandRepository(db *sqlx.DB) *sliderCommandRepository {
	return &sliderCommandRepository{
		db: db,
	}
}

func (r *sliderCommandRepository) Create(ctx context.Context, request *requests.CreateSliderRequest) (*db.CreateSliderRow, error) {
	var slider db.CreateSliderRow
	err := r.db.GetContext(ctx, &slider, createSlider,
		request.Nama,
		request.FilePath,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, slider_errors.ErrSliderNotFound
		}
		return nil, slider_errors.ErrCreateSlider
	}

	return &slider, nil
}

func (r *sliderCommandRepository) Update(ctx context.Context, request *requests.UpdateSliderRequest) (*db.UpdateSliderRow, error) {
	var slider db.UpdateSliderRow
	err := r.db.GetContext(ctx, &slider, updateSlider,
		int32(*request.ID),
		request.Nama,
		request.FilePath,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, slider_errors.ErrSliderNotFound
		}
		return nil, slider_errors.ErrUpdateSlider
	}

	return &slider, nil
}

func (r *sliderCommandRepository) Trash(ctx context.Context, slider_id int) (*db.Slider, error) {
	var slider db.Slider
	err := r.db.GetContext(ctx, &slider, trashSlider, int32(slider_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, slider_errors.ErrSliderNotFound
		}
		return nil, slider_errors.ErrTrashSlider
	}

	return &slider, nil
}

func (r *sliderCommandRepository) Restore(ctx context.Context, slider_id int) (*db.Slider, error) {
	var slider db.Slider
	err := r.db.GetContext(ctx, &slider, restoreSlider, int32(slider_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, slider_errors.ErrSliderNotFound
		}
		return nil, slider_errors.ErrRestoreSlider
	}

	return &slider, nil
}

func (r *sliderCommandRepository) DeletePermanent(ctx context.Context, slider_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteSliderPermanently, int32(slider_id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, slider_errors.ErrSliderNotFound
		}
		return false, slider_errors.ErrDeletePermanentSlider
	}

	return true, nil
}

func (r *sliderCommandRepository) RestoreAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllSliders); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, slider_errors.ErrSliderNotFound
		}
		return false, slider_errors.ErrRestoreAllSlider
	}

	return true, nil
}

func (r *sliderCommandRepository) DeleteAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentSliders); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, slider_errors.ErrSliderNotFound
		}
		return false, slider_errors.ErrDeleteAllPermanentSlider
	}

	return true, nil
}
