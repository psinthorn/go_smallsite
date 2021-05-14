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
	StringToTime(s1, s2 string) (time.Time, time.Time, error)
}
