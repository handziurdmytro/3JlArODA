package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/common"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	storeproductpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/storeproduct"
)

type StoreProductClient interface {
	Create(ctx context.Context, req models.CreateStoreProductRequest) (*storeproductpb.StoreProduct, error)
	Update(ctx context.Context, upc string, req models.UpdateStoreProductRequest) (*storeproductpb.StoreProduct, error)
	Delete(ctx context.Context, upc string) error
	GetByUPC(ctx context.Context, upc string) (*storeproductpb.DetailedStoreProduct, error)
	GetAllSortedByQuantity(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error)
	GetAllSortedByName(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error)
	GetPromoSortedByQuantity(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error)
	GetPromoSortedByName(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error)
	GetNonPromoSortedByQuantity(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error)
	GetNonPromoSortedByName(ctx context.Context) ([]*storeproductpb.DetailedStoreProduct, error)
	GetByCategorySortedByName(ctx context.Context, categoryNumber int) ([]*storeproductpb.DetailedStoreProduct, error)
	GetCashiersWhoSoldAllProductsFromCategory(ctx context.Context, categoryNumber int, from, to time.Time) ([]*storeproductpb.CashierSoldAllCategoryProducts, error)
}

type StoreProductHandler struct {
	storeProductClient StoreProductClient
}

func NewStoreProductHandler(storeProductClient StoreProductClient) *StoreProductHandler {
	return &StoreProductHandler{storeProductClient: storeProductClient}
}

func (h *StoreProductHandler) Create(c *gin.Context) {
	var req models.CreateStoreProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	product, err := h.storeProductClient.Create(c.Request.Context(), req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *StoreProductHandler) GetByUPC(c *gin.Context) {
	upc := c.Param("upc")
	if upc == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "UPC is required"})
		return
	}

	product, err := h.storeProductClient.GetByUPC(c.Request.Context(), upc)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *StoreProductHandler) List(c *gin.Context) {
	categoryNumber := c.Query("category_number")
	if categoryNumber != "" {
		parsed, err := strconv.Atoi(categoryNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid category_number"})
			return
		}

		products, err := h.storeProductClient.GetByCategorySortedByName(c.Request.Context(), parsed)
		if err != nil {
			respondGRPCError(c, err)
			return
		}

		c.JSON(http.StatusOK, products)
		return
	}

	promo := c.Query("promo")
	sortBy := c.DefaultQuery("sort", "quantity")

	var (
		products []*storeproductpb.DetailedStoreProduct
		err      error
	)

	switch {
	case promo == "true" && sortBy == "name":
		products, err = h.storeProductClient.GetPromoSortedByName(c.Request.Context())
	case promo == "true":
		products, err = h.storeProductClient.GetPromoSortedByQuantity(c.Request.Context())
	case promo == "false" && sortBy == "name":
		products, err = h.storeProductClient.GetNonPromoSortedByName(c.Request.Context())
	case promo == "false":
		products, err = h.storeProductClient.GetNonPromoSortedByQuantity(c.Request.Context())
	case sortBy == "name":
		products, err = h.storeProductClient.GetAllSortedByName(c.Request.Context())
	case sortBy == "quantity":
		products, err = h.storeProductClient.GetAllSortedByQuantity(c.Request.Context())
	default:
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "invalid sort"})
		return
	}

	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *StoreProductHandler) GetCashiersWhoSoldAllProductsFromCategory(c *gin.Context) {
	categoryNumber, ok := parseIntQuery(c, "category_number")
	if !ok {
		return
	}
	from, to, ok := parsePeriod(c)
	if !ok {
		return
	}

	cashiers, err := h.storeProductClient.GetCashiersWhoSoldAllProductsFromCategory(c.Request.Context(), categoryNumber, from, to)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, cashiers)
}

func (h *StoreProductHandler) Update(c *gin.Context) {
	upc := c.Param("upc")
	if upc == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "UPC is required"})
		return
	}

	var req models.UpdateStoreProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Invalid request: " + err.Error()})
		return
	}

	product, err := h.storeProductClient.Update(c.Request.Context(), upc, req)
	if err != nil {
		respondGRPCError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *StoreProductHandler) Delete(c *gin.Context) {
	upc := c.Param("upc")
	if upc == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "UPC is required"})
		return
	}

	if err := h.storeProductClient.Delete(c.Request.Context(), upc); err != nil {
		respondGRPCError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
