// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisibleOnScreenEnabler2DEnableModeEnableModeInherit VisibleOnScreenEnabler2DEnableMode = 0
  VisibleOnScreenEnabler2DEnableModeEnableModeAlways VisibleOnScreenEnabler2DEnableMode = 1
  VisibleOnScreenEnabler2DEnableModeEnableModeWhenPaused VisibleOnScreenEnabler2DEnableMode = 2
)

func  (me *VisibleOnScreenEnabler2D) SetEnableMode(mode VisibleOnScreenEnabler2DEnableMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisibleOnScreenEnabler2D) GetEnableMode() { // TODO: return value
  // TODO: implement
}

func  (me *VisibleOnScreenEnabler2D) SetEnableNodePath(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisibleOnScreenEnabler2D) GetEnableNodePath() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
