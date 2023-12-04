package byte

import (
	"main.go/logger"
	"os"
)

func DeserVid(filename, data string) {
	// Открытие файла для записи
	file, err := os.Create(filename)
	if err != nil {
		logger.Logger.Error("трабл с десерелизацией", err)
		return
	}
	defer file.Close()

	bytes := []byte(data)

	// Запись данных в файл
	_, err = file.Write(bytes)
	if err != nil {
		logger.Logger.Error("проблема с десерелизацией:", err)
		return
	}

	logger.Logger.Info("Видео файл успешно создан.")
}

//qwerty
