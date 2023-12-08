// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GrooveJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *GrooveJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GrooveJoint2D) BaseClass() string {
  return "GrooveJoint2D"
}
