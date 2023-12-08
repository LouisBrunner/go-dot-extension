// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState2D) BaseClass() string {
  return "PhysicsDirectSpaceState2D"
}
