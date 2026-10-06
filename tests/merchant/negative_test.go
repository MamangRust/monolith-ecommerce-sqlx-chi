package merchant_test

import (
	"context"

	pbmerchant "github.com/MamangRust/monolith-ecommerce-pb/merchant"
	"github.com/MamangRust/monolith-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent merchant must map to codes.NotFound (404), not Internal.
func (s *MerchantGapiTestSuite) TestMerchantGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pbmerchant.FindByIdMerchantRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent merchant must be NotFound, got %v: %s", st.Code(), st.Message())
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *MerchantRepositoryTestSuite) TestMerchantFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.queryRepo.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
