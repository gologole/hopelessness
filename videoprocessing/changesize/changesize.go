package changesize

import (
	"github.com/mowshon/moviego"
	"main.go/logger"
)

func ChangeSize(inputfile string, outputfile string, width int, height int) error {
	first, err := moviego.Load(inputfile)
	if err != nil {
		logger.Logger.Error("Ошибка при загрузке видео для изменения размера", err)
	}
	err = first.Resize(int64(width), int64(height)).Output(outputfile).Run()
	if err != nil {
		logger.Logger.Error("Ошибка в RUn", err)
		return err
	}
	logger.Logger.Info("Изменен размер видео", inputfile)
	return nil
}
