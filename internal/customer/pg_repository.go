package customer

import (
	"context"
	"customer-service/internal/domain/dto"

	"github.com/google/uuid"
)

type CustomerPgRepository interface {
	GetCustomerProfile(
		ctx context.Context,
		customerId uuid.UUID) (customerProfile dto.CustomerDto, err error)
	UpdateCustomerProfile(
		ctx context.Context,
		customerId uuid.UUID,
		customerDto dto.CustomerDto) (customerProfile dto.CustomerDto, err error)
	//добавить все нужные запросы к бд
}
