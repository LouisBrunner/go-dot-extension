// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Marker2D struct {
  obj gdc.ObjectPtr
}

func (me *Marker2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Marker2D) BaseClass() string {
  return "Marker2D"
}
