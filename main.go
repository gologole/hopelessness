package main

import (
	"main.go/adminpanel"
	"main.go/logger"
)

func main() {
	logger.CreateLogger()
	adminpanel.StartMenu()
}
