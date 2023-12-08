// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasTexture struct {
  obj gdc.ObjectPtr
}

func (me *CanvasTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasTexture) BaseClass() string {
  return "CanvasTexture"
}
