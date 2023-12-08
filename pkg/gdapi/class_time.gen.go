// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Time struct {
  obj gdc.ObjectPtr
}

func createTime(obj gdc.ObjectPtr) *Time {
  return &Time{
    obj: obj,
  }
}

func (me *Time) BaseClass() string {
  return "Time"
}
