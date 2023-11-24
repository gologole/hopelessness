package cutting

import (
	"main.go/logger"
	"os/exec"
)

func cropVideo(inputFile string, outputFile string, startTime string, duration string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-ss", startTime, "-t", duration, "-c", "copy", outputFile)
	err := cmd.Run()
	if err != nil {
		logger.Logger.Error("Error cropping video in cropVideo:", err)
		return err
	}
	return nil
}

func CutFile(inputFile, outputFile string, startTime, duration string) error {
	logger.Logger.Info("Запущена обрезка видео.Название файла:", inputFile)
	//inputFile := "zamay.mp4"
	//outputFile := "output.mp4"
	//startTime := "00:00:05" // Начиная с 5-й секунды
	//duration := "00:00:10"  // Длительность обрезки - 10 секунд

	err := cropVideo(inputFile, outputFile, startTime, duration)
	if err != nil {
		logger.Logger.Error("Cut.go 27:Error cropping video:", err)
		return err
	}
	logger.Logger.Info("Video cropped successfully")
	return nil
}
