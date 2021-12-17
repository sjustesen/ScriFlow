package main

import (
	"fmt"

	"github.com/sjustesen/scriflow/core/configuration"
	"github.com/sjustesen/scriflow/server"
)

func main() {
	projectpath := configuration.GetProjectPath()
	fmt.Println(projectpath)
	server.Bootup()
}
