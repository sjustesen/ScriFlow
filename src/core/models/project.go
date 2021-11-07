package models

import (
	"database/sql"
	"time"
)

type Project struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
