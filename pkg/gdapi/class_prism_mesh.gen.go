// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PrismMesh struct {
  obj gdc.ObjectPtr
}

func (me *PrismMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PrismMesh) BaseClass() string {
  return "PrismMesh"
}
