// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsMaterial struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsMaterial) BaseClass() string {
  return "PhysicsMaterial"
}
