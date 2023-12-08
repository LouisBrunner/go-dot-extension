// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CurveXYZTexture struct {
  obj gdc.ObjectPtr
}

func createCurveXYZTexture(obj gdc.ObjectPtr) *CurveXYZTexture {
  return &CurveXYZTexture{
    obj: obj,
  }
}

func (me *CurveXYZTexture) BaseClass() string {
  return "CurveXYZTexture"
}
