// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextServerExtension struct {
  obj gdc.ObjectPtr
}

func createTextServerExtension(obj gdc.ObjectPtr) *TextServerExtension {
  return &TextServerExtension{
    obj: obj,
  }
}

func (me *TextServerExtension) BaseClass() string {
  return "TextServerExtension"
}
