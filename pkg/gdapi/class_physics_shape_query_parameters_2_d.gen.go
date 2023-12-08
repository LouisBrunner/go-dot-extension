// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsShapeQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsShapeQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsShapeQueryParameters2D) BaseClass() string {
  return "PhysicsShapeQueryParameters2D"
}
