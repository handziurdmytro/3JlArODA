package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func respondGRPCError(c *gin.Context, err error) {
	switch status.Code(err) {
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: status.Convert(err).Message()})
	case codes.NotFound:
		c.JSON(http.StatusNotFound, common.ErrorResponse{Error: status.Convert(err).Message()})
	case codes.AlreadyExists:
		c.JSON(http.StatusConflict, common.ErrorResponse{Error: status.Convert(err).Message()})
	case codes.Unauthenticated:
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: status.Convert(err).Message()})
	case codes.PermissionDenied:
		c.JSON(http.StatusForbidden, common.ErrorResponse{Error: status.Convert(err).Message()})
	case codes.Unavailable, codes.DeadlineExceeded:
		c.JSON(http.StatusServiceUnavailable, common.ErrorResponse{Error: "service unavailable"})
	default:
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "internal server error"})
	}
}
