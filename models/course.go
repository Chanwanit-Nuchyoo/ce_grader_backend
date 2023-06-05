package models

import "time"

type Course struct {
	ID          uint
	OwnerID     string
	Owner       User   `gorm:"foreignKey:OwnerID"`
	Name        string `gorm:"size:255"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LabStatus   []Class_Lab_Status `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Labs        []Lab              `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Topics      []Topic            `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Classes     []Class            `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
