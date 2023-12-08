// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImmediateMesh struct {
  obj gdc.ObjectPtr
}

func (me *ImmediateMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImmediateMesh) BaseClass() string {
  return "ImmediateMesh"
}
