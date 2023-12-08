// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState2DExtension struct {
  obj gdc.ObjectPtr
}

func createPhysicsDirectSpaceState2DExtension(obj gdc.ObjectPtr) *PhysicsDirectSpaceState2DExtension {
  return &PhysicsDirectSpaceState2DExtension{
    obj: obj,
  }
}

func (me *PhysicsDirectSpaceState2DExtension) BaseClass() string {
  return "PhysicsDirectSpaceState2DExtension"
}
