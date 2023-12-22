package panelfuncs

import (
	"fmt"
	"main.go/configs/setconfigs"
	"main.go/logger/loggermenu"
	"main.go/server"
	"time"
)

func OpenLogger() {
	loggermenu.LoggerMenu()

}

func MonitorResourcesAndState() {
	fmt.Println("Не работает")
}

func PlanServerWork() {
	fmt.Println("не работает")
}

func ToggleServer() {
	server.StartServer()
}
func TestProcessing() {
	server.TestProcess()
}

// надо сделать обработку ошибок и добавить функцию в админ панель
func DownloadFFMPEG() {
	fmt.Println("Введите путь к файлу с скриптом в таком формате:C:\\Users\\nikit\\Desktop\\ИС\\ffmpeg_install.ps1")
	var a string
	fmt.Scan(&a)
	setconfigs.SetPATH(a)
	fmt.Scan(&a)
}

// хз поч это здесь
func secondsToTimeString(seconds int) string {
	duration := time.Second * time.Duration(seconds)
	parsedTime, _ := time.Parse("15:04:05", "00:00:00")
	timeToAdd := parsedTime.Add(duration)
	result := timeToAdd.Format("15:04:05")
	return result
}
