package handler

import (
	pbreview_detail "github.com/MamangRust/monolith-ecommerce-pb/review_detail"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/MamangRust/monolith-ecommerce-review-detail/service"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

type Handler struct {
	ReviewDetail        ReviewDetailHandleGrpc
	ReviewDetailQuery   pbreview_detail.ReviewDetailQueryServiceServer
	ReviewDetailCommand pbreview_detail.ReviewDetailCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		ReviewDetailQuery:   NewReviewDetailQueryHandler(deps.Service.ReviewDetailQuery, deps.Logger),
		ReviewDetailCommand: NewReviewDetailCommandHandler(deps.Service.ReviewDetailCommand, deps.Logger),
	}
}
