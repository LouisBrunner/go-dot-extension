// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XROrigin3D struct {
  obj gdc.ObjectPtr
}

func createXROrigin3D(obj gdc.ObjectPtr) *XROrigin3D {
  return &XROrigin3D{
    obj: obj,
  }
}

func (me *XROrigin3D) BaseClass() string {
  return "XROrigin3D"
}
