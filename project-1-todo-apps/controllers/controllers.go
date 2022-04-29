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

type todoModels struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Selesai bool   `json:"Selesai"`
}

// @Tags to-do
// @Summary Menampilkan semua ToDo
// @Description Menampilkan semua daftar data ToDo
// @Produce json
// @Success 200 {array} models.Todo
// @Router /to-do [get]
func AmbilAll(c *gin.Context) {
	if todos != [] {
		c.JSON(http.StatusOK, todos)
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"pesan": "data tidak ada"})
	}

}

// @Summary      Menampilkan sebuah data ToDo
// @Description  Menampilkan sebuah data ToDo berdasarkan Id yang dimasukkan
// @Tags         to-do
// @Accept       json
// @Produce      json
// @Param        id   path int  true  "Ambil By Id"
// @Success      200  {object}  models.Todo
// @Router       /to-do/{id} [get]
func AmbilTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	hasil, err := AmbilById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"pesan": "data tidak ada"})
		return
	}
	c.JSON(http.StatusOK, hasil)
}

func AmbilById(id int) (*models.Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("data tidak ada")
}

// @Summary Menambahkan ToDo Baru
// @Description Menambahkan satu data ToDo baru
// @Tags to-do
// @Accept json
// @Produce json
// @Success 200 {object} models.Todo
// @Param todo body models.Todo true "Buat Todo"
// @Router /to-do [post]
func TambahTodo(c *gin.Context) {
	var (
		tambahTodo todoModels
		reqTodos   models.Todo
	)

	err := c.ShouldBindJSON(&tambahTodo)
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
		reqTodos.Nama = tambahTodo.Nama
		reqTodos.Selesai = tambahTodo.Selesai
	}
	todos = append(todos, reqTodos)
	c.JSON(http.StatusCreated, tambahTodo)
}

// @Summary Mengubah ToDo
// @Description Mengubah ToDo yang sudah ada
// @Tags to-do
// @Accept json
// @Produce json
// @Success 200 {object} models.Todo
// @Param id path int true "Ubah ToDo"
// @Param todo body models.Todo true "minta data json"
// @Router /to-do/{id} [put]
func UbahTodo(c *gin.Context) {
	var (
		hasil    gin.H
		ubahTodo todoModels
	)
	if err := c.ShouldBindJSON(&ubahTodo); err != nil {
		hasil = gin.H{
			"error": true,
			"pesan": err,
		}
	}
	inputId := c.Param("id")
	id, _ := strconv.Atoi(inputId)
	for k, v := range todos {
		if v.Id == id {
			todos[k].Nama = ubahTodo.Nama
			todos[k].Selesai = ubahTodo.Selesai
			hasil = gin.H{
				"hasil": todos[k],
			}
		}
	}
	c.JSON(http.StatusOK, hasil)
}

// @Summary Menghapus ToDo
// @Description Menghapus ToDo yang sudah ada
// @Tags to-do
// @Accept json
// @Produce json
// @Success 200 {object} models.Todo
// @Param id path int true "Hapus ToDo"
// @Router /to-do/{id} [delete]
func HapusTodo(c *gin.Context) {
	var hasil gin.H
	inputId := c.Param("id")
	id, _ := strconv.Atoi(inputId)
	id--
	temp := todos[id]
	todos = append(todos[:id], todos[id+1:]...)
	hasil = gin.H{
		"ToDo yang dihapus": temp,
		"ToDo baru":         todos,
	}
	c.JSON(http.StatusOK, hasil)
}
