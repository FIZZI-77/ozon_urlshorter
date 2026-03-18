package handler

import (
	"github.com/gin-gonic/gin"
	"ozon/src/core/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Api is working"})
	})

	api := router.Group("/api")
	{
		api.POST("/shorter", h.CreateShort)
		api.GET("/link/:short", h.GetOriginal)
	}

	return router

}
