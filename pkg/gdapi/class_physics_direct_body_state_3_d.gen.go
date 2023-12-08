// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsDirectBodyState3D(obj gdc.ObjectPtr) *PhysicsDirectBodyState3D {
  return &PhysicsDirectBodyState3D{
    obj: obj,
  }
}

func (me *PhysicsDirectBodyState3D) BaseClass() string {
  return "PhysicsDirectBodyState3D"
}
