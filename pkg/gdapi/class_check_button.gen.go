// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CheckButton struct {
  obj gdc.ObjectPtr
}

func createCheckButton(obj gdc.ObjectPtr) *CheckButton {
  return &CheckButton{
    obj: obj,
  }
}

func (me *CheckButton) BaseClass() string {
  return "CheckButton"
}
