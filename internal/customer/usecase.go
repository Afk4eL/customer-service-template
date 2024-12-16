package customer

import (
	"context"
	"customer-service/internal/domain/dto"

	"github.com/google/uuid"
)

type CustomerUseCase interface {
	GetCustomerProfile(
		ctx context.Context,
		customerId uuid.UUID) (customerProfile dto.CustomerDto, err error)
	UpdateCustomerProfile(
		ctx context.Context,
		customerId uuid.UUID,
		customerDto dto.CustomerDto) (customerProfile dto.CustomerDto, err error)
}
