// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionObject2D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionObject2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionObject2D) BaseClass() string {
  return "CollisionObject2D"
}
