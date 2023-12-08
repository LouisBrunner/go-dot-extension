// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GeometryInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *GeometryInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GeometryInstance3D) BaseClass() string {
  return "GeometryInstance3D"
}
