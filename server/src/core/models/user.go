package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserGroups struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	UsersInGroup []User `gorm:"foreignKey:Id;references:Id"`
}
