// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SubViewport struct {
  obj gdc.ObjectPtr
}

func (me *SubViewport) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SubViewport) BaseClass() string {
  return "SubViewport"
}



// Enums

type SubViewportClearMode int
const (
  SubViewportClearModeClearModeAlways SubViewportClearMode = 0
  SubViewportClearModeClearModeNever SubViewportClearMode = 1
  SubViewportClearModeClearModeOnce SubViewportClearMode = 2
)

type SubViewportUpdateMode int
const (
  SubViewportUpdateModeUpdateDisabled SubViewportUpdateMode = 0
  SubViewportUpdateModeUpdateOnce SubViewportUpdateMode = 1
  SubViewportUpdateModeUpdateWhenVisible SubViewportUpdateMode = 2
  SubViewportUpdateModeUpdateWhenParentVisible SubViewportUpdateMode = 3
  SubViewportUpdateModeUpdateAlways SubViewportUpdateMode = 4
)

func (me *SubViewport) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SubViewport) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SubViewport) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SubViewport) SetSize(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *SubViewport) GetSize()  {
  panic("TODO: implement")
}

func  (me *SubViewport) SetSize2DOverride(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *SubViewport) GetSize2DOverride()  {
  panic("TODO: implement")
}

func  (me *SubViewport) SetSize2DOverrideStretch(enable bool, )  {
  panic("TODO: implement")
}

func  (me *SubViewport) IsSize2DOverrideStretchEnabled()  {
  panic("TODO: implement")
}

func  (me *SubViewport) SetUpdateMode(mode SubViewportUpdateMode, )  {
  panic("TODO: implement")
}

func  (me *SubViewport) GetUpdateMode()  {
  panic("TODO: implement")
}

func  (me *SubViewport) SetClearMode(mode SubViewportClearMode, )  {
  panic("TODO: implement")
}

func  (me *SubViewport) GetClearMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
