package migrate

import (
	"dashboardNotes/config"
	"dashboardNotes/models"
)

func init() {
	//config.LoadEnvVars()
	config.ConnectToDb()
}

func Migrate() {
	var exists bool
	err := config.DB.Model(models.Note{}).Select("count(*) > 0").Where("body = ?", models.NoteRequest.Body).
		Find(&exists).Error
	if err != nil {
		err := config.DB.AutoMigrate(&models.Note{})
		if err != nil {
			return
		}
	}
}
