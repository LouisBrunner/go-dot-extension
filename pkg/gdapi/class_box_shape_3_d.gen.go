// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxShape3D struct {
  obj gdc.ObjectPtr
}

func (me *BoxShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoxShape3D) BaseClass() string {
  return "BoxShape3D"
}
