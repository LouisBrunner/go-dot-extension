// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState2D) BaseClass() string {
  return "PhysicsDirectBodyState2D"
}
