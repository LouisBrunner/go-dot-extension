// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRCamera3D struct {
  obj gdc.ObjectPtr
}

func createXRCamera3D(obj gdc.ObjectPtr) *XRCamera3D {
  return &XRCamera3D{
    obj: obj,
  }
}

func (me *XRCamera3D) BaseClass() string {
  return "XRCamera3D"
}
