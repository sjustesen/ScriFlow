package models

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          uint `gorm:"primaryKey"`
	ProjectOwnerId uint
	AssignedUsers string
	Name        string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}