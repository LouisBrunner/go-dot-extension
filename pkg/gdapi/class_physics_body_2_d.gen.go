// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsBody2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsBody2D) BaseClass() string {
  return "PhysicsBody2D"
}
