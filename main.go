package main

import (
	"main.go/adminpanel"
	"main.go/logger"
)

func main() {
	logger.CreateLogger()
	adminpanel.StartMenu()

	//configs.SetPATHFFMPEG() //сделать обработку ошибок и добавить функцию в админ панель

	//err := video.ChangeSize("zamay.mp4", "output.mp4", 1280, 120)
	//if err != nil {
	//	fmt.Println("Обработка ошибки")
	//}
	//err := video.CutVideo("zamay.mp4", "output.mp4", "00:00:05", "00:00:10")
	//if err != nil {
	//	fmt.Println("Обработка ошибки")
	//}

}
