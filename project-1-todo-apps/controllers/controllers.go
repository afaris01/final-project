package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/todo-apps/models"
)

var (
	todos   = make([]models.Todo, 0, 10)
	counter int
)

type InputModels struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Complete bool   `json:"complete"`
}

// swagger
func AmbilAll(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// swagger
func AmbilTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	singleTodo, err := AmbilById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"pesan": "data tidak ada"})
		return
	}
	c.JSON(http.StatusOK, singleTodo)
}

func AmbilById(id int) (*models.Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("data tidak ada")
}

// swagger
func BuatTodo(c *gin.Context) {
	var (
		inputTodos InputModels
		reqTodos   models.Todo
	)

	err := c.ShouldBindJSON(&inputTodos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"pesan": "Request Invalid",
		})
		return
	}
	counter++
	{
		reqTodos.Id = counter
		reqTodos.Nama = inputTodos.Nama
		reqTodos.Complete = inputTodos.Complete
	}
	todos = append(todos, reqTodos)
	c.JSON(http.StatusCreated, inputTodos)
}

// swagger
func UbahTodo(c *gin.Context) {
	var (
		result     gin.H
		inputTodos InputModels
	)
	if err := c.ShouldBindJSON(&inputTodos); err != nil {
		result = gin.H{
			"error": true,
			"pesan": err,
		}
	}
	inputId := c.Param("id")
	id, _ := strconv.Atoi(inputId)
	for k, v := range todos {
		if v.Id == id {
			todos[k].Nama = inputTodos.Nama
			todos[k].Complete = inputTodos.Complete
			result = gin.H{
				"hasil": todos[k],
			}
		}
	}
	c.JSON(http.StatusOK, result)
}

// swagger
func HapusTodo(c *gin.Context) {
	var result gin.H
	inputId := c.Param("id")
	id, _ := strconv.Atoi(inputId)
	id--
	temp := todos[id]
	todos = append(todos[:id], todos[id+1:]...)
	result = gin.H{
		"todos yang dihapus": temp,
		"todos baru":         todos,
	}
	c.JSON(http.StatusOK, result)
}
