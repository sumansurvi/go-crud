package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sumansurvi/go-crud/config"
)

// HandlerFunc is a type that represents a handler with access to AppContext
type HandlerFunc func(*gin.Context, *config.AppContext)

// GenericHandler is a higher-order function that takes your custom handler and returns gin.HandlerFunc
func GenericHandler(appCtx *config.AppContext, handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pass gin context and app context to the handler
		handler(c, appCtx)
	}
}
