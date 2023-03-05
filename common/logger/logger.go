package logger

import (
	"log"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *log.Logger
)

func init() {
	once.Do(func() {
		logger = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

// Info ...
func Info(args ...interface{}) {
	logger.SetPrefix("[INFO]")
	logger.Println(args...)
}

// Error  ...
func Error(args ...interface{}) {
	logger.SetPrefix("[ERROR]")
	logger.Fatal(args...)
}

// Warning ...
func Warning(args ...interface{}) {
	logger.SetPrefix("[WARNING]")
	logger.Println(args...)
}

// DeBug ...
func DeBug(args ...interface{}) {
	logger.SetPrefix("[DeBug]")
	logger.Println(args...)
}
