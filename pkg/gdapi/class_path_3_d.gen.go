// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Path3D struct {
  obj gdc.ObjectPtr
}

func createPath3D(obj gdc.ObjectPtr) *Path3D {
  return &Path3D{
    obj: obj,
  }
}

func (me *Path3D) BaseClass() string {
  return "Path3D"
}
