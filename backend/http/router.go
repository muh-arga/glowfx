package http

import (
	"glow-fx/repository"
	"glow-fx/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	productRepo := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productHandler := NewProductHandler(productUseCase)

	productRoutes := api.Group("/products")
	{
		productRoutes.GET("", productHandler.GetProducts)
		productRoutes.GET("/:id", productHandler.GetProductById)
		productRoutes.POST("", productHandler.CreateProduct)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

}
