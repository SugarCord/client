package errorHandling

import "log"

func FatalCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogCheck(err error) {
	if err != nil {
		log.Print(err)
	}
}
