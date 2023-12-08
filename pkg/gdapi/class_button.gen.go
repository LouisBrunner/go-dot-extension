// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Button struct {
  obj gdc.ObjectPtr
}

func createButton(obj gdc.ObjectPtr) *Button {
  return &Button{
    obj: obj,
  }
}

func (me *Button) BaseClass() string {
  return "Button"
}
