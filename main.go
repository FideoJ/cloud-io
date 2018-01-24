package main

import (
	"github.com/FideoJ/cloudgo-io/service"
)

func main() {
	server := service.NewServer()
	server.Run(":8080")
}
