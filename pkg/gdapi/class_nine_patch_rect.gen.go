// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NinePatchRect struct {
  obj gdc.ObjectPtr
}

func (me *NinePatchRect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NinePatchRect) BaseClass() string {
  return "NinePatchRect"
}

type NinePatchRectAxisStretchMode int
const (
  NinePatchRectAxisStretchModeStretch NinePatchRectAxisStretchMode = 0
  NinePatchRectAxisStretchModeTile NinePatchRectAxisStretchMode = 1
  NinePatchRectAxisStretchModeTileFit NinePatchRectAxisStretchMode = 2
)
