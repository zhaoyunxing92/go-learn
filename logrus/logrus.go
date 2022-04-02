package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type LogFormatter struct{}

func (s *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	//fmt.Println(entry.Data)
	msg := fmt.Sprintf("%s [%s:%d][GOID:%d][%s] %s\n", timestamp, file, len, getGID(), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
func main() {
	log.SetLevel(log.DebugLevel)
	//log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(new(LogFormatter))
	//log.SetFormatter(&log.JSONFormatter{
	//	TimestampFormat:"2006-01-02 15:04:05",
	//	//PrettyPrint: true,
	//})
	log.Info("hello")
	log.WithFields(log.Fields{
		"age":  14,
		"name": "xiaofang",
		"sex":  1,
	}).Fatal("小芳来了")

	log.Debug("debug")
}
