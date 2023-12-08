// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Joint3D struct {
  obj gdc.ObjectPtr
}

func (me *Joint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Joint3D) BaseClass() string {
  return "Joint3D"
}
