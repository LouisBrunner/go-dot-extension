// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionResult3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsTestMotionResult3D(obj gdc.ObjectPtr) *PhysicsTestMotionResult3D {
  return &PhysicsTestMotionResult3D{
    obj: obj,
  }
}

func (me *PhysicsTestMotionResult3D) BaseClass() string {
  return "PhysicsTestMotionResult3D"
}
