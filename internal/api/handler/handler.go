package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {

}

type Service interface {} 

func NewHandler(service Service) *Handler {
	return &Handler{
	}
}

func (h *Handler)  InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		order := api.Group("/order")
		{
			// TODO: add handle function
			order.GET("/:id")
		}
	}

	return router
}
