// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FogMaterial struct {
  obj gdc.ObjectPtr
}

func createFogMaterial(obj gdc.ObjectPtr) *FogMaterial {
  return &FogMaterial{
    obj: obj,
  }
}

func (me *FogMaterial) BaseClass() string {
  return "FogMaterial"
}
