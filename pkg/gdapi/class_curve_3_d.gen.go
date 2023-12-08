// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve3D struct {
  obj gdc.ObjectPtr
}

func createCurve3D(obj gdc.ObjectPtr) *Curve3D {
  return &Curve3D{
    obj: obj,
  }
}

func (me *Curve3D) BaseClass() string {
  return "Curve3D"
}
