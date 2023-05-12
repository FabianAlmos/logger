package logger

import "fmt"

func Info(data any) {
	fmt.Println("Log info: ", data)
}

func Error(data any) {
	fmt.Println("Log error: ", data)
}

func Warn(data any) {
	fmt.Println("Log warn: ", data)
}
