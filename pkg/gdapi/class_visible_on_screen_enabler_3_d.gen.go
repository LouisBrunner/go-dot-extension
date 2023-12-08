// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenEnabler3D struct {
  obj gdc.ObjectPtr
}

func createVisibleOnScreenEnabler3D(obj gdc.ObjectPtr) *VisibleOnScreenEnabler3D {
  return &VisibleOnScreenEnabler3D{
    obj: obj,
  }
}

func (me *VisibleOnScreenEnabler3D) BaseClass() string {
  return "VisibleOnScreenEnabler3D"
}
