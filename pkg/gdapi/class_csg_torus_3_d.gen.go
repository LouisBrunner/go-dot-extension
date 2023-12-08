// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGTorus3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGTorus3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGTorus3D) BaseClass() string {
  return "CSGTorus3D"
}
