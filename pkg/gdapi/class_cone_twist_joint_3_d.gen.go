// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConeTwistJoint3D struct {
  obj gdc.ObjectPtr
}

func createConeTwistJoint3D(obj gdc.ObjectPtr) *ConeTwistJoint3D {
  return &ConeTwistJoint3D{
    obj: obj,
  }
}

func (me *ConeTwistJoint3D) BaseClass() string {
  return "ConeTwistJoint3D"
}
