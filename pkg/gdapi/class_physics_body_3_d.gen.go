// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsBody3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsBody3D(obj gdc.ObjectPtr) *PhysicsBody3D {
  return &PhysicsBody3D{
    obj: obj,
  }
}

func (me *PhysicsBody3D) BaseClass() string {
  return "PhysicsBody3D"
}
