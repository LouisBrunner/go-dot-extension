// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RectangleShape2D struct {
  obj gdc.ObjectPtr
}

func createRectangleShape2D(obj gdc.ObjectPtr) *RectangleShape2D {
  return &RectangleShape2D{
    obj: obj,
  }
}

func (me *RectangleShape2D) BaseClass() string {
  return "RectangleShape2D"
}
