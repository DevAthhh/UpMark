package handlers

import (
	"github.com/DevAthhh/upmark/internal/routes"
	"github.com/gin-gonic/gin"
)

func Handle() *gin.Engine {
	router := gin.Default()

	routes.Route(router)
	return router
}
