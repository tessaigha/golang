package routes

import (
	"go-kasir-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/barang", controllers.GetBarang)
	r.POST("/barang", controllers.AddBarang)
	r.DELETE("/barang/:id", controllers.DeleteBarang)
	r.PUT("/barang/:id", controllers.UpdateBarang)
	r.GET("/barang/:id", controllers.GetBarangByID)

	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.AddUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.GET("/user/:id", controllers.GetUserByID)

	r.GET("/pembelian", controllers.GetPembelian)
	// r.POST("/pembelian", controllers.AddPembelian)
	// r.DELETE("/pembelian/:id", controllers.DeleteUser)
	// r.PUT("/pembelian/:id", controllers.UpdateUser)
	// r.GET("/pembelian/:id", controllers.GetUserByID)
	return r
}
