// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RayCast2D struct {
  obj gdc.ObjectPtr
}

func (me *RayCast2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RayCast2D) BaseClass() string {
  return "RayCast2D"
}
