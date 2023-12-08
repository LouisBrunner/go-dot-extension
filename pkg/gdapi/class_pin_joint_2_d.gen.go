// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PinJoint2D struct {
  obj gdc.ObjectPtr
}

func createPinJoint2D(obj gdc.ObjectPtr) *PinJoint2D {
  return &PinJoint2D{
    obj: obj,
  }
}

func (me *PinJoint2D) BaseClass() string {
  return "PinJoint2D"
}
