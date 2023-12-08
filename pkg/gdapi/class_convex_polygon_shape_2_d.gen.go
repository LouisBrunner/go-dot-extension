// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConvexPolygonShape2D struct {
  obj gdc.ObjectPtr
}

func (me *ConvexPolygonShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConvexPolygonShape2D) BaseClass() string {
  return "ConvexPolygonShape2D"
}
