// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState3DExtension struct {
  obj gdc.ObjectPtr
}

func createPhysicsDirectBodyState3DExtension(obj gdc.ObjectPtr) *PhysicsDirectBodyState3DExtension {
  return &PhysicsDirectBodyState3DExtension{
    obj: obj,
  }
}

func (me *PhysicsDirectBodyState3DExtension) BaseClass() string {
  return "PhysicsDirectBodyState3DExtension"
}
