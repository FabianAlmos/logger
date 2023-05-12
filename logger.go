package logger

import "fmt"

func Info(message string, params any) {
	fmt.Println("Log info message: ", message, "params: ", params)
}

func Error(message string, params any) {
	fmt.Println("Log error message: ", message, "params: ", params)
}

func Warn(message string, params any) {
	fmt.Println("Log warn message: ", message, "params: ", params)
}

func Fatal(message string, params any) {
	fmt.Println("Log fatal message: ", message, "params: ", params)
}
