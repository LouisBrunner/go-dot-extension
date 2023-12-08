// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shape2D struct {
  obj gdc.ObjectPtr
}

func createShape2D(obj gdc.ObjectPtr) *Shape2D {
  return &Shape2D{
    obj: obj,
  }
}

func (me *Shape2D) BaseClass() string {
  return "Shape2D"
}
