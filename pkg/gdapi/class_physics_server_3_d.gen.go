// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3D) BaseClass() string {
  return "PhysicsServer3D"
}
