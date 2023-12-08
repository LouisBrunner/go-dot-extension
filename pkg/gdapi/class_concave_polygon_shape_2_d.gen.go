// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConcavePolygonShape2D struct {
  obj gdc.ObjectPtr
}

func (me *ConcavePolygonShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConcavePolygonShape2D) BaseClass() string {
  return "ConcavePolygonShape2D"
}
