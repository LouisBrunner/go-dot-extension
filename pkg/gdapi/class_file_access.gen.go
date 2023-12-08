// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileAccess struct {
  obj gdc.ObjectPtr
}

func (me *FileAccess) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FileAccess) BaseClass() string {
  return "FileAccess"
}

type FileAccessModeFlags int
const (
  FileAccessRead FileAccessModeFlags = 1
  FileAccessWrite FileAccessModeFlags = 2
  FileAccessReadWrite FileAccessModeFlags = 3
  FileAccessWriteRead FileAccessModeFlags = 7
)

type FileAccessCompressionMode int
const (
  FileAccessCompressionFastlz FileAccessCompressionMode = 0
  FileAccessCompressionDeflate FileAccessCompressionMode = 1
  FileAccessCompressionZstd FileAccessCompressionMode = 2
  FileAccessCompressionGzip FileAccessCompressionMode = 3
  FileAccessCompressionBrotli FileAccessCompressionMode = 4
)
