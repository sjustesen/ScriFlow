package core

import (
	"os"
	"path/filepath"

	config "github.com/sjustesen/onlinedtp/core"
)

func listProjects() []string {
	filepath.Walk(config.GetProjectPath(), func(path string, fileinfo os.FileInfo, err error) error {
	})
}

func ProcessFile() {

}
