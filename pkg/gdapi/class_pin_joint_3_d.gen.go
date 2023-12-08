// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PinJoint3D struct {
  obj gdc.ObjectPtr
}

func createPinJoint3D(obj gdc.ObjectPtr) *PinJoint3D {
  return &PinJoint3D{
    obj: obj,
  }
}

func (me *PinJoint3D) BaseClass() string {
  return "PinJoint3D"
}
