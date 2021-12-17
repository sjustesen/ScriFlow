package main

import (
	"fmt"
	"github.com/sjustesen/scriflow/core"
	"github.com/sjustesen/scriflow/core/config"
	"github.com/sjustesen/scriflow/server"
)

func main() {
	database.Open()
	database.Setup()
	projectpath := config.GetProjectPath()
	fmt.Println(projectpath)
	server.Bootup()
}
