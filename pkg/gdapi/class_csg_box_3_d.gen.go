// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGBox3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGBox3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGBox3D) BaseClass() string {
  return "CSGBox3D"
}
