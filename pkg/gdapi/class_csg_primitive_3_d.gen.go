// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGPrimitive3D struct {
  obj gdc.ObjectPtr
}

func createCSGPrimitive3D(obj gdc.ObjectPtr) *CSGPrimitive3D {
  return &CSGPrimitive3D{
    obj: obj,
  }
}

func (me *CSGPrimitive3D) BaseClass() string {
  return "CSGPrimitive3D"
}
