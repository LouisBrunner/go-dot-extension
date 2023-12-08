// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImporterMesh struct {
  obj gdc.ObjectPtr
}

func (me *ImporterMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImporterMesh) BaseClass() string {
  return "ImporterMesh"
}
