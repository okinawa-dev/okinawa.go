package log

import (
	"log"
)

// Initialize prepares the log component of the engine
func Initialize() {
}

// Log prints a log message to the console
func Log(fileName string, message string) {
	log.Println("[" + fileName + "] " + message)
}
