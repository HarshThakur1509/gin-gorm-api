package main

import (
	"github.com/HarshThakur1509/gin-gorm-api/controllers"
	"github.com/HarshThakur1509/gin-gorm-api/initializers"
	"github.com/HarshThakur1509/gin-gorm-api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.Use(cors.Default())

	//User endpoint
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	//Post endpoint
	post := "/posts/"
	r.POST(post, controllers.PostsCreate)
	r.PUT(post+":id", controllers.PostsUpdate)
	r.GET(post, controllers.PostsIndex)
	r.GET(post+":id", controllers.PostsShow)
	r.DELETE(post+":id", controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
