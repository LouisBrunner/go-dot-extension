// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionParameters3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsTestMotionParameters3D(obj gdc.ObjectPtr) *PhysicsTestMotionParameters3D {
  return &PhysicsTestMotionParameters3D{
    obj: obj,
  }
}

func (me *PhysicsTestMotionParameters3D) BaseClass() string {
  return "PhysicsTestMotionParameters3D"
}
