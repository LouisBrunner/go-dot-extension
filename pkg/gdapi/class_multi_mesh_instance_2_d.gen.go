// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMeshInstance2D struct {
  obj gdc.ObjectPtr
}

func (me *MultiMeshInstance2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMeshInstance2D) BaseClass() string {
  return "MultiMeshInstance2D"
}
