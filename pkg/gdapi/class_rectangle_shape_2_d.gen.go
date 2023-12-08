// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RectangleShape2D struct {
  obj gdc.ObjectPtr
}

func (me *RectangleShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RectangleShape2D) BaseClass() string {
  return "RectangleShape2D"
}
