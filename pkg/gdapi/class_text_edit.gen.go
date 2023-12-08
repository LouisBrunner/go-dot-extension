// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextEdit struct {
  obj gdc.ObjectPtr
}

func createTextEdit(obj gdc.ObjectPtr) *TextEdit {
  return &TextEdit{
    obj: obj,
  }
}

func (me *TextEdit) BaseClass() string {
  return "TextEdit"
}
