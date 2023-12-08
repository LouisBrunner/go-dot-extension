// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve2D struct {
  obj gdc.ObjectPtr
}

func createCurve2D(obj gdc.ObjectPtr) *Curve2D {
  return &Curve2D{
    obj: obj,
  }
}

func (me *Curve2D) BaseClass() string {
  return "Curve2D"
}
