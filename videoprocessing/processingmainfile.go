package videoprocessing

import (
	"main.go/server"
	byte2 "main.go/videoprocessing/byte"
	"main.go/videoprocessing/changesize"
	"main.go/videoprocessing/cutting"
	"strconv"
)

var ScoreVideo int

func ProcessVideo(videostruct server.Videoinfo) {

	format := ".mp4"
	filename := "input"
	outputfile := "output"
	filename += strconv.Itoa(ScoreVideo) + format
	outputfile += strconv.Itoa(ScoreVideo) + format
	go goProcessVideo(filename, outputfile, videostruct)
	ScoreVideo++
}

func goProcessVideo(filename string, outputfile string, videostruct server.Videoinfo) {
	byte2.DeserVid(filename, string(videostruct.Bytevideo))
	//это выглядит странно,но это вызов функции
	CutVideo(
		filename,
		outputfile,
		videostruct.Starttime,
		videostruct.Endtime,
	)
	ChangeSize(
		filename,
		outputfile,
		videostruct.Newheight,
		videostruct.Newwidth,
	)
}

func CutVideo(inputFile, outputFile string, startTime, duration string) error {
	er := cutting.CutFile(inputFile, outputFile, startTime, duration)
	if er != nil {
		return er
	}
	return nil
}

func ChangeSize(inputfile string, outputfile string, width int, height int) error {
	er := changesize.ChangeSize(inputfile, outputfile, width, height)
	if er != nil {
		return er
	}
	return nil
}
