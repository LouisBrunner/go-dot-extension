// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState3D) BaseClass() string {
  return "PhysicsDirectBodyState3D"
}
