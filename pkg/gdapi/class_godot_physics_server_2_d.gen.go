// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GodotPhysicsServer2D struct {
  obj gdc.ObjectPtr
}

func createGodotPhysicsServer2D(obj gdc.ObjectPtr) *GodotPhysicsServer2D {
  return &GodotPhysicsServer2D{
    obj: obj,
  }
}

func (me *GodotPhysicsServer2D) BaseClass() string {
  return "GodotPhysicsServer2D"
}
