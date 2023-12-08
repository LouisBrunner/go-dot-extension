// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRController3D struct {
  obj gdc.ObjectPtr
}

func createXRController3D(obj gdc.ObjectPtr) *XRController3D {
  return &XRController3D{
    obj: obj,
  }
}

func (me *XRController3D) BaseClass() string {
  return "XRController3D"
}
