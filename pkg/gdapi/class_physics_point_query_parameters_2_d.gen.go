// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsPointQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func createPhysicsPointQueryParameters2D(obj gdc.ObjectPtr) *PhysicsPointQueryParameters2D {
  return &PhysicsPointQueryParameters2D{
    obj: obj,
  }
}

func (me *PhysicsPointQueryParameters2D) BaseClass() string {
  return "PhysicsPointQueryParameters2D"
}
