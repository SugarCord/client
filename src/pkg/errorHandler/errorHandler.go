package errorHandler

import (
	// Standard Library
	"log"
)

// if err, fatal
func FatalCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// if err, log
func LogCheck(err error) {
	if err != nil {
		log.Print(err)
	}
}
