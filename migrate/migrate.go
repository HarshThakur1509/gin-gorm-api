package main

import (
	"GinGorm/initializers"
	"GinGorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	Post := &models.Post{}
	initializers.DB.Debug().AutoMigrate(Post)

}
