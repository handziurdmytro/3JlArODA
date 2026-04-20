package handlers

import (
	"net/http"
	"strconv"
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

func parseIntQuery(c *gin.Context, name string) (int, bool) {
	value := c.Query(name)
	if value == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: name + " is required"})
		return 0, false
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid " + name})
		return 0, false
	}

	return parsed, true
}

func parseFloatQuery(c *gin.Context, name string) (float64, bool) {
	value := c.Query(name)
	if value == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: name + " is required"})
		return 0, false
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid " + name})
		return 0, false
	}

	return parsed, true
}
