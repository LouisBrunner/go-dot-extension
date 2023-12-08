// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Light2D struct {
  obj gdc.ObjectPtr
}

func (me *Light2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Light2D) BaseClass() string {
  return "Light2D"
}

type Light2DShadowFilter int
const (
  Light2DShadowFilterNone Light2DShadowFilter = 0
  Light2DShadowFilterPcf5 Light2DShadowFilter = 1
  Light2DShadowFilterPcf13 Light2DShadowFilter = 2
)

type Light2DBlendMode int
const (
  Light2DBlendModeAdd Light2DBlendMode = 0
  Light2DBlendModeSub Light2DBlendMode = 1
  Light2DBlendModeMix Light2DBlendMode = 2
)
