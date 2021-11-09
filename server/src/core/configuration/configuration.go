package configuration

const (
	ProjectFolder   = "/Projects"
	TemplatesFolder = "/Templates"
)

func GetProjectPath() string {
	return ProjectFolder
}

func GetTemplatesFolder() string {
	return TemplatesFolder
}
