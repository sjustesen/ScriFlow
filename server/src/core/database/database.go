package database

import (
	"log"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/sjustesen/scriflow/core/config"
	"github.com/sjustesen/scriflow/core/models"
	"github.com/harranali/authority"

)

type DbConn struct {
	username string
	password string
}

func Setup() {
	
}

func Open() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.GetDatabaseDSN()), &gorm.Config{})
	db.AutoMigrate(models.Project{}, models.User{}, models.UserGroups{})
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
