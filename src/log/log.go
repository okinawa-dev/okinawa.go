package log

import (
	"fmt"
	"log"
)

// Stringer is an interface for types that can be converted to a string
// type Stringer interface {
// 	String() string
// }

// Initialize prepares the log component of the engine
func Initialize() {
}

// toString converts multiple types to string
func toString(message any) string {
	switch message.(type) {
	case fmt.Stringer:
		return message.(fmt.Stringer).String()
	case string:
		return message.(string)
	default:
		return "<something strange here>"
	}
}

// Info prints a log message to the console
func Info(location string, messages ...any) {

	result := "[" + location + "] "

	for _, msg := range messages {
		result += toString(msg) + " "
	}

	log.Println(result)
}
