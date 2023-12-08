// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenEnabler3D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenEnabler3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenEnabler3D) BaseClass() string {
  return "VisibleOnScreenEnabler3D"
}

type VisibleOnScreenEnabler3DEnableMode int
const (
  VisibleOnScreenEnabler3DEnableModeInherit VisibleOnScreenEnabler3DEnableMode = 0
  VisibleOnScreenEnabler3DEnableModeAlways VisibleOnScreenEnabler3DEnableMode = 1
  VisibleOnScreenEnabler3DEnableModeWhenPaused VisibleOnScreenEnabler3DEnableMode = 2
)
