package glog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Log struct {
	dl *log.Logger
	il *log.Logger
	wl *log.Logger
	el *log.Logger
}

func NewLogger(filename string) *Log {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file '", filename, "':", err)
	}

	l := new(Log)
	l.dl = log.New(io.MultiWriter(file, ioutil.Discard), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.il = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.wl = log.New(io.MultiWriter(file, os.Stdout), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.el = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return l
}

func (l *Log) Debugf(format string, a ...interface{}) {
	l.dl.Printf(format, a...)
}

func (l *Log) Debug(line string) {
	l.dl.Println(line)
}

func (l *Log) Infof(format string, a ...interface{}) {
	l.il.Printf(format, a...)
}

func (l *Log) Info(line string) {
	l.il.Println(line)
}

func (l *Log) Warnf(format string, a ...interface{}) {
	l.wl.Printf(format, a...)
}

func (l *Log) Warn(line string) {
	l.wl.Println(line)
}

func (l *Log) Errorf(format string, a ...interface{}) {
	l.el.Printf(format, a...)
}

func (l *Log) Error(line string) {
	l.el.Println(line)
}
