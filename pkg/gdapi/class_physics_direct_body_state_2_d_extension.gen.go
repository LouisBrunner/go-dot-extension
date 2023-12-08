// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState2DExtension struct {
  obj gdc.ObjectPtr
}

func createPhysicsDirectBodyState2DExtension(obj gdc.ObjectPtr) *PhysicsDirectBodyState2DExtension {
  return &PhysicsDirectBodyState2DExtension{
    obj: obj,
  }
}

func (me *PhysicsDirectBodyState2DExtension) BaseClass() string {
  return "PhysicsDirectBodyState2DExtension"
}
