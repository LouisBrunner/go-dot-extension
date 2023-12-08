// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionResult2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsTestMotionResult2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsTestMotionResult2D) BaseClass() string {
  return "PhysicsTestMotionResult2D"
}
