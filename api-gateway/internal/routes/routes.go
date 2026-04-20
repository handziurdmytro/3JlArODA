package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/handlers"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/middleware"
)

func SetupRoutes(
	router *gin.Engine,
	authHandler *handlers.AuthHandler,
	authValidator middleware.AuthValidator,
	productHandler *handlers.ProductHandler,
	categoryHandler *handlers.CategoryHandler,
	storeProductHandler *handlers.StoreProductHandler,
	customerCardHandler *handlers.CustomerCardHandler,
	employeeHandler *handlers.EmployeeHandler,
	checkHandler *handlers.CheckHandler,
	saleHandler *handlers.SaleHandler,
) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(authValidator))

		employee := protected.Group("/employees")
		{
			employee.GET("/me", employeeHandler.GetMe)
			employee.GET("/contacts", employeeHandler.GetContacts)
			employee.GET("/", employeeHandler.List)
			employee.GET("/:id", employeeHandler.GetByID)
			employee.POST("/", employeeHandler.Create)
			employee.PUT("/:id", employeeHandler.Update)
			employee.DELETE("/:id", employeeHandler.Delete)
		}

		category := protected.Group("/categories")
		{
			category.GET("/", categoryHandler.List)
			category.GET("/:number", categoryHandler.GetByNumber)
			category.POST("/", categoryHandler.Create)
			category.PUT("/:number", categoryHandler.Update)
			category.DELETE("/:number", categoryHandler.Delete)
		}

		product := protected.Group("/products")
		{
			product.GET("/", productHandler.List)
			product.GET("/:id", productHandler.GetByID)
			product.POST("/", productHandler.Create)
			product.PUT("/:id", productHandler.Update)
			product.DELETE("/:id", productHandler.Delete)
		}

		storeProduct := protected.Group("/store-products")
		{
			storeProduct.GET("/", storeProductHandler.List)
			storeProduct.GET("/:upc", storeProductHandler.GetByUPC)
			storeProduct.POST("/", storeProductHandler.Create)
			storeProduct.PUT("/:upc", storeProductHandler.Update)
			storeProduct.DELETE("/:upc", storeProductHandler.Delete)
		}

		customerCard := protected.Group("/customer-cards")
		{
			customerCard.GET("/", customerCardHandler.List)
			customerCard.GET("/:number", customerCardHandler.GetByNumber)
			customerCard.POST("/", customerCardHandler.Create)
			customerCard.PUT("/:number", customerCardHandler.Update)
			customerCard.DELETE("/:number", customerCardHandler.Delete)
		}

		check := protected.Group("/checks")
		{
			check.GET("/", checkHandler.List)
			check.GET("/:number", checkHandler.GetByNumber)
			check.POST("/", checkHandler.Create)
			check.POST("/:number/items", saleHandler.Create)
			check.DELETE("/:number", checkHandler.Delete)
		}

		reports := protected.Group("/reports")
		{
			reports.GET("/checks/details", checkHandler.GetDetailsReport)
			reports.GET("/checks/total", checkHandler.GetTotalReport)
			reports.GET("/products/:id/sold-quantity", saleHandler.GetProductSoldQuantity)
		}

		individualTasks := protected.Group("/individual-tasks")
		{
			individualTasks.GET("/products/sales-stats", productHandler.GetSalesStatsByCategoryAndPeriod)
			individualTasks.GET("/customer-cards/bought-all-products-from-category", customerCardHandler.GetWhoBoughtAllProductsFromCategory)
			individualTasks.GET("/categories/stock-summary", categoryHandler.GetStockSummary)
			individualTasks.GET("/employees/cashier-performance", employeeHandler.GetCashierPerformance)
			individualTasks.GET("/employees/best-cashiers-by-promo", employeeHandler.GetBestCashiersByPromo)
			individualTasks.GET("/store-products/cashiers-sold-all-category-products", storeProductHandler.GetCashiersWhoSoldAllProductsFromCategory)
		}
	}
}
