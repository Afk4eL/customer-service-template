package grpc_service

import (
	"customer-service/config"
	"customer-service/internal/customer"
	//customer_contracts "github.com/Afk4eL/wildberries2.0/gen/go/customer-service"
)

type customerService struct {
	customer_contracts.UnimplementedUsersStorageServer

	cfg        *config.Config
	customerUC customer.CustomerUseCase
}

func NewCustomerServiceGRPC(cfg *config.Config, customerUC customer.CustomerUseCase) *customerService {
	return &customerService{
		cfg:        cfg,
		customerUC: customerUC,
	}
}
