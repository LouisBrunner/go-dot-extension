// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SoftBody3D struct {
  obj gdc.ObjectPtr
}

func (me *SoftBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SoftBody3D) BaseClass() string {
  return "SoftBody3D"
}
