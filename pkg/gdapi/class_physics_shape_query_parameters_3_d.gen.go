// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsShapeQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsShapeQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsShapeQueryParameters3D) BaseClass() string {
  return "PhysicsShapeQueryParameters3D"
}
