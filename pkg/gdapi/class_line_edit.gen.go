// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LineEdit struct {
  obj gdc.ObjectPtr
}

func createLineEdit(obj gdc.ObjectPtr) *LineEdit {
  return &LineEdit{
    obj: obj,
  }
}

func (me *LineEdit) BaseClass() string {
  return "LineEdit"
}
