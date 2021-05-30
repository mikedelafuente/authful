package logger

import (
	"log"
)

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}
