// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DManager struct {
  obj gdc.ObjectPtr
}

func createPhysicsServer3DManager(obj gdc.ObjectPtr) *PhysicsServer3DManager {
  return &PhysicsServer3DManager{
    obj: obj,
  }
}

func (me *PhysicsServer3DManager) BaseClass() string {
  return "PhysicsServer3DManager"
}
