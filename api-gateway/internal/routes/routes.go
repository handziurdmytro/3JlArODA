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
	customerCardHandler *handlers.CustomerCardHandler,
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

		employee := protected.Group("/employee")
		{
			employee.GET("/", handlers.ListEmployees)
			employee.GET("/:id", handlers.GetEmployeeByID)
			employee.POST("/", handlers.CreateEmployee)
			employee.PUT("/:id", handlers.UpdateEmployee)
			employee.DELETE("/:id", handlers.DeleteEmployee)
		}

		category := protected.Group("/category")
		{
			category.GET("/", handlers.ListCategories)
			category.GET("/:number", handlers.GetCategoryByNumber)
			category.POST("/", handlers.CreateCategory)
			category.PUT("/:number", handlers.UpdateCategory)
			category.DELETE("/:number", handlers.DeleteCategory)
		}

		product := protected.Group("/product")
		{
			product.GET("/", productHandler.List)
			product.GET("/:id", productHandler.GetByID)
			product.POST("/", productHandler.Create)
			product.PUT("/:id", productHandler.Update)
			product.DELETE("/:id", productHandler.Delete)
		}

		storeProduct := protected.Group("/store-product")
		{
			storeProduct.GET("/", handlers.ListStoreProducts)
			storeProduct.GET("/:upc", handlers.GetStoreProductByUPC)
			storeProduct.POST("/", handlers.CreateStoreProduct)
			storeProduct.PUT("/:upc", handlers.UpdateStoreProduct)
			storeProduct.DELETE("/:upc", handlers.DeleteStoreProduct)
		}

		customerCard := protected.Group("/customer-card")
		{
			customerCard.GET("/", customerCardHandler.List)
			customerCard.GET("/:number", customerCardHandler.GetByNumber)
			customerCard.POST("/", customerCardHandler.Create)
			customerCard.PUT("/:number", customerCardHandler.Update)
			customerCard.DELETE("/:number", customerCardHandler.Delete)
		}

		check := protected.Group("/check")
		{
			check.GET("/", handlers.ListChecks)
			check.GET("/:number", handlers.GetCheckByNumber)
			check.POST("/", handlers.CreateCheck)
			check.DELETE("/:number", handlers.DeleteCheck)
		}

		sale := protected.Group("/sale")
		{
			sale.GET("/", handlers.ListSales)
			sale.GET("/item", handlers.GetSale)
			sale.POST("/", handlers.CreateSale)
		}
	}
}
