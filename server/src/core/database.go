package core

import (
	"log"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/sjustesen/scriflow/core/models"
	"github.com/harranali/authority"

)

type DbConn struct {
	username string
	password string
}

var db *gorm.DB

func Setup() {
	db.AutoMigrate(models.Project{})
}

func Open() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/scriflow?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect")
	}
	return db
}

func SetupAuthRoles() {
	fmt.Println("Setting up roles...")

	db := Open()

	auth := authority.New(
		authority.Options{
			TablesPrefix: "authority_",DB: db})

	err := auth.CreateRole("admin")
	if err != nil {
		log.Fatal("Couldn't create role admin")
	}
}

func down() {
}
