// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ColorPickerButton struct {
  obj gdc.ObjectPtr
}

func createColorPickerButton(obj gdc.ObjectPtr) *ColorPickerButton {
  return &ColorPickerButton{
    obj: obj,
  }
}

func (me *ColorPickerButton) BaseClass() string {
  return "ColorPickerButton"
}
