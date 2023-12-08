// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HingeJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *HingeJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HingeJoint3D) BaseClass() string {
  return "HingeJoint3D"
}
