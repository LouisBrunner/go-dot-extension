// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SliderJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *SliderJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SliderJoint3D) BaseClass() string {
  return "SliderJoint3D"
}
