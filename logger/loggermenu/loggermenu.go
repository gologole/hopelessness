package loggermenu

import (
	"fmt"
	"main.go/configs"
	"os"
)

var logfile = configs.NameofLogfile

func OpenLogger() {
	fmt.Println("оно случилось")
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
