package main

import (
	"fmt"
	"main.go/logger"
	video "main.go/videoprocessing"
)

func main() {
	logger.CreateLogger()
	//panel.StartMenu()
	err := video.ChangeSize("zamay.mp4", "output.mp4", 1280, 720)
	if err != nil {
		fmt.Println("Обработка ошибки")
	}
}
