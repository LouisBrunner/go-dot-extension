// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGShape3D) BaseClass() string {
  return "CSGShape3D"
}
