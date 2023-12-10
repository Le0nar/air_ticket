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

	// TODO: add routes here

	return router
}
