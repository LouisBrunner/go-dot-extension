// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextLine struct {
  obj gdc.ObjectPtr
}

func createTextLine(obj gdc.ObjectPtr) *TextLine {
  return &TextLine{
    obj: obj,
  }
}

func (me *TextLine) BaseClass() string {
  return "TextLine"
}
