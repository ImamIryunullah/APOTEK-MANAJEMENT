package routes

import (
	"apotek-management/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/stok", controllers.GetAllStok)
	r.POST("/stok", controllers.CreateStok)
	r.PUT("/stok/:id", controllers.UpdateStok)
	r.DELETE("/stok/:id", controllers.DeleteStok)
}
