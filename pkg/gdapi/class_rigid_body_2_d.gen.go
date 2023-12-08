// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody2D struct {
  obj gdc.ObjectPtr
}

func (me *RigidBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RigidBody2D) BaseClass() string {
  return "RigidBody2D"
}
