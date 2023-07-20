package links

import (
	"github.com/gin-gonic/gin"
	"go-ushort/app/routers/middlewares"
)

func Router(route *gin.RouterGroup) {
	route.Use(middlewares.AuthMiddleware(true))
	route.GET("/", GetAll)
	route.POST("/", Create)
	route.PATCH("/:linkId", Update)

}
