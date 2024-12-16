package repository

import (
	"context"
	"customer-service/internal/domain/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	database *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{database: db}
}

func (r *CustomerRepository) GetCustomerProfile(
	ctx context.Context,
	customerId uuid.UUID) (customerProfile dto.CustomerDto, err error) {
	const op = "customer.repository.GetCustomerProfile"

	return customerProfile, nil
}

func (r *CustomerRepository) UpdateCustomerProfile(
	ctx context.Context,
	customerId uuid.UUID,
	customerDto dto.CustomerDto) (customerProfile dto.CustomerDto, err error) {
	const op = "customer.repository.GetCustomerProfile"

	return customerProfile, nil
}
