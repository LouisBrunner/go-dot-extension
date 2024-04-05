// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForSubViewportContainerList struct {
  fnXPropagateInputEvent gdc.MethodBindPtr
  fnSetStretch gdc.MethodBindPtr
  fnIsStretchEnabled gdc.MethodBindPtr
  fnSetStretchShrink gdc.MethodBindPtr
  fnGetStretchShrink gdc.MethodBindPtr
}

var ptrsForSubViewportContainer ptrsForSubViewportContainerList

func initSubViewportContainerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SubViewportContainer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_stretch")
    defer methodName.Destroy()
    ptrsForSubViewportContainer.fnSetStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_stretch_enabled")
    defer methodName.Destroy()
    ptrsForSubViewportContainer.fnIsStretchEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_stretch_shrink")
    defer methodName.Destroy()
    ptrsForSubViewportContainer.fnSetStretchShrink = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_stretch_shrink")
    defer methodName.Destroy()
    ptrsForSubViewportContainer.fnGetStretchShrink = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type SubViewportContainer struct {
  Container
}

func (me *SubViewportContainer) BaseClass() string {
  return "SubViewportContainer"
}

func NewSubViewportContainer() *SubViewportContainer {
  str := StringNameFromStr("SubViewportContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SubViewportContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SubViewportContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SubViewportContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SubViewportContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SubViewportContainer) SetStretch(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewportContainer.fnSetStretch), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewportContainer) IsStretchEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewportContainer.fnIsStretchEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SubViewportContainer) SetStretchShrink(amount int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewportContainer.fnSetStretchShrink), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewportContainer) GetStretchShrink() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSubViewportContainer.fnGetStretchShrink), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
