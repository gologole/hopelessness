package setconfigs

import (
	"fmt"
	"os"
	"os/exec"
)

var scriptPath = "C:\\Users\\Никита\\Desktop\\remakeProj\\script.ps1" //засунуть в конфигурации

func SetPATH(scriptpath1 string) {
	// Путь к скрипту ffmpeg_install.ps1

	if scriptpath1 == "-----" {
		// Создание команды powershell для запуска скрипта
		cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "-File", scriptPath)

		// Установка stdout и stderr для вывода результатов выполнения скрипта
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Запуск команды
		err := cmd.Run()
		if err != nil {
			fmt.Println("Ошибка в запуске скрипта автоустановки ffmpeg и PATH:", err)
		}
	} else {
		// Создание команды powershell для запуска скрипта
		cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "-File", scriptpath1)

		// Установка stdout и stderr для вывода результатов выполнения скрипта
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Запуск команды
		err := cmd.Run()
		if err != nil {
			fmt.Println("Ошибка в запуске скрипта автоустановки ffmpeg и PATH:", err)
		}
	}
}

//
//func main() {
//	fileName := "example.txt"
//	startDir := "."  // Начальная директория для поиска (можно указать путь)
//
//	filePath, err := findFile(fileName, startDir)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("Путь к файлу:", filePath)
//	}
//}
