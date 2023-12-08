// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConvexPolygonShape3D struct {
  obj gdc.ObjectPtr
}

func (me *ConvexPolygonShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConvexPolygonShape3D) BaseClass() string {
  return "ConvexPolygonShape3D"
}
