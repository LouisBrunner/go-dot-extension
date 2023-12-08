// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsShapeQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func createPhysicsShapeQueryParameters2D(obj gdc.ObjectPtr) *PhysicsShapeQueryParameters2D {
  return &PhysicsShapeQueryParameters2D{
    obj: obj,
  }
}

func (me *PhysicsShapeQueryParameters2D) BaseClass() string {
  return "PhysicsShapeQueryParameters2D"
}
