package models

import "time"

type LabProblem struct {
	ID            uint
	LabID         uint
	Name          string `gorm:"size:255"`
	Prompt        string
	FullScore     float32
	ScoringMethod string
	Testcases     []Testcase `gorm:"foreignKey:LabProblemID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
