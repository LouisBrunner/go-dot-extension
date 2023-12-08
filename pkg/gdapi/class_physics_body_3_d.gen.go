// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsBody3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsBody3D) BaseClass() string {
  return "PhysicsBody3D"
}
