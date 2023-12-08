// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody3D struct {
  obj gdc.ObjectPtr
}

func createRigidBody3D(obj gdc.ObjectPtr) *RigidBody3D {
  return &RigidBody3D{
    obj: obj,
  }
}

func (me *RigidBody3D) BaseClass() string {
  return "RigidBody3D"
}
