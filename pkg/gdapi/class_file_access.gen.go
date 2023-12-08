// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileAccess struct {
  obj gdc.ObjectPtr
}

func createFileAccess(obj gdc.ObjectPtr) *FileAccess {
  return &FileAccess{
    obj: obj,
  }
}

func (me *FileAccess) BaseClass() string {
  return "FileAccess"
}
