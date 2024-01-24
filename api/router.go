package api

import (
	"github.com/gin-gonic/gin"
	"test/api/handler"
	"test/storage"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.New(store)

	r := gin.New()

	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUser)
	r.GET("/users", h.GetUserList)
	r.PUT("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	r.PATCH("/user/:id", h.UpdateUserPassword)

	r.POST("/category", h.CreateCategory)
	r.GET("/category/:id", h.GetCategory)
	r.GET("/categories", h.GetCategoryList)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetProduct)
	r.GET("/products", h.GetProductList)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)

	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetBasket)
	r.GET("/basket", h.GetBasketList)
	r.PUT("basket/:id", h.UpdateBasket)
	r.DELETE("basket/:id", h.DeleteBasket)
	return r
}
