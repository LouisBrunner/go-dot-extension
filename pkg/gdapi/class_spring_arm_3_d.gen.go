// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpringArm3D struct {
  obj gdc.ObjectPtr
}

func (me *SpringArm3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpringArm3D) BaseClass() string {
  return "SpringArm3D"
}
