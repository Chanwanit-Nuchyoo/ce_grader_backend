package models

import "time"

type Topic struct {
	ID          uint
	CourseID    uint
	Name        string `gorm:"size:255"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
