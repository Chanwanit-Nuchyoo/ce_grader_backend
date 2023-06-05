package main

import (
	"ce_grader_backend/db"
	"ce_grader_backend/environment"
	"ce_grader_backend/models"
)

func main() {
	environment.LoadEnvVariables()
	db := db.GetDB()

	db.AutoMigrate(
		&models.Class{},
		&models.ClassMemberRole{},
		&models.Class_Lab_Status{},
		&models.Course{},
		&models.Lab{},
		&models.LabProblem{},
		&models.Testcase{},
		&models.Topic{},
		&models.Topic_File_Mat{},
		&models.User{},
	)
}
