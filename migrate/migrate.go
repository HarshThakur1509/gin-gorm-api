package main

import (
	"github.com/HarshThakur1509/gin-gorm-api/initializers"
	"github.com/HarshThakur1509/gin-gorm-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	Post := &models.Post{}
	User := &models.User{}
	initializers.DB.Debug().AutoMigrate(Post, User)

}
