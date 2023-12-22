package byte

import (
	"io/ioutil"
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
		logger.Logger.Error("проблема с переводом байтов в видео:", err)
		return
	}

	logger.Logger.Info("Видео файл успешно создан.")
}

func VideoToBytes(filename string) ([]byte, error) {
	// Чтение видеофайла в байтовый срез
	videoBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return videoBytes, nil
}
