// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsRayQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func createPhysicsRayQueryParameters2D(obj gdc.ObjectPtr) *PhysicsRayQueryParameters2D {
  return &PhysicsRayQueryParameters2D{
    obj: obj,
  }
}

func (me *PhysicsRayQueryParameters2D) BaseClass() string {
  return "PhysicsRayQueryParameters2D"
}
