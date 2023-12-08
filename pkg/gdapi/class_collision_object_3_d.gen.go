// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionObject3D struct {
  obj gdc.ObjectPtr
}

func createCollisionObject3D(obj gdc.ObjectPtr) *CollisionObject3D {
  return &CollisionObject3D{
    obj: obj,
  }
}

func (me *CollisionObject3D) BaseClass() string {
  return "CollisionObject3D"
}
