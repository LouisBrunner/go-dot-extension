// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PathFollow2D struct {
  obj gdc.ObjectPtr
}

func (me *PathFollow2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PathFollow2D) BaseClass() string {
  return "PathFollow2D"
}
