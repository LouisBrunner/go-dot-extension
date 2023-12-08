// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DampedSpringJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *DampedSpringJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DampedSpringJoint2D) BaseClass() string {
  return "DampedSpringJoint2D"
}
