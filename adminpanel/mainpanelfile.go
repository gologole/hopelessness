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
	instance1 = f{1, "Открыть меню логера(пока просто вывод записей) ", panelfuncs.OpenLogger}
	instance2 = f{2, "Мониторинг ресурсов и состояния-", panelfuncs.MonitorResourcesAndState}
	instance3 = f{3, "Планирование работы сервера-", panelfuncs.PlanServerWork}
	instance4 = f{4, "Включить/выключить сервер-", panelfuncs.ToggleServer}
	instance5 = f{5, "Обработать видео с замаем", panelfuncs.TestProcessing}
	instance6 = f{6, "скачать ffmpeg(один раз точно сработало)", panelfuncs.DownloadFFMPEG}
)

var flist1 = []f{instance1, instance2, instance3, instance4, instance5, instance6}

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
	fmt.Scan(&a)
	//if(a!=) проверка на валидность значения
	usefunc(a, flist1)
}

func StartMenu() {
	mainmenu()
}
