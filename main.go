package main

import (
	"dashboardNotes/config"
	"dashboardNotes/controllers"
	"dashboardNotes/migrate"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVars()
	config.ConnectToDb()
	migrate.Migrate()
}

func main() {
	var URL = "/api/v1/notes"

	router := gin.Default()
	routerConfig := cors.DefaultConfig()
	routerConfig.AllowAllOrigins = true
	router.Use(cors.New(routerConfig))
	router.GET(URL+"/all", controllers.GetNotes)
	router.GET(URL+"/:id", controllers.GetNote)
	router.GET(URL+"/deletedNotes", controllers.GetAllDeleted)
	router.POST(URL+"/", controllers.CreateNote)
	router.PUT(URL+"/:id", controllers.UpdateNote)
	router.DELETE(URL+"/:id", controllers.DeleteNote)

	router.Run() // listen and serve on 0.0.0.0:8080
}
