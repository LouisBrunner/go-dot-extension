// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DampedSpringJoint2D struct {
  obj gdc.ObjectPtr
}

func createDampedSpringJoint2D(obj gdc.ObjectPtr) *DampedSpringJoint2D {
  return &DampedSpringJoint2D{
    obj: obj,
  }
}

func (me *DampedSpringJoint2D) BaseClass() string {
  return "DampedSpringJoint2D"
}
