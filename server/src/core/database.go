package core

import (
	"log"

	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConn struct {
	username string
	password string
}

func (db *gorm.DB) Setup() {
	db.AutoMigrate(&Project{})
}

func (db *gorm.DB) Open() {
	dsn := "root@tcp(127.0.0.1:3306)/wwfs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect")
	}
	return db
}
func down() {
}
