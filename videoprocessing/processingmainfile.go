package videoprocessing

import (
	byte2 "main.go/videoprocessing/byte"
	"main.go/videoprocessing/changesize"
	"main.go/videoprocessing/cutting"
	"os"
	"strconv"
)

type RequestStruct struct {
	UserID    int
	Bytevideo []byte
	Hash      uint32
}

type Videoinfo struct {
	UserID    int
	Flag      int
	Bytevideo []byte
	Starttime string
	Endtime   string
	Hash      uint32
	Newheight int
	Newwidth  int
	URL       string
}

var ScoreVideo int

func ProcessVideo(videostruct Videoinfo) (Videoinfo, string, error) {
	format := ".mp4"
	filename := "input"
	outputfile := "output"
	filename += strconv.Itoa(ScoreVideo) + format
	outputfile += strconv.Itoa(ScoreVideo) + format

	v, o, err := goProcessVideo(filename, outputfile, videostruct)
	ScoreVideo++
	return v, o, err
}

func goProcessVideo(filename string, outputfile string, videostruct Videoinfo) (Videoinfo, string, error) {

	byte2.DeserVid(filename, string(videostruct.Bytevideo))
	//это выглядит странно,но это вызов функции
	var err error
	if videostruct.Flag == 1 {
		err = CutVideo(
			filename,
			outputfile,
			videostruct.Starttime,
			videostruct.Endtime)
	}

	if videostruct.Flag == 2 {
		err = ChangeSize(
			filename,
			outputfile,
			videostruct.Newheight,
			videostruct.Newwidth,
		)
	}

	if videostruct.Flag == 3 {
		err = CutVideo(
			filename,
			outputfile,
			videostruct.Starttime,
			videostruct.Endtime)
		//изменяется размер необработанного видоса
		os.Remove(filename)
		err = ChangeSize(
			outputfile,
			filename,
			videostruct.Newheight,
			videostruct.Newwidth,
		)
		os.Remove(outputfile)
		if err != nil {
			return videostruct, filename, err
		}
		return videostruct, outputfile, nil
	}

	go os.Remove(filename)
	if err != nil {
		return videostruct, outputfile, err
	}
	return videostruct, outputfile, nil
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
