// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState3DExtension struct {
  obj gdc.ObjectPtr
}

func createPhysicsDirectSpaceState3DExtension(obj gdc.ObjectPtr) *PhysicsDirectSpaceState3DExtension {
  return &PhysicsDirectSpaceState3DExtension{
    obj: obj,
  }
}

func (me *PhysicsDirectSpaceState3DExtension) BaseClass() string {
  return "PhysicsDirectSpaceState3DExtension"
}
