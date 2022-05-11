package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/afaris01/final-project/project-1-todo-apps/controllers"
	docs "github.com/afaris01/final-project/project-1-todo-apps/docs"
)

// @title           ToDo Apps - Rest API
// @version         1.0
// @description     Ini adalah contoh Aplikasi ToDo - Rest API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /to-do

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	Routes := router.Group("/to-do")
	{
		Routes.GET("/", controllers.AmbilAll)
		Routes.GET("/:id", controllers.AmbilTodo)
		Routes.POST("/", controllers.TambahTodo)
		Routes.PUT("/:id", controllers.UbahTodo)
		Routes.DELETE("/:id", controllers.HapusTodo)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
