package Frontend

//
//import (
//	"fmt"
//	"io"
//	"main.go/videoprocessing"
//	byte2 "main.go/videoprocessing/byte"
//	"net/http"
//	"os"
//	"strconv"
//	"strings"
//	"time"
//)
//
//func StartFrontend() {
//	http.HandleFunc("/", indexHandler)
//	http.HandleFunc("/process", processVideo)
//	http.ListenAndServe(":8080", nil)
//}
//
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	http.ServeFile(w, r, "index.html")
//}
//
//func processVideo(w http.ResponseWriter, r *http.Request) {
//	// Получить файл видео из формы
//	var videostruct videoprocessing.Videoinfo
//	file, header, err := r.FormFile("video")
//	Starttime := r.FormValue("startTime")
//	Endtime := r.FormValue("endTime")
//	resolution := r.FormValue("resolution")
//	if err != nil {
//		http.Error(w, "Не удалось", http.StatusBadRequest)
//		return
//	}
//	defer file.Close()
//
//	// Создать файл для сохранения видео на сервере
//	out, err := os.Create(header.Filename)
//	if err != nil {
//		http.Error(w, "Не удалось создать файл для сохранения видео", http.StatusInternalServerError)
//		return
//	}
//	defer out.Close()
//
//	// Скопировать содержимое файла видео в созданный файл на сервере
//	_, err = io.Copy(out, file)
//	if err != nil {
//		http.Error(w, "Не удалось сохранить видео на сервере", http.StatusInternalServerError)
//		return
//	}
//	Starttim, _ := strconv.Atoi(Starttime)
//	Endtim, _ := strconv.Atoi(Endtime)
//	videostruct.Starttime = secondsToTimeString(Starttim)
//	videostruct.Endtime = secondsToTimeString(Endtim)
//	parts := strings.Split(resolution, " ")
//	width, _ := strconv.Atoi(parts[0])
//	height, _ := strconv.Atoi(parts[1])
//	videostruct.Newwidth = width
//	videostruct.Newheight = height
//	videostruct.Flag = 3
//	fmt.Println(videostruct)
//	arr, _ := byte2.VideoToBytes(header.Filename)
//	videostruct.Bytevideo = arr
//	_, finalfile, err := videoprocessing.ProcessVideo(videostruct)
//	videoHandler(w, r, finalfile)
//	os.Remove(header.Filename)
//}
//
//func secondsToTimeString(seconds int) string {
//	duration := time.Second * time.Duration(seconds)
//	parsedTime, _ := time.Parse("15:04:05", "00:00:00")
//	timeToAdd := parsedTime.Add(duration)
//	result := timeToAdd.Format("15:04:05")
//	return result
//}
//
//func videoHandler(w http.ResponseWriter, r *http.Request, filePath string) {
//	// Открываем видеофайл
//
//	file, err := os.Open(filePath)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	defer file.Close()
//
//	// Получаем информацию о файле
//	fileInfo, err := file.Stat()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// Устанавливаем заголовки ответа
//	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
//	w.Header().Set("Content-Type", "video/mp4")
//	w.Header().Set("Content-Length", string(fileInfo.Size()))
//
//	// Отправляем содержимое видеофайла обратно клиенту
//	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
//}
