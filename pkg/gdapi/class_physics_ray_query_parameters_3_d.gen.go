// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsRayQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsRayQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsRayQueryParameters3D) BaseClass() string {
  return "PhysicsRayQueryParameters3D"
}
