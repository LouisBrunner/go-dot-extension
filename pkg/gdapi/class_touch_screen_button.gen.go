// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TouchScreenButton struct {
  obj gdc.ObjectPtr
}

func createTouchScreenButton(obj gdc.ObjectPtr) *TouchScreenButton {
  return &TouchScreenButton{
    obj: obj,
  }
}

func (me *TouchScreenButton) BaseClass() string {
  return "TouchScreenButton"
}
