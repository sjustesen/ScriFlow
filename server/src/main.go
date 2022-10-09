package main

import (
	"github.com/sjustesen/scriflow/core/database"
	"github.com/sjustesen/scriflow/server"
)

func main() {
	database.Open()
	server.Bootup()
}
