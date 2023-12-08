// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRAnchor3D struct {
  obj gdc.ObjectPtr
}

func createXRAnchor3D(obj gdc.ObjectPtr) *XRAnchor3D {
  return &XRAnchor3D{
    obj: obj,
  }
}

func (me *XRAnchor3D) BaseClass() string {
  return "XRAnchor3D"
}
