package logger

import (
	"os"
	"service/internal/config"

	"github.com/sirupsen/logrus"
)

func Init(cfg config.Config) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})

	switch cfg.GetLogLevel() {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetOutput(os.Stdout)
	// switch cfg.GetLogOut() {
	// case "":
	// 	log.SetOutput(os.Stdout)
	// default:
	// 	log.SetOutput(os.Stdout) //TODO: not yet
	// }

	return log
}
