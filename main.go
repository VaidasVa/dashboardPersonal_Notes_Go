package main

import (
	"dashboardNotes/config"
	"dashboardNotes/controllers"
	"dashboardNotes/migrate"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVars()
	config.ConnectToDb()
	migrate.Migrate()
}

func main() {
	var URL = "/api/v1/notes"

	r := gin.Default()
	r.GET(URL+"/all", controllers.GetNotes)
	r.GET(URL+"/:id", controllers.GetNote)
	r.GET(URL+"/deletedNotes", controllers.GetAllDeleted)
	r.POST(URL+"/add", controllers.CreateNote)
	r.PUT(URL+"/:id", controllers.UpdateNote)
	r.DELETE(URL+"/:id", controllers.DeleteNote)

	r.Run() // listen and serve on 0.0.0.0:8080
}
