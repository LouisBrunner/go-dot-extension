// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshInstance2D struct {
  obj gdc.ObjectPtr
}

func (me *MeshInstance2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshInstance2D) BaseClass() string {
  return "MeshInstance2D"
}
