// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody3D struct {
  obj gdc.ObjectPtr
}

func (me *RigidBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RigidBody3D) BaseClass() string {
  return "RigidBody3D"
}
