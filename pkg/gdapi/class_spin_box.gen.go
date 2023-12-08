// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpinBox struct {
  obj gdc.ObjectPtr
}

func createSpinBox(obj gdc.ObjectPtr) *SpinBox {
  return &SpinBox{
    obj: obj,
  }
}

func (me *SpinBox) BaseClass() string {
  return "SpinBox"
}
