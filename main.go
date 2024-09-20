package main

import (
	"dashboardNotes/config"
	"dashboardNotes/controllers"
	"dashboardNotes/migrate"
	"dashboardNotes/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
	var exists bool
	err := config.DB.Model(models.Note{}).Select("count(*) > 0").Where("body = ?", models.NoteRequest.Body).
		Find(&exists).Error
	if err != nil {
		migrate.Main()
	}
}

func main() {
	var URL = "/api/v1/notes"

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))
	r.GET(URL+"/all", controllers.GetNotes)
	r.GET(URL+"/:id", controllers.GetNote)
	r.GET(URL+"/deletedNotes", controllers.GetAllDeleted)
	r.POST(URL+"/", controllers.CreateNote)
	r.PUT(URL+"/:id", controllers.UpdateNote)
	r.DELETE(URL+"/:id", controllers.DeleteNote)

	err := r.Run(":8081")
	if err != nil {
		return
	}
}
