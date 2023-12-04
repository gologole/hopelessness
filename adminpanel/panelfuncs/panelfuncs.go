package panelfuncs

import (
	"fmt"
	"main.go/configs/setconfigs"
	"main.go/logger/loggermenu"
	"main.go/videoprocessing"
	"time"
)

func OpenLogger() {
	loggermenu.OpenLogger()
}

func MonitorResourcesAndState() {
	fmt.Println("Потрачено ресурсов:")
}

func PlanServerWork() {
	fmt.Println("Как спланируем работу сервера")
}

func ToggleServer() {
	fmt.Println("включить/выключить модуль/сервер")
}
func TestProcessing() {
	fmt.Println("введите с какой по какую секунду хотите обрезать видео ,где замай благославляет валю карнавал")
	var a, b int
	fmt.Scan(&a, &b)
	a1 := secondsToTimeString(a)
	b1 := secondsToTimeString(b)
	err := videoprocessing.CutVideo("zamay.mp4", "output.mp4", a1, b1)
	if err != nil {
		fmt.Println("Обработка ошибки")
	}

	fmt.Println("введите в какое качество вы хотите зашакалить видео ,где замай благославляет валю карнавал ")
	fmt.Scan(&a, &b)

	err = videoprocessing.ChangeSize("output.mp4", "theEnd.mp4", a, b)
	if err != nil {
		fmt.Println("Обработка ошибки")
	}
	fmt.Println("сейчас должен появится файл TheEnd.mp4")
}

// надо сделать обработку ошибок и добавить функцию в админ панель
func DownloadFFMPEG() {
	fmt.Println("Введите путь к файлу с скриптом в таком формате:C:\\Users\\nikit\\Desktop\\ИС\\ffmpeg_install.ps1")
	var a string
	fmt.Scan(&a)
	setconfigs.SetPATH(a)
}

func secondsToTimeString(seconds int) string {
	duration := time.Second * time.Duration(seconds)
	parsedTime, _ := time.Parse("15:04:05", "00:00:00")
	timeToAdd := parsedTime.Add(duration)
	result := timeToAdd.Format("15:04:05")
	return result
}
