// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Control struct {
  obj gdc.ObjectPtr
}

func createControl(obj gdc.ObjectPtr) *Control {
  return &Control{
    obj: obj,
  }
}

func (me *Control) BaseClass() string {
  return "Control"
}
