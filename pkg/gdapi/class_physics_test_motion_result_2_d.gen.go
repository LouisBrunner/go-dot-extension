// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionResult2D struct {
  obj gdc.ObjectPtr
}

func createPhysicsTestMotionResult2D(obj gdc.ObjectPtr) *PhysicsTestMotionResult2D {
  return &PhysicsTestMotionResult2D{
    obj: obj,
  }
}

func (me *PhysicsTestMotionResult2D) BaseClass() string {
  return "PhysicsTestMotionResult2D"
}
