package router

import (
	"oracle-crud-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(personHandler *handler.PersonHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/persons", personHandler.Create)
		api.GET("/persons/:id", personHandler.GetByID)
		api.GET("/persons", personHandler.GetAll)
		api.PUT("/persons/:id", personHandler.Update)
		api.DELETE("/persons/:id", personHandler.Delete)
	}

	return r
}
