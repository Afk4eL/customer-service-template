package grpc_service

import (
	"context"
	"customer-service/pkg/grpc_errors"

	"google.golang.org/grpc/status"
)

func (s *customerService) GetCustomerProfile(
	ctx context.Context,
	req *customer_contracts.GetCustomerProfileRequest) (*customer_contracts.GetCustomerProfileResponse, error) {

	customerId := req.GetCustomerId()
	customerProfile, err := s.customerUC.GetCustomerProfile(ctx, customerId)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "Get profile failed: %v", err)
	}

	return &customer_contracts.GetCustomerProfileResponse{CustomerProfile: customerProfile}, nil
}

func (s *customerService) UpdateCustomerProfile(
	ctx context.Context,
	req *customer_contracts.UpdateCustomerProfileRequest) (*customer_contracts.UpdateCustomerProfileResponse, error) {

	customerId := req.GetCustomerId()
	updatedCustomerProfile := req.GetCustomerProfile()
	customerProfile, err := s.customerUC.UpdateCustomerProfile(ctx, customerId, updatedCustomerProfile)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "Update profile failed: %v", err)
	}

	return &customer_contracts.UpdateCustomerProfileResponse{CustomerProfile: customerProfile}, nil
}
