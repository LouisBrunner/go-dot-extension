// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PanoramaSkyMaterial struct {
  obj gdc.ObjectPtr
}

func createPanoramaSkyMaterial(obj gdc.ObjectPtr) *PanoramaSkyMaterial {
  return &PanoramaSkyMaterial{
    obj: obj,
  }
}

func (me *PanoramaSkyMaterial) BaseClass() string {
  return "PanoramaSkyMaterial"
}
