// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenEnabler2D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenEnabler2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenEnabler2D) BaseClass() string {
  return "VisibleOnScreenEnabler2D"
}

type VisibleOnScreenEnabler2DEnableMode int
const (
  VisibleOnScreenEnabler2DEnableModeInherit VisibleOnScreenEnabler2DEnableMode = 0
  VisibleOnScreenEnabler2DEnableModeAlways VisibleOnScreenEnabler2DEnableMode = 1
  VisibleOnScreenEnabler2DEnableModeWhenPaused VisibleOnScreenEnabler2DEnableMode = 2
)
