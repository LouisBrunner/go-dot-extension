// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Mesh struct {
  obj gdc.ObjectPtr
}

func (me *Mesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Mesh) BaseClass() string {
  return "Mesh"
}
