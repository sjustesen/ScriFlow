package core

const (
	ProjectFolder = "/Public"
	TemplatesFolder = "/Templates"
)

func GetProjectPath() string {
	return ProjectFolder;
}

func GetTemplatesFolder() string {
	return TemplatesFolder;
}