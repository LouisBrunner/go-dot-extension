// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GodotPhysicsServer3D struct {
  obj gdc.ObjectPtr
}

func createGodotPhysicsServer3D(obj gdc.ObjectPtr) *GodotPhysicsServer3D {
  return &GodotPhysicsServer3D{
    obj: obj,
  }
}

func (me *GodotPhysicsServer3D) BaseClass() string {
  return "GodotPhysicsServer3D"
}
