// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type World2D struct {
  obj gdc.ObjectPtr
}

func createWorld2D(obj gdc.ObjectPtr) *World2D {
  return &World2D{
    obj: obj,
  }
}

func (me *World2D) BaseClass() string {
  return "World2D"
}
