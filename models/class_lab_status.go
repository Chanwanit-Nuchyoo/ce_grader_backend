package models

type Class_Lab_Status struct {
	CourseID uint
	ClassID  uint
	LabID    uint
	IsOpen   bool `gorm:"default: false"`
}
