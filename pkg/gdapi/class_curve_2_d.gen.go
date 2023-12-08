// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve2D struct {
  obj gdc.ObjectPtr
}

func (me *Curve2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve2D) BaseClass() string {
  return "Curve2D"
}
