// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Material struct {
  obj gdc.ObjectPtr
}

func createMaterial(obj gdc.ObjectPtr) *Material {
  return &Material{
    obj: obj,
  }
}

func (me *Material) BaseClass() string {
  return "Material"
}
