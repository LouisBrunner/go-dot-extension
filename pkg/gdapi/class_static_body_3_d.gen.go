// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StaticBody3D struct {
  obj gdc.ObjectPtr
}

func createStaticBody3D(obj gdc.ObjectPtr) *StaticBody3D {
  return &StaticBody3D{
    obj: obj,
  }
}

func (me *StaticBody3D) BaseClass() string {
  return "StaticBody3D"
}
