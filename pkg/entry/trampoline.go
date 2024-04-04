package entry

import (
	"fmt"
	"log"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdextension"
)

var gext gdextension.Extension

func Logf(level gdextension.LogLevel, format string, v ...interface{}) {
	if gext == nil {
		log.Printf("error: extension not initialized")
		log.Printf(format, v...)
		return
	}
	gext.Logf(level, format, v...)
}

func CreateClass[T gdextension.Class]() (*T, error) {
	if gext == nil {
		return nil, fmt.Errorf("extension not initialized")
	}
	return gdextension.CreateClass[T](gext)
}
