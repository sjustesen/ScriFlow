package project

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sjustesen/scriflow/core/config"
)

func ListScribusProjects() []string {
	var files []string
	filepath.Walk(config.GetProjectPath(), func(path string, fileinfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("FilePath Walking Error")
			fmt.Println(err)
			return err
		}

		if !fileinfo.IsDir() && filepath.Ext(path) == ".sla" {
			files = append(files, path)
		}
		return nil
	})

	return files
}
