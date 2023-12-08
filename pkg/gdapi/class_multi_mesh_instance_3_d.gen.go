// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *MultiMeshInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMeshInstance3D) BaseClass() string {
  return "MultiMeshInstance3D"
}
