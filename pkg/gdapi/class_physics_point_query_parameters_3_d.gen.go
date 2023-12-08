// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsPointQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsPointQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsPointQueryParameters3D) BaseClass() string {
  return "PhysicsPointQueryParameters3D"
}
