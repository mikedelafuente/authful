package logger

import (
	"fmt"
	"log"
)

func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v)
	log.Printf(format, v)
}

func Fatal(v ...interface{}) {
	fmt.Println(v)
	log.Fatal(v)
}

func Println(v ...interface{}) {
	fmt.Println(v)
	log.Println(v)
}
