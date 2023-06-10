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
		&models.ClassLabStatus{},
		&models.ClassMemberRole{},
		&models.Course{},
		&models.Lab{},
		&models.LabProblem{},
		&models.LabProblemSubmission{},
		&models.Post{},
		&models.Testcase{},
		&models.TestcaseSubmission{},
		&models.Topic{},
		&models.TopicFileMat{},
	)
}
