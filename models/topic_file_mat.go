package models

import "time"

type Topic_File_Mat struct {
	ID        uint
	Name      string `gorm:"size:255"`
	TopicID   uint
	FilePath  string `gorm:"type:text"`
	CreatedAt time.Time
}
