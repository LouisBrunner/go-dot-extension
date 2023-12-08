// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicalBone2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicalBone2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicalBone2D) BaseClass() string {
  return "PhysicalBone2D"
}
