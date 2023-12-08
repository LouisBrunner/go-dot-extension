// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsBody2D struct {
  obj gdc.ObjectPtr
}

func createPhysicsBody2D(obj gdc.ObjectPtr) *PhysicsBody2D {
  return &PhysicsBody2D{
    obj: obj,
  }
}

func (me *PhysicsBody2D) BaseClass() string {
  return "PhysicsBody2D"
}
