// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionShape2D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionShape2D) BaseClass() string {
  return "CollisionShape2D"
}
