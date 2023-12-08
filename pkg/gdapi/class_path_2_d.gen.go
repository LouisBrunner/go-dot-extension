// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Path2D struct {
  obj gdc.ObjectPtr
}

func createPath2D(obj gdc.ObjectPtr) *Path2D {
  return &Path2D{
    obj: obj,
  }
}

func (me *Path2D) BaseClass() string {
  return "Path2D"
}
