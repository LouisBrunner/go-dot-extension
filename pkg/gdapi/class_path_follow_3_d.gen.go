// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PathFollow3D struct {
  obj gdc.ObjectPtr
}

func createPathFollow3D(obj gdc.ObjectPtr) *PathFollow3D {
  return &PathFollow3D{
    obj: obj,
  }
}

func (me *PathFollow3D) BaseClass() string {
  return "PathFollow3D"
}
