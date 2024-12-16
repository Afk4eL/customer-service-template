package grpc_errors

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
)

// здесь пропишем все возможные ошибки на сервисе
var (
	ErrNotFound      = errors.New("Not found")
	ErrNoCtxMetaData = errors.New("No ctx metadata")
)

func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows) || errors.Is(err, ErrNotFound):
		return codes.NotFound
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case strings.Contains(err.Error(), "Validate"):
		return codes.InvalidArgument
	}
	return codes.Internal
}

func MapGRPCErrCodeToHttpStatus(code codes.Code) int {
	switch code {
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.AlreadyExists:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.InvalidArgument:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
