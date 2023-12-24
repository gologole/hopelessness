package adminpanel

import (
	"fmt"
	"main.go/adminpanel/panelfuncs"
)

// Чтобы добавить страницу-создатить menufunc и массив структур для неё
// Чтобы добавить функцию -добавить структуру
// добавить проверку изменений с гитхаба и версию

type f struct {
	num      int
	name     string
	function func()
}

var ( //главная страница
	instance1 = f{1, "Открыть меню логера", panelfuncs.OpenLogger}
	instance2 = f{2, "Планирование работы сервера", panelfuncs.PlanServerWork}
	instance3 = f{3, "Включить/выключить сервер и сайт", panelfuncs.ToggleServer}
	instance4 = f{4, "Тест обработки видео", panelfuncs.TestProcessing}
	instance5 = f{5, "скачать ffmpeg(один раз точно сработало)", panelfuncs.DownloadFFMPEG}
)

var flist1 = []f{instance1, instance2, instance3, instance4, instance5}

func rangef(flist []f) {
	for _, f := range flist {
		fmt.Println(f.num, ".", f.name)
	}
}

func usefunc(input int, flist []f) {
	for _, f := range flist {
		if input == f.num {
			f.function()
		}
	}
}

func mainmenu() {
	fmt.Println("Доступные функции админ-панели:")
	rangef(flist1)
	var a int
	b, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("ВВедите валидное значение", b)
		mainmenu()
	}
	usefunc(a, flist1)
}

func StartMenu() {
	mainmenu()
}
