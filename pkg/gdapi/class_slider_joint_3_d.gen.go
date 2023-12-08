// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SliderJoint3D struct {
  obj gdc.ObjectPtr
}

func createSliderJoint3D(obj gdc.ObjectPtr) *SliderJoint3D {
  return &SliderJoint3D{
    obj: obj,
  }
}

func (me *SliderJoint3D) BaseClass() string {
  return "SliderJoint3D"
}
