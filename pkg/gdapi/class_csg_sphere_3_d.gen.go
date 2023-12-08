// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGSphere3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGSphere3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGSphere3D) BaseClass() string {
  return "CSGSphere3D"
}
