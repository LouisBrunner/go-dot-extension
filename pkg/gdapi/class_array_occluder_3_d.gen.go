// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ArrayOccluder3D struct {
  obj gdc.ObjectPtr
}

func createArrayOccluder3D(obj gdc.ObjectPtr) *ArrayOccluder3D {
  return &ArrayOccluder3D{
    obj: obj,
  }
}

func (me *ArrayOccluder3D) BaseClass() string {
  return "ArrayOccluder3D"
}
