package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sumansurvi/go-crud/config"
	"github.com/sumansurvi/go-crud/controllers/getItems"
)

func StartServer(appContext *config.AppContext) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// CRUD routes
	r.GET("/items", GenericHandler(appContext, getItems.GetAllItems))
	// r.GET("/items/:id", getItem)
	// r.POST("/items", createItem)
	// r.PUT("/items/:id", updateItem)
	// r.DELETE("/items/:id", deleteItem)

	return r
}
