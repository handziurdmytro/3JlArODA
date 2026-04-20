package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
)

func parseTimeQuery(c *gin.Context, name string) (time.Time, bool) {
	value := c.Query(name)
	if value == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: name + " is required"})
		return time.Time{}, false
	}

	for _, layout := range []string{time.RFC3339, time.DateOnly} {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			return parsed, true
		}
	}

	c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid " + name + ", expected RFC3339 or YYYY-MM-DD"})
	return time.Time{}, false
}
