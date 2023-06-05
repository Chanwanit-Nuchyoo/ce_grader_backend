package models

import "time"

type Lab struct {
	ID        uint
	Name      string `gorm:"size:255"`
	CourseID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	LabStatus []Class_Lab_Status `gorm:"foreignKey:LabID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Problems  []LabProblem       `gorm:"foreignKey:LabID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
