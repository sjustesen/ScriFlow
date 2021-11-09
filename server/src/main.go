package main

import (
	"fmt"

	"github.com/sjustesen/onlinedtp/core/configuration"
	"github.com/sjustesen/onlinedtp/server"
)

func main() {
	projectpath := configuration.GetProjectPath()
	fmt.Println(projectpath)
	server.Bootup()
}
