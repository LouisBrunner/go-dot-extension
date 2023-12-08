// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Marker3D struct {
  obj gdc.ObjectPtr
}

func createMarker3D(obj gdc.ObjectPtr) *Marker3D {
  return &Marker3D{
    obj: obj,
  }
}

func (me *Marker3D) BaseClass() string {
  return "Marker3D"
}
