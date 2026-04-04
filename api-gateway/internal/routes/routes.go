package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		employee := api.Group("/employee")
		{
			employee.GET("/", handlers.ListEmployees)
			employee.GET("/:id", handlers.GetEmployeeByID)
			employee.POST("/", handlers.CreateEmployee)
			employee.PUT("/:id", handlers.UpdateEmployee)
			employee.DELETE("/:id", handlers.DeleteEmployee)
		}

		category := api.Group("/category")
		{
			category.GET("/", handlers.ListCategories)
			category.GET("/:number", handlers.GetCategoryByNumber)
			category.POST("/", handlers.CreateCategory)
			category.PUT("/:number", handlers.UpdateCategory)
			category.DELETE("/:number", handlers.DeleteCategory)
		}

		product := api.Group("/product")
		{
			product.GET("/", handlers.ListProducts)
			product.GET("/:id", handlers.GetProductByID)
			product.POST("/", handlers.CreateProduct)
			product.PUT("/:id", handlers.UpdateProduct)
			product.DELETE("/:id", handlers.DeleteProduct)
		}

		storeProduct := api.Group("/store-product")
		{
			storeProduct.GET("/", handlers.ListStoreProducts)
			storeProduct.GET("/:upc", handlers.GetStoreProductByUPC)
			storeProduct.POST("/", handlers.CreateStoreProduct)
			storeProduct.PUT("/:upc", handlers.UpdateStoreProduct)
			storeProduct.DELETE("/:upc", handlers.DeleteStoreProduct)
		}

		customerCard := api.Group("/customer-card")
		{
			customerCard.GET("/", handlers.ListCustomerCards)
			customerCard.GET("/:number", handlers.GetCustomerCardByNumber)
			customerCard.POST("/", handlers.CreateCustomerCard)
			customerCard.PUT("/:number", handlers.UpdateCustomerCard)
			customerCard.DELETE("/:number", handlers.DeleteCustomerCard)
		}

		check := api.Group("/check")
		{
			check.GET("/", handlers.ListChecks)
			check.GET("/:number", handlers.GetCheckByNumber)
			check.POST("/", handlers.CreateCheck)
			check.PUT("/:number", handlers.UpdateCheck)
			check.DELETE("/:number", handlers.DeleteCheck)
		}

		sale := api.Group("/sale")
		{
			sale.GET("/", handlers.ListSales)
			sale.GET("/item", handlers.GetSale)
			sale.POST("/", handlers.CreateSale)
			sale.PUT("/item", handlers.UpdateSale)
			sale.DELETE("/item", handlers.DeleteSale)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/login")
		}
	}
}
