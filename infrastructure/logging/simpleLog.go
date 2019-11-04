package logging

import (
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

// init logs configuration
func init() {
	Debug = log.New(os.Stdout, "DEBUG : ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "INFO  : ",
		log.Ltime)

	Error = log.New(os.Stderr, "ERROR : ",
		log.Ldate|log.Ltime|log.Lshortfile)
}