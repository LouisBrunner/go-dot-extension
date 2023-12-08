// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionShape2D struct {
  obj gdc.ObjectPtr
}

func createCollisionShape2D(obj gdc.ObjectPtr) *CollisionShape2D {
  return &CollisionShape2D{
    obj: obj,
  }
}

func (me *CollisionShape2D) BaseClass() string {
  return "CollisionShape2D"
}
