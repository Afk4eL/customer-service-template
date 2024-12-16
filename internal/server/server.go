package server

import (
	"customer-service/config"
	"customer-service/internal/customer"
	"customer-service/internal/customer/repository"
	"customer-service/internal/customer/usecase"
	"customer-service/pkg/logger"
	"fmt"
	"net"

	//customer_contracts "github.com/Afk4eL/wildberries2.0/gen/go/customer-service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	cfg        *config.Config
	db         *gorm.DB
	gRPCServer *grpc.Server
}

func NewCustomerServer(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) grpcRun(customerUC customer.CustomerUseCase) error {
	const op = "grpc_app.grpcRun"

	s.gRPCServer = grpc.NewServer()

	customergRPCServer := grpc_service.customerService(s.cfg, customerUC)

	customer_contracts.RegisterCustomerServiceServer(s.gRPCServer, customergRPCServer)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Server.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer l.Close()

	logger.Logger.Info().Msg(fmt.Sprintf("gRRPC server started on %d port", s.cfg.Server.Port))

	if err := s.gRPCServer.Serve(l); err != nil {
		logger.Logger.Fatal().Str(op, err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Run() error {
	const op = "grpc_app.Run"

	userRepo := repository.NewCustomerRepository(s.db)

	userUC := usecase.NewCustomerUsecase(userRepo)

	go s.grpcRun(userUC)

	return nil
}

func (a *Server) Stop() {
	logger.Logger.Info().Msg("stopping gRPC server")

	pg, _ := a.db.DB()
	pg.Close()
	logger.Logger.Info().Msg("Postgres stopped")

	a.gRPCServer.GracefulStop()
	logger.Logger.Info().Msg("gRPC server stopped")
}
