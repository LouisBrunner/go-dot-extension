// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsPointQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsPointQueryParameters3D(obj gdc.ObjectPtr) *PhysicsPointQueryParameters3D {
  return &PhysicsPointQueryParameters3D{
    obj: obj,
  }
}

func (me *PhysicsPointQueryParameters3D) BaseClass() string {
  return "PhysicsPointQueryParameters3D"
}
