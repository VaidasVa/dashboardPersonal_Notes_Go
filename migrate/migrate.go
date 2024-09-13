package migrate

import (
	"dashboardNotes/config"
	"dashboardNotes/models"
)

func init() {
	config.LoadEnvVars()
	config.ConnectToDb()
}

func Main() {
	err := config.DB.AutoMigrate(&models.Note{})
	if err != nil {
		return
	}
}
