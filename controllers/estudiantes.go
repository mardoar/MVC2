package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"sait.mx/MVC2/model" // MVC2 nombre de la carpeta prncipal
)

func InsertEstudiante(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	var estudiante model.Estudiante
	err = json.Unmarshal(body, &estudiante)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	err = model.InsertEstudiante(estudiante)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Insertado correctamente")
}

func ListEstudiantes(c *gin.Context) {
	estudiantes, err := model.ListEstudiantes()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, estudiantes)
}

func GetEstudiante(c *gin.Context) {
	id := c.Param("id")
	cliente, err := model.GetEstudiante(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, cliente)
}

func UpdateEstudiante(c *gin.Context) {
	id := c.Param("id")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var estudiante model.Estudiante
	err = json.Unmarshal(body, &estudiante)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	estudiante.ID = id
	err = model.UpdateEstudiante(estudiante)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Actualizado correctamente")
}

func DeleteEstudiante(c *gin.Context) {
	id := c.Param("id")
	err := model.DeleteEstudiante(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "Eliminado correctamente")
}
