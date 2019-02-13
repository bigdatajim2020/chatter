package logger

import (
	"io"
	"log"
	"os"
)

var (
	// Trace logger for debugging use.
	Trace *log.Logger
	// Info logger logs infomation.
	Info *log.Logger
	// Warning logger gives alerts.
	Warning *log.Logger
	// Error logger logs errors to both file and stderr.
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open: %v", err)
	}

	Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Llongfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, file), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
