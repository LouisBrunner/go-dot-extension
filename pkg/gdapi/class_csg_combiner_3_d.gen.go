// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGCombiner3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGCombiner3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGCombiner3D) BaseClass() string {
  return "CSGCombiner3D"
}
