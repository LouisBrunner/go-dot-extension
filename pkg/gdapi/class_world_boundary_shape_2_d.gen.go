// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorldBoundaryShape2D struct {
  obj gdc.ObjectPtr
}

func (me *WorldBoundaryShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorldBoundaryShape2D) BaseClass() string {
  return "WorldBoundaryShape2D"
}
