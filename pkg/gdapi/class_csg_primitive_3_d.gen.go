// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGPrimitive3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGPrimitive3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGPrimitive3D) BaseClass() string {
  return "CSGPrimitive3D"
}
