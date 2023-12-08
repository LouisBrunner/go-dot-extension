// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRNode3D struct {
  obj gdc.ObjectPtr
}

func createXRNode3D(obj gdc.ObjectPtr) *XRNode3D {
  return &XRNode3D{
    obj: obj,
  }
}

func (me *XRNode3D) BaseClass() string {
  return "XRNode3D"
}
