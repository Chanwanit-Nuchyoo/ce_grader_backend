package models

type ClassMemberRole struct {
	UserID  string `gorm:"primaryKey"`
	ClassID uint   `gorm:"primaryKey;autoIncrement:false"`
	Role    string `gorm:"size:15"`
}
