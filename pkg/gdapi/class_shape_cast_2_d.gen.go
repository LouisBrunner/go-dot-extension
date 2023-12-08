// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShapeCast2D struct {
  obj gdc.ObjectPtr
}

func (me *ShapeCast2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShapeCast2D) BaseClass() string {
  return "ShapeCast2D"
}
