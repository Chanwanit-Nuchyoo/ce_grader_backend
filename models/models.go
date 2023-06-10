package models

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID             uint 
	Email          string    `gorm:"unique;not null"`
	Username       string    `gorm:"unique;not null"`
	Password       string    `gorm:"not null;size:255"`
	Name           string    `gorm:"not null;size:255"`
	ImagePath      string    `json:"image_path" gorm:"type:text"`
	LastOnline     time.Time `json:"last_online"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	OwnedTestcases []Testcase        `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	OwnedCourses   []Course          `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	OwnedClasses   []Class           `gorm:"foreignKey:OwnerID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	OwnedPosts     []Post            `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	ClassRoles     []ClassMemberRole `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	SubmissionRecord     []LabProblemSubmission `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type Course struct {
	ID          uint 
	OwnerID     uint
	Owner User `gorm:"foreignKey:OwnerID;"`
	Name        string             `gorm:"size:255"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LabStatus   []ClassLabStatus    `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Labs        []Lab              `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Topics      []Topic            `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Classes     []Class            `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type Class struct {
	ID          uint
	OwnerID     uint
	Owner User `gorm:"foreignKey:OwnerID;"`
	CourseID    uint
	Name        string             `gorm:"size:255"`
	DefaultLang string             `gorm:"size:15"`
	Description string
	DayOtw      int8
	StartTime   datatypes.Time        
	EndTime     datatypes.Time   
	LabStatus   []ClassLabStatus    `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Members     []ClassMemberRole   `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Lab struct {
	ID        uint 
	Name      string             `gorm:"size:255"`
	CourseID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	LabStatus []ClassLabStatus    `gorm:"foreignKey:LabID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Problems  []LabProblem       `gorm:"foreignKey:LabID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type Testcase struct {
	ID           uint
	OwnerID      uint
	Owner User `gorm:"foreignKey:OwnerID;"`
	LabProblemID uint
	Input        string
	Output       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ClassLabStatus struct {
	ID       uint
	CourseID uint
	ClassID  uint
	LabID    uint
	IsOpen   bool `gorm:"default:false"`
}

type ClassMemberRole struct {
	UserID  uint   `gorm:"primaryKey"`
	ClassID uint   `gorm:"primaryKey"`
	Role    string `gorm:"size:15"`
}

type TopicFileMat struct {
	ID        uint
	Name      string             `gorm:"size:255"`
	TopicID   uint
	FilePath  string             `gorm:"type:text"`
	CreatedAt time.Time
}

type Topic struct {
	ID          uint
	CourseID    uint
	Files []TopicFileMat `gorm:"foreignKey:TopicID;"`
	Name        string             `gorm:"size:255"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

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

type LabProblemSubmission struct {
	ID                  uint
	UserID              uint
	LabProblemID        uint
	Code                string
	Score               int
	Feedback            string
	InstructorScore     int
	InstructorFeedback  string
	TestcaseSubmissions []TestcaseSubmission `gorm:"foreignKey:SubmissionID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type TestcaseSubmission struct {
	ID               uint
	SubmissionID     uint
	TestcaseID       uint
	Input            string
	ExpectedOutput   string
	ActualOutput     string
	Passed           bool
	Score            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Post struct {
	ID        uint
	ClassID   uint
	AuthorID  uint
	Author User `gorm:"foreignKey:AuthorID;"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
