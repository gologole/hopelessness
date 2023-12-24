package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"main.go/logger"
	"main.go/videoprocessing"
	byte2 "main.go/videoprocessing/byte"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var finalfile string

type Videoinfo = videoprocessing.Videoinfo
type RequestStruct = videoprocessing.RequestStruct

var mux = http.NewServeMux()

func StartServer() {
	//mux := http.NewServeMux()
	mux.HandleFunc("/site", indexHandler)
	mux.HandleFunc("/process", processVideo)
	mux.HandleFunc("/catchvideo", gocatchvideo)
	mux.HandleFunc("/video", videoHandler)
	mux.HandleFunc("/sendlogs", sendlogs)
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
	sum := CalculateCrc32(video.Bytevideo)
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

	arrbyte, err := byte2.VideoToBytes(filename)
	a := len(arrbyte) < 5
	logger.Logger.Info("массив видео пустой = ", a)
	if err != nil {
		logger.Logger.Error("Проблема перевода видео в байты", err)
	}
	sb.Bytevideo = arrbyte

	sb.Hash = CalculateCrc32(arrbyte)

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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func processVideo(w http.ResponseWriter, r *http.Request) {
	// Получить файл видео из формы
	var videostruct videoprocessing.Videoinfo
	file, header, err := r.FormFile("video")
	Starttime := r.FormValue("startTime")
	Endtime := r.FormValue("endTime")
	resolution := r.FormValue("resolution")
	if err != nil {
		http.Error(w, "Не удалось", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Создать файл для сохранения видео на сервере
	out, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, "Не удалось создать файл для сохранения видео", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Скопировать содержимое файла видео в созданный файл на сервере
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Не удалось сохранить видео на сервере", http.StatusInternalServerError)
		return
	}
	Starttim, _ := strconv.Atoi(Starttime)
	Endtim, _ := strconv.Atoi(Endtime)
	videostruct.Starttime = secondsToTimeString(Starttim)
	videostruct.Endtime = secondsToTimeString(Endtim)
	parts := strings.Split(resolution, " ")
	width, _ := strconv.Atoi(parts[0])
	height, _ := strconv.Atoi(parts[1])
	videostruct.Newwidth = width
	videostruct.Newheight = height
	videostruct.Flag = 3

	arr, _ := byte2.VideoToBytes(header.Filename)
	videostruct.Bytevideo = arr
	fmt.Println(videostruct)

	_, finalfile, err = videoprocessing.ProcessVideo(videostruct)
	w.WriteHeader(200)
	go DeleteFile(finalfile, file)

}

func secondsToTimeString(seconds int) string {
	duration := time.Second * time.Duration(seconds)
	parsedTime, _ := time.Parse("15:04:05", "00:00:00")
	timeToAdd := parsedTime.Add(duration)
	result := timeToAdd.Format("15:04:05")
	return result
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	// Откройте файл видео
	file, err := os.Open(finalfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("attachment; filename=video.mp4", "attachment; filename=video.mp4")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Отправьте содержимое файла в ответ
	io.Copy(w, file)
}

func DeleteFile(file string, fileToClose io.Closer) {
	time.Sleep(30 * time.Minute)
	fileToClose.Close()
	os.Remove(file)
}

func sendlogs(w http.ResponseWriter, r *http.Request) {
	// Открываем файл logs.log
	file, err := ioutil.ReadFile("logs.log")
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка чтения файла: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Отправляем содержимое файла в запросе
	w.Write(file)
}
func CalculateCrc32(data []byte) uint32 {
	crc := crc32.NewIEEE()
	crc.Write(data)
	return crc.Sum32()
}
