package models

import "time"

type User struct {
	ID             string    `gorm:"primaryKey;size:15"`
	Email          string    `gorm:"unique;not null"`
	Password       string    `gorm:"not null;size:255"`
	Name           string    `gorm:"not null;size:255"`
	ImagePath      string    `json:"image_path" gorm:"type:text"`
	LastOneline    time.Time `json:"last_online"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	OwnedTestcases []Testcase        `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	OwnedCourses   []Course          `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	OwnedClasses   []Class           `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	ClassRoles     []ClassMemberRole `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
