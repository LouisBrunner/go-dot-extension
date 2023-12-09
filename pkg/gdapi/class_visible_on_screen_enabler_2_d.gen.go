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



// Enums

type VisibleOnScreenEnabler2DEnableMode int
const (
  VisibleOnScreenEnabler2DEnableModeEnableModeInherit VisibleOnScreenEnabler2DEnableMode = 0
  VisibleOnScreenEnabler2DEnableModeEnableModeAlways VisibleOnScreenEnabler2DEnableMode = 1
  VisibleOnScreenEnabler2DEnableModeEnableModeWhenPaused VisibleOnScreenEnabler2DEnableMode = 2
)

func (me *VisibleOnScreenEnabler2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenEnabler2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisibleOnScreenEnabler2D) SetEnableMode(mode VisibleOnScreenEnabler2DEnableMode, )  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenEnabler2D) GetEnableMode()  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenEnabler2D) SetEnableNodePath(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenEnabler2D) GetEnableNodePath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
