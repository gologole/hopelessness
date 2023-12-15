package server

import (
	"encoding/json"
	"hash/crc32"
	"io/ioutil"
	"main.go/videoprocessing"
	"net/http"
)

/*
post запрос для отправки видео
bytevideo:байтовый массив
time:строка времени в виде "00:00:05"
flag 1-обрезать 2-изменить формат 3-и то и то
hash:string
struct {  если без изменения размеров то nul поля

		newheight
		newwidth
	}
*/
type RequestStruct struct {
	userID    int
	Bytevideo []byte
	Hash      uint32
}

type Videoinfo struct {
	userID    int
	Flag      int
	Bytevideo []byte
	Starttime string
	Endtime   string
	Hash      uint32
	Newheight int
	Newwidth  int
	URL       string
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}
	// Декодирование JSON из тела запроса в структуру Videoinfo
	var video Videoinfo
	err = json.Unmarshal(body, &video)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	sum := crc32.ChecksumIEEE(video.Bytevideo)
	if video.Hash == sum {
		//отправка ответа внутри
		videoprocessing.ProcessVideo(video)
	} else {
		http.Error(w, "Хэш не сошелся", http.StatusTeapot)
	}
}
