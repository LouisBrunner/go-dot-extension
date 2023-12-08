// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CurveTexture struct {
  obj gdc.ObjectPtr
}

func createCurveTexture(obj gdc.ObjectPtr) *CurveTexture {
  return &CurveTexture{
    obj: obj,
  }
}

func (me *CurveTexture) BaseClass() string {
  return "CurveTexture"
}
