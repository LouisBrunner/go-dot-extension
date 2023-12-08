// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Joint3D struct {
  obj gdc.ObjectPtr
}

func createJoint3D(obj gdc.ObjectPtr) *Joint3D {
  return &Joint3D{
    obj: obj,
  }
}

func (me *Joint3D) BaseClass() string {
  return "Joint3D"
}
