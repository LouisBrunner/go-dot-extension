// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Generic6DOFJoint3D struct {
  obj gdc.ObjectPtr
}

func (me *Generic6DOFJoint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Generic6DOFJoint3D) BaseClass() string {
  return "Generic6DOFJoint3D"
}
