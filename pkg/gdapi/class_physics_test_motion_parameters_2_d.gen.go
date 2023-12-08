// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionParameters2D struct {
  obj gdc.ObjectPtr
}

func createPhysicsTestMotionParameters2D(obj gdc.ObjectPtr) *PhysicsTestMotionParameters2D {
  return &PhysicsTestMotionParameters2D{
    obj: obj,
  }
}

func (me *PhysicsTestMotionParameters2D) BaseClass() string {
  return "PhysicsTestMotionParameters2D"
}
