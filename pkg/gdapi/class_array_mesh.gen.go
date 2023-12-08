// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ArrayMesh struct {
  obj gdc.ObjectPtr
}

func (me *ArrayMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ArrayMesh) BaseClass() string {
  return "ArrayMesh"
}
