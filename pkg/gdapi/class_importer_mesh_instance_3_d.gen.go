// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImporterMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *ImporterMeshInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImporterMeshInstance3D) BaseClass() string {
  return "ImporterMeshInstance3D"
}
