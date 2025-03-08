package config

import (
	"os"
	"users-api/app/src/constants"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func InitLog() {

	log.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
		//    TimestampFormat: "2006-01-02T15:04:05",
		TimestampFormat: constants.LoggerTimestampFormat,
		ShowFullLevel:   true,
		CallerFirst:     true,
	})

}

func getLoggerLevel(value string) log.Level {
	switch value {
	case constants.LoggerLevelDebug:
		return log.DebugLevel
	case constants.LoggerLevelTrace:
		return log.TraceLevel
	default:
		return log.InfoLevel
	}
}
