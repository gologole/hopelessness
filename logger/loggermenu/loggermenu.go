package loggermenu

import (
	"fmt"
	"log"
	"main.go/configs"
	"os"
	"os/exec"
)

var logfile = configs.NameofLogfile
var a int

func LoggerMenu() {
	fmt.Println("введите 1 если хотите открыть логи в файле(не работает)" +
		"введите 2 если хотите вывести логи в консоль" +
		"введите 3 если хотите очистить логи")

	fmt.Scan(&a)
	if a == 1 {
		OpenLogFile()
	}
	if a == 2 {
		OpenLogger()
		fmt.Scan()
	}
	if a == 3 {
		err := ClearLogs("logs.log")
		if err != nil {
			fmt.Println("Кто-то удалил файл логов")
		}
	}
}

func OpenLogFile() {
	filePath := "logs.log"

	// Открываем файл с использованием программы по умолчанию для просмотра файла
	err := exec.Command("open", filePath).Start()

	if err != nil {
		log.Fatal(err)
	}
}

func ClearLogs(filename string) error {
	// Открываем файл для чтения
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Очищаем содержимое файла
	if err := file.Truncate(0); err != nil {
		return err
	}

	// Перемещаем указатель в начало файла
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	// Записываем пустую строку в файл
	if _, err := file.WriteString(""); err != nil {
		return err
	}

	return nil
}

func OpenLogger() {
	file, err := os.Open(logfile)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	stat, _ := file.Stat()
	fileSize := stat.Size()
	data := make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println(string(data))
	fmt.Scan(&a)
}
