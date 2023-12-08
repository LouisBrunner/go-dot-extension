// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionResult3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsTestMotionResult3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsTestMotionResult3D) BaseClass() string {
  return "PhysicsTestMotionResult3D"
}
