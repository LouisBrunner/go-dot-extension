// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ColorPicker struct {
  obj gdc.ObjectPtr
}

func createColorPicker(obj gdc.ObjectPtr) *ColorPicker {
  return &ColorPicker{
    obj: obj,
  }
}

func (me *ColorPicker) BaseClass() string {
  return "ColorPicker"
}
