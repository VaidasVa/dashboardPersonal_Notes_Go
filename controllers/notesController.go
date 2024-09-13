package controllers

import (
	"dashboardNotes/config"
	"dashboardNotes/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateNote(c *gin.Context) {
	err := c.BindJSON(&models.NoteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	note := models.Note{Title: models.NoteRequest.Title, Body: models.NoteRequest.Body}
	result := config.DB.Create(&note)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	c.JSON(http.StatusCreated, gin.H{"note": note})
}

func GetNotes(c *gin.Context) {
	var notes []models.Note
	result := config.DB.Find(&notes)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func GetNote(c *gin.Context) {
	var note models.Note
	result := config.DB.First(&note, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"note": note})
	}
}

func GetAllDeleted(c *gin.Context) {
	var notes []models.Note
	result := config.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	err := c.Bind(&models.NoteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var note models.Note
	result := config.DB.First(&note, id)
	if result.Error != nil {
		log.Fatal(result.Error, "Note not found for update")
	}
	config.DB.Model(&note).Updates(models.Note{
		Title: models.NoteRequest.Title,
		Body:  models.NoteRequest.Body})

	c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	result := config.DB.First(&models.Note{}, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
	}
	config.DB.Delete(&models.Note{}, id)
	c.JSON(http.StatusNoContent, gin.H{"Deleted": "yes"})
}
