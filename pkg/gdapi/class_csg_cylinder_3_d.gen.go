// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGCylinder3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGCylinder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGCylinder3D) BaseClass() string {
  return "CSGCylinder3D"
}
