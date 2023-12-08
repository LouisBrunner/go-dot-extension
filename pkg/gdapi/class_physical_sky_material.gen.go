// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicalSkyMaterial struct {
  obj gdc.ObjectPtr
}

func createPhysicalSkyMaterial(obj gdc.ObjectPtr) *PhysicalSkyMaterial {
  return &PhysicalSkyMaterial{
    obj: obj,
  }
}

func (me *PhysicalSkyMaterial) BaseClass() string {
  return "PhysicalSkyMaterial"
}
