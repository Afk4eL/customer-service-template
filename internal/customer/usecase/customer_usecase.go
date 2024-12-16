package usecase

import (
	"context"
	"customer-service/internal/customer"
	"customer-service/internal/domain/dto"

	"github.com/google/uuid"
)

type customerUseCase struct {
	customerPgRepo customer.CustomerPgRepository
}

func NewCustomerUsecase(customerPgRepo customer.CustomerPgRepository) *customerUseCase {
	return &customerUseCase{customerPgRepo: customerPgRepo}
}

func (u *customerUseCase) GetCustomerProfile(
	ctx context.Context,
	customerId uuid.UUID) (customerProfile dto.CustomerDto, err error) {
	const op = "customer.usecase.GetCustomerProfile"

	customerProfile, err = u.customerPgRepo.GetCustomerProfile(ctx, customerId)
	if err != nil {
		return dto.CustomerDto{}, err
	}

	return customerProfile, nil
}

func (u *customerUseCase) UpdateCustomerProfile(
	ctx context.Context,
	customerId uuid.UUID,
	customerDto dto.CustomerDto) (customerProfile dto.CustomerDto, err error) {
	const op = "customer.usecase.UpdateCustomerProfile"

	customerProfile, err = u.customerPgRepo.GetCustomerProfile(ctx, customerId)
	if err != nil {
		return dto.CustomerDto{}, err
	}

	return customerProfile, nil
}
