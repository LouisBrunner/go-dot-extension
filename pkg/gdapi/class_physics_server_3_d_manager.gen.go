// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DManager struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3DManager) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3DManager) BaseClass() string {
  return "PhysicsServer3DManager"
}
