// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Joint2D struct {
  obj gdc.ObjectPtr
}

func createJoint2D(obj gdc.ObjectPtr) *Joint2D {
  return &Joint2D{
    obj: obj,
  }
}

func (me *Joint2D) BaseClass() string {
  return "Joint2D"
}
