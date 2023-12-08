// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody2D struct {
  obj gdc.ObjectPtr
}

func createRigidBody2D(obj gdc.ObjectPtr) *RigidBody2D {
  return &RigidBody2D{
    obj: obj,
  }
}

func (me *RigidBody2D) BaseClass() string {
  return "RigidBody2D"
}
