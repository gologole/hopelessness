package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// для импорта в main
var Logger *logrus.Logger

func CreateLogger() {
	// Создание нового логгера
	Logger = logrus.New()

	// Установка уровня логирования для информационных сообщений
	Logger.SetLevel(logrus.InfoLevel)

	// Настройка формата вывода логов
	Logger.SetFormatter(&logrus.TextFormatter{})

	// Настройка вывода логов в файл
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.Info("Не удалось открыть файл логов. Логирование будет осуществляться в стандартный вывод.")
	}

}
