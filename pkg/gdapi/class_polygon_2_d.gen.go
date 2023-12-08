// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Polygon2D struct {
  obj gdc.ObjectPtr
}

func (me *Polygon2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Polygon2D) BaseClass() string {
  return "Polygon2D"
}
