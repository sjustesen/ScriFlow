package database

import (
	"database/sql"
	"time"

	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConn struct {
	username string
	password string
}

type Project struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (db *gorm.DB) Setup() {
	db.AutoMigrate(&Project)
}

func down() {
}
