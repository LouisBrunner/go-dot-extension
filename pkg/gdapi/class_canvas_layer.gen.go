// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasLayer struct {
  obj gdc.ObjectPtr
}

func (me *CanvasLayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasLayer) BaseClass() string {
  return "CanvasLayer"
}
