// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileDialog struct {
  obj gdc.ObjectPtr
}

func createFileDialog(obj gdc.ObjectPtr) *FileDialog {
  return &FileDialog{
    obj: obj,
  }
}

func (me *FileDialog) BaseClass() string {
  return "FileDialog"
}
