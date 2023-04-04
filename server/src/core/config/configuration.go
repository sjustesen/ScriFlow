package config

import (
	"os"
)

const (
	ProjectFolder   = "/Projects"
	TemplatesFolder = "/Templates"
	dsn             = "simon:verystrongpassword!@tcp(127.0.0.1:3306)/scriflow?charset=utf8mb4&parseTime=True&loc=Local"
)

func GetSystemPath() string {
	var path = ""
	path, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return path
}

func GetDatabaseDSN() string {
	return dsn
}

func GetProjectPath() string {
	path, _ := os.Getwd()
	return path + ProjectFolder
}

func GetTemplatesFolder() string {
	return TemplatesFolder
}
