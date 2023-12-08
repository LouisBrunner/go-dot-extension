// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer2DManager struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer2DManager) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer2DManager) BaseClass() string {
  return "PhysicsServer2DManager"
}
