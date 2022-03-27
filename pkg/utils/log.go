package utils

import (
	log "github.com/sirupsen/logrus"
)

func LogFormat(layer, service string, message interface{}) log.Fields {

	return log.Fields{
		"layer":   layer,
		"service": service,
		"message": message,
	}

}

// func NewLog() *log.Logger {
// 	log := log.New()

// 	// / You could set this to any `io.Writer` such as a file
// 	file, _ := os.OpenFile("../logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

// 	log.SetOutput(file)

// 	return log

// }
