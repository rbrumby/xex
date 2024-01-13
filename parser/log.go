package parser

import (
	"log"
	"os"
)

var (
	logInf *log.Logger
	logErr *log.Logger
)

func init() {
	logInf = log.New(os.Stdout, "INFO: ", log.Llongfile)
	logErr = log.New(os.Stdout, "ERROR: ", log.Llongfile)
}
