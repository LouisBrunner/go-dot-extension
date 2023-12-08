// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BaseButton struct {
  obj gdc.ObjectPtr
}

func createBaseButton(obj gdc.ObjectPtr) *BaseButton {
  return &BaseButton{
    obj: obj,
  }
}

func (me *BaseButton) BaseClass() string {
  return "BaseButton"
}
