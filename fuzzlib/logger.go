package fuzzlib

import (
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	debugLog *log.Logger
	debug    bool
}

func NewLogger() Logger {
	return Logger{
		errorLog: log.New(os.Stdout, color.New(color.FgRed).Sprint("ERROR "), 0),
		infoLog:  log.New(os.Stdout, color.New(color.FgWhite).Sprint("INFO "), 0),
		debugLog: log.New(os.Stdout, color.New(color.FgHiBlue).Sprint("DEBUG "), 0),
	}
}

func (l Logger) Error(v ...any) {
	l.errorLog.Print(v...)
}

func (l Logger) Errorf(format string, v ...any) {
	l.errorLog.Printf(format, v...)
}

func (l Logger) Info(v ...any) {
	l.infoLog.Print(v...)
}

func (l Logger) Infof(format string, v ...any) {
	l.infoLog.Printf(format, v...)
}

func (l Logger) Debug(v ...any) {
	if !l.debug {
		return
	}
	l.debugLog.Print(v...)
}

func (l Logger) Debugf(format string, v ...any) {
	if !l.debug {
		return
	}
	l.debugLog.Printf(format, v...)
}
