// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Geometry3D struct {
  obj gdc.ObjectPtr
}

func (me *Geometry3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Geometry3D) BaseClass() string {
  return "Geometry3D"
}
