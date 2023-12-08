// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CubemapArray struct {
  obj gdc.ObjectPtr
}

func createCubemapArray(obj gdc.ObjectPtr) *CubemapArray {
  return &CubemapArray{
    obj: obj,
  }
}

func (me *CubemapArray) BaseClass() string {
  return "CubemapArray"
}
