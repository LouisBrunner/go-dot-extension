// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Timer struct {
  obj gdc.ObjectPtr
}

func createTimer(obj gdc.ObjectPtr) *Timer {
  return &Timer{
    obj: obj,
  }
}

func (me *Timer) BaseClass() string {
  return "Timer"
}
