// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Path2D struct {
  obj gdc.ObjectPtr
}

func (me *Path2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Path2D) BaseClass() string {
  return "Path2D"
}
