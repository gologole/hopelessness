package videoprocessing

import (
	"main.go/videoprocessing/changesize"
	"main.go/videoprocessing/cutting"
)

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
