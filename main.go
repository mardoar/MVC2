package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sait.mx/MVC2/controllers"
	"sait.mx/MVC2/model"
)

func main() {
	fmt.Println("Pruebas con bases de Datos Estudiantes")
	model.OpenDB()

	r := gin.Default()
	r.POST("/api/v1/estudiantes", controllers.InsertEstudiante)
	r.GET("/api/v1/estudiantes", controllers.ListEstudiantes)
	r.GET("/api/v1/estudiantes/:id", controllers.GetEstudiante)
	r.PUT("/api/v1/estudiantes/:id", controllers.UpdateEstudiante)
	r.DELETE("/api/v1/estudiantes/:id", controllers.DeleteEstudiante)
	// puerto al servidor
	r.Run(":9053")
}
