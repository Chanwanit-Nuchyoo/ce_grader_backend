package models

import "time"

type Testcase struct {
	ID           uint
	OwnerID      string
	Owner        User `gorm:"foreignKey:OwnerID"`
	LabProblemID uint
	Input        string
	Output       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
