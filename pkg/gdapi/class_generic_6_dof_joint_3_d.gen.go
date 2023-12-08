// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Generic6DOFJoint3D struct {
  obj gdc.ObjectPtr
}

func createGeneric6DOFJoint3D(obj gdc.ObjectPtr) *Generic6DOFJoint3D {
  return &Generic6DOFJoint3D{
    obj: obj,
  }
}

func (me *Generic6DOFJoint3D) BaseClass() string {
  return "Generic6DOFJoint3D"
}
