// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StaticBody2D struct {
  obj gdc.ObjectPtr
}

func createStaticBody2D(obj gdc.ObjectPtr) *StaticBody2D {
  return &StaticBody2D{
    obj: obj,
  }
}

func (me *StaticBody2D) BaseClass() string {
  return "StaticBody2D"
}
