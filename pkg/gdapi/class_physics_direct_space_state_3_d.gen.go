// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsDirectSpaceState3D(obj gdc.ObjectPtr) *PhysicsDirectSpaceState3D {
  return &PhysicsDirectSpaceState3D{
    obj: obj,
  }
}

func (me *PhysicsDirectSpaceState3D) BaseClass() string {
  return "PhysicsDirectSpaceState3D"
}
