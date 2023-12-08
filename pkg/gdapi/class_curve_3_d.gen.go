// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve3D struct {
  obj gdc.ObjectPtr
}

func (me *Curve3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve3D) BaseClass() string {
  return "Curve3D"
}
