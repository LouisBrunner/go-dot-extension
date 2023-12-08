// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShapeCast3D struct {
  obj gdc.ObjectPtr
}

func (me *ShapeCast3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShapeCast3D) BaseClass() string {
  return "ShapeCast3D"
}
