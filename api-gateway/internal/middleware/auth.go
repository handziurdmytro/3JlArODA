package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthValidator interface {
	Validate(ctx context.Context, token string) (*pb.ValidateResponse, error)
}

func AuthMiddleware(validator AuthValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{Error: "unauthorized access"})
			return
		}

		prefix := "Bearer "
		if !strings.HasPrefix(header, prefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{Error: "missing token"})
			return
		}

		token := strings.TrimPrefix(header, prefix)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{Error: "empty token"})
			return
		}

		resp, err := validator.Validate(c.Request.Context(), token)
		if err != nil {
			switch status.Code(err) {
			case codes.Unauthenticated, codes.NotFound:
				c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{Error: "invalid token"})
			case codes.Unavailable, codes.DeadlineExceeded:
				c.AbortWithStatusJSON(http.StatusServiceUnavailable, common.ErrorResponse{Error: "authentication service unavailable"})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrorResponse{Error: "failed to validate token"})
			}
			return
		}

		if !resp.IsValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{Error: "invalid token"})
			return
		}

		c.Set("user_id", resp.UserId)
		c.Set("username", resp.Username)
		c.Next()
	}
}
