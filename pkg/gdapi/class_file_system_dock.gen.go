// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileSystemDock struct {
  obj gdc.ObjectPtr
}

func createFileSystemDock(obj gdc.ObjectPtr) *FileSystemDock {
  return &FileSystemDock{
    obj: obj,
  }
}

func (me *FileSystemDock) BaseClass() string {
  return "FileSystemDock"
}
