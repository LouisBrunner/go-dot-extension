// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SphereMesh struct {
  obj gdc.ObjectPtr
}

func (me *SphereMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SphereMesh) BaseClass() string {
  return "SphereMesh"
}
