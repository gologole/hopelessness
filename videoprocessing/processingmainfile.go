package videoprocessing

import (
	"fmt"
	byte2 "main.go/videoprocessing/byte"
	"main.go/videoprocessing/changesize"
	"main.go/videoprocessing/cutting"
	"os"
	"strconv"
)

type RequestStruct struct {
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

var ScoreVideo int = 1

func ProcessVideo(videostruct Videoinfo) (Videoinfo, string) {
	format := ".mp4"
	filename := "input"
	outputfile := "output"
	filename += strconv.Itoa(ScoreVideo) + format
	outputfile += strconv.Itoa(ScoreVideo) + format

	v, o := goProcessVideo(filename, outputfile, videostruct)
	ScoreVideo++
	return v, o
}

func goProcessVideo(filename string, outputfile string, videostruct Videoinfo) (Videoinfo, string) {
	byte2.DeserVid(filename, string(videostruct.Bytevideo))
	//это выглядит странно,но это вызов функции
	var err error
	if videostruct.Flag == 1 {
		err = cutting.CutFile(
			filename,
			outputfile,
			videostruct.Starttime,
			videostruct.Endtime)
	}

	if videostruct.Flag == 2 {
		err = changesize.ChangeSize(
			filename,
			outputfile,
			videostruct.Newheight,
			videostruct.Newwidth,
		)
	}

	if videostruct.Flag == 3 {
		theend := "TheEnd" + strconv.Itoa(ScoreVideo) + ".mp4"
		printVideoinfo(videostruct)
		err = cutting.CutFile(filename, outputfile, videostruct.Starttime, videostruct.Endtime)
		os.Remove(filename)
		err = changesize.ChangeSize(outputfile, theend, videostruct.Newheight, videostruct.Newwidth)
		//os.Remove(outputfile)
		if err != nil {
			return videostruct, filename
		}
		return videostruct, theend
	}

	//go os.Remove(filename)
	if err != nil {
		return videostruct, outputfile
	}
	return videostruct, outputfile
}

func printVideoinfo(vi Videoinfo) {
	fmt.Println("UserID:", vi.UserID)
	fmt.Println("Flag:", vi.Flag)
	fmt.Println("Starttime:", vi.Starttime)
	fmt.Println("Endtime:", vi.Endtime)
	fmt.Println("Hash:", vi.Hash)
	fmt.Println("Newheight:", vi.Newheight)
	fmt.Println("Newwidth:", vi.Newwidth)
	fmt.Println("URL:", vi.URL)
}
