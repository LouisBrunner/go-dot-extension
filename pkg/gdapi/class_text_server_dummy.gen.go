// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextServerDummy struct {
  obj gdc.ObjectPtr
}

func createTextServerDummy(obj gdc.ObjectPtr) *TextServerDummy {
  return &TextServerDummy{
    obj: obj,
  }
}

func (me *TextServerDummy) BaseClass() string {
  return "TextServerDummy"
}
