package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"main.go/logger"
	"main.go/videoprocessing"
	byte2 "main.go/videoprocessing/byte"
	"net/http"
	"os"
	"time"
)

type Videoinfo = videoprocessing.Videoinfo
type RequestStruct = videoprocessing.RequestStruct

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/catchvideo", gocatchvideo)
	//	mux.HandleFunc("/sendlogs", sendlogs)
	fmt.Println("Сервер запущен")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Logger.Error("Не получается запустить сервер : ", err)
	}

}

func gocatchvideo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) //УСТАРЕЛ????????????
	if err != nil {
		logger.Logger.Error("ошибка чтения бади ", err)
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}
	// Декодирование JSON из тела запроса в структуру Videoinfo
	var video Videoinfo
	fmt.Println(string(body))
	err = json.Unmarshal(body, &video)
	if err != nil {
		logger.Logger.Error("ошибка декодирования бади :", err)
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}
	logger.Logger.Info("Принят запрос : ", video)
	sum := crc32.ChecksumIEEE(video.Bytevideo)
	if video.Hash == sum {
		//отправка ответа внутри
		logger.Logger.Info("хэш суммы сошлись")
		videostruct, outputfile, _ := videoprocessing.ProcessVideo(video)

		SendReq(videostruct, outputfile)
		err := os.Remove(outputfile)
		if err != nil {
			fmt.Println("НЕ УДАЛОСЬ УДАЛИТЬ ФАЙЛ ПОСЛЕ ОТПРАВКИ")
			logger.Logger.Error("НЕ УДАЛОСЬ УДАЛИТЬ ФАЙЛ ПОСЛЕ ОТПРАВКИ ", err)
		}

	} else {
		http.Error(w, "Хэш не сошелся", http.StatusTeapot)
		logger.Logger.Error("Хэш суммы не сошлись")
	}
}

func SendReq(video Videoinfo, filename string) {
	client := &http.Client{}

	//составления структуры ответа
	var sb RequestStruct
	sb.UserID = video.UserID

	arrbyte, err := byte2.VideoToBytes(filename)
	a := len(arrbyte) < 5
	logger.Logger.Info("массив видео пустой = ", a)
	if err != nil {
		logger.Logger.Error("Проблема перевода видео в байты", err)
	}
	sb.Bytevideo = arrbyte

	sb.Hash = crc32.ChecksumIEEE(arrbyte)

	data, err := json.Marshal(sb)
	r := bytes.NewReader(data) //хз поч он требует ридер ,а не пустой интерфейс
	logger.Logger.Info("Отправляю POST запрос :", sb)
	req, err := http.NewRequest("POST", video.URL, r)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Error("server.go ошибка в client.Do(req): ", err)
		return
	}
	defer resp.Body.Close()
}

func TestProcess() {
	fmt.Println("введите с какой по какую секунду хотите обрезать видео ,где замай благославляет валю карнавал")
	var a, b int
	fmt.Scan(&a, &b)
	a1 := secondsToTimeString(a)
	b1 := secondsToTimeString(b)
	err := videoprocessing.CutVideo("zamay.mp4", "output.mp4", a1, b1)
	if err != nil {
		fmt.Println("Обработка ошибки")
	}

	fmt.Println("введите в какое качество вы хотите зашакалить видео ,где замай благославляет валю карнавал ")
	fmt.Scan(&a, &b)

	err = videoprocessing.ChangeSize("output.mp4", "theEnd.mp4", a, b)
	if err != nil {
		fmt.Println("Обработка ошибки")
	}
	fmt.Println("сейчас должен появится файл TheEnd.mp4")
}
func secondsToTimeString(seconds int) string {
	duration := time.Second * time.Duration(seconds)
	parsedTime, _ := time.Parse("15:04:05", "00:00:00")
	timeToAdd := parsedTime.Add(duration)
	result := timeToAdd.Format("15:04:05")
	return result
}
