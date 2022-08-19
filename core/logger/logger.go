package logger

import "fmt"

const colorReset = "\033[0m"

const colorRed = "\033[31m"
const colorCyan = "\033[36m"
const colorGreen = "\033[32m"

// const colorYellow = "\033[33m"
// const colorBlue = "\033[34m"
// const colorPurple = "\033[35m"
// const colorWhite = "\033[37m"

func Error(format string, n ...any) {
	fmt.Print(colorRed)
	fmt.Printf(format, n...)
	fmt.Print(colorReset)
}

func Verbose(format string, n ...any) {
	fmt.Print(colorCyan)
	fmt.Printf(format, n...)
	fmt.Print(colorReset)
}

func Info(format string, n ...any) {
	fmt.Print(colorGreen)
	fmt.Printf(format, n...)
	fmt.Print(colorReset)
}
