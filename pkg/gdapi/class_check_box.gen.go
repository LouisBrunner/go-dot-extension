// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CheckBox struct {
  obj gdc.ObjectPtr
}

func createCheckBox(obj gdc.ObjectPtr) *CheckBox {
  return &CheckBox{
    obj: obj,
  }
}

func (me *CheckBox) BaseClass() string {
  return "CheckBox"
}
