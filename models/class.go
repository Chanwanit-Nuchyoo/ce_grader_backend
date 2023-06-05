package models

import "time"

type Class struct {
	ID          uint
	OwnerID     string
	Owner       User `gorm:"foreignKey:OwnerID"`
	CourseID    uint
	Name        string `gorm:"size:255"`
	DefaultLang string `gorm:"size:15"`
	Description string
	DayOtw      int8
	StartTime   time.Time          `gorm:"type:time"`
	EndTime     time.Time          `gorm:"type:time"`
	LabStatus   []Class_Lab_Status `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Members     []ClassMemberRole  `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
