// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionShape3D struct {
  obj gdc.ObjectPtr
}

func createCollisionShape3D(obj gdc.ObjectPtr) *CollisionShape3D {
  return &CollisionShape3D{
    obj: obj,
  }
}

func (me *CollisionShape3D) BaseClass() string {
  return "CollisionShape3D"
}
