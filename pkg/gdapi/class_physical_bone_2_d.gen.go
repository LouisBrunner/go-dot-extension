// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicalBone2D struct {
  obj gdc.ObjectPtr
}

func createPhysicalBone2D(obj gdc.ObjectPtr) *PhysicalBone2D {
  return &PhysicalBone2D{
    obj: obj,
  }
}

func (me *PhysicalBone2D) BaseClass() string {
  return "PhysicalBone2D"
}
