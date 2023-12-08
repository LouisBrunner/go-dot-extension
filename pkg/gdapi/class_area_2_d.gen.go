// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Area2D struct {
  obj gdc.ObjectPtr
}

func createArea2D(obj gdc.ObjectPtr) *Area2D {
  return &Area2D{
    obj: obj,
  }
}

func (me *Area2D) BaseClass() string {
  return "Area2D"
}
