// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Geometry2D struct {
  obj gdc.ObjectPtr
}

func (me *Geometry2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Geometry2D) BaseClass() string {
  return "Geometry2D"
}
