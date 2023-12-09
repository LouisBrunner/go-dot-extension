// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type VisibleOnScreenEnabler3DEnableMode int
const (
  VisibleOnScreenEnabler3DEnableModeEnableModeInherit VisibleOnScreenEnabler3DEnableMode = 0
  VisibleOnScreenEnabler3DEnableModeEnableModeAlways VisibleOnScreenEnabler3DEnableMode = 1
  VisibleOnScreenEnabler3DEnableModeEnableModeWhenPaused VisibleOnScreenEnabler3DEnableMode = 2
)

func (me *VisibleOnScreenEnabler3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenEnabler3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisibleOnScreenEnabler3D) SetEnableMode(mode VisibleOnScreenEnabler3DEnableMode, )  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenEnabler3D) GetEnableMode()  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenEnabler3D) SetEnableNodePath(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenEnabler3D) GetEnableNodePath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
