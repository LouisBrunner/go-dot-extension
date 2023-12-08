// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PrimitiveMesh struct {
  obj gdc.ObjectPtr
}

func (me *PrimitiveMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PrimitiveMesh) BaseClass() string {
  return "PrimitiveMesh"
}
