// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Marker2D struct {
  obj gdc.ObjectPtr
}

func createMarker2D(obj gdc.ObjectPtr) *Marker2D {
  return &Marker2D{
    obj: obj,
  }
}

func (me *Marker2D) BaseClass() string {
  return "Marker2D"
}
