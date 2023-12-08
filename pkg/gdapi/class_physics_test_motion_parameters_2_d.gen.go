// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsTestMotionParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsTestMotionParameters2D) BaseClass() string {
  return "PhysicsTestMotionParameters2D"
}
