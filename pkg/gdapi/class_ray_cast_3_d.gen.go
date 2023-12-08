// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RayCast3D struct {
  obj gdc.ObjectPtr
}

func (me *RayCast3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RayCast3D) BaseClass() string {
  return "RayCast3D"
}
