package config

const (
	ProjectFolder   = "/Projects"
	TemplatesFolder = "/Templates"
	dsn = "simon:strongpass@tcp(127.0.0.1:3306)/scriflow?charset=utf8mb4&parseTime=True&loc=Local"

)

func GetDatabaseDSN() string {
	return dsn
}

func GetProjectPath() string {
	return ProjectFolder
}

func GetTemplatesFolder() string {
	return TemplatesFolder
}
