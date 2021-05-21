package utils

import (
	"time"

	"github.com/psinthorn/go_smallsite/configs"
)

var UtilsService utilsInterface = &Utils{}

type Utils utils
type utils struct{}

type utilsInterface interface {
	IsProduction(appConfig *configs.AppConfig)
	StringToTime(timeString string) (time.Time, error)
}
