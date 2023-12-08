// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ProceduralSkyMaterial struct {
  obj gdc.ObjectPtr
}

func createProceduralSkyMaterial(obj gdc.ObjectPtr) *ProceduralSkyMaterial {
  return &ProceduralSkyMaterial{
    obj: obj,
  }
}

func (me *ProceduralSkyMaterial) BaseClass() string {
  return "ProceduralSkyMaterial"
}
