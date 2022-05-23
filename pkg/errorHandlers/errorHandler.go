package errorHandlers

import (
	"log"
)

func PrintError(msg string, err error) {
	if err == nil {
		return
	}
	log.Printf("Message: %s, Error: %v\n", msg, err)
}

func PanicError(msg string, err error) {
	if err == nil {
		return
	}
	log.Panicf("Message: %s, Error: %v\n", msg, err)
}

func FatalError(msg string, err error) {
	if err == nil {
		return
	}
	log.Fatalf("Message: %s, Error: %v\n", msg, err)
}
