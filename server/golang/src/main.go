package main

import (
	"fmt"

	"github.com/sjustesen/scriflow/core/config"
	"github.com/sjustesen/scriflow/core/database"
	"github.com/sjustesen/scriflow/server"
)

func main() {
	database.Open()
	projectpath := config.GetProjectPath()
	fmt.Println(projectpath)
	server.Bootup()
}
