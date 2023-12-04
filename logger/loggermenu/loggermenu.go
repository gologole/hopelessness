package loggermenu

import (
	"fmt"
	"main.go/configs"
	"main.go/logger"
	"os"
)

var logfile = configs.NameofLogfile

func ClearLogs() {
	var a string
	fmt.Println("Вы точно хотите стереть что-то из логов?y/n")
	fmt.Scan(&a)
	if !(a == "y" || a == "n") {
		ClearLogs()
	}
	if a == "n" {
		return
	}

	fmt.Println("Вы хотите очистить логи полностью?y/n")
	fmt.Scan(&a)
	if !(a == "y" || a == "n") {
		ClearLogs()
	}
	//УДАЛЕНИЕ КАКОГО ТО КОЛВА ЛОГОВ
	if a == "n" {
		return
	}

	fmt.Println("Стераю всё")
	f, err := os.Open("logs.log")
	if err != nil {
		logger.Logger.Error("ПРоблема в открытии файла с логами:", err)
	}
	defer f.Close()
	//удаление всего тек
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
}
