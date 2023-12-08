// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicalBone3D struct {
  obj gdc.ObjectPtr
}

func createPhysicalBone3D(obj gdc.ObjectPtr) *PhysicalBone3D {
  return &PhysicalBone3D{
    obj: obj,
  }
}

func (me *PhysicalBone3D) BaseClass() string {
  return "PhysicalBone3D"
}
