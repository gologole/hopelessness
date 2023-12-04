package server

import (
	"fmt"
	"main.go/logger"
	"main.go/videoprocessing"
	"net/http"
)

/*
post запрос для отправки видео
bytevideo:байтовый массив
time:строка времени в виде "00:00:05"
hash:string
struct {  если без изменения размеров то nul поля

		newheight
		newwidth
	}
*/
type Sender struct {
}

type Videoinfo struct {
	Bytevideo []byte
	Starttime string
	Endtime   string
	hash      string
	Newheight int
	Newwidth  int
}

func StartServer() {

	mux := http.NewServeMux()

	mux.HandleFunc("/catchvideo", gocatchvideo)
	//	mux.HandleFunc("/sendlogs", sendlogs)

	http.ListenAndServe(":8080", mux)
}

func gocatchvideo(w http.ResponseWriter, r *http.Request) {
	go catchvideo(w, r)
}

func catchvideo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			logger.Logger.Error("Ошибка при парсинге формы постзапроса:", err)
		}
		video := r.Form.Get("bytevideo")
		//сделать возврат ошибок и слушатель окончание обработки и перевод тела запроса в структуру
		videoprocessing.ProcessVideo(video)

	} else {
		fmt.Println("Ктото пытается отправить видео гетзапросом,но зачем?")
	}
}
