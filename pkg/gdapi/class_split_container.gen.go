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

type ptrsForSplitContainerList struct {
  fnSetSplitOffset gdc.MethodBindPtr
  fnGetSplitOffset gdc.MethodBindPtr
  fnClampSplitOffset gdc.MethodBindPtr
  fnSetCollapsed gdc.MethodBindPtr
  fnIsCollapsed gdc.MethodBindPtr
  fnSetDraggerVisibility gdc.MethodBindPtr
  fnGetDraggerVisibility gdc.MethodBindPtr
  fnSetVertical gdc.MethodBindPtr
  fnIsVertical gdc.MethodBindPtr
}

var ptrsForSplitContainer ptrsForSplitContainerList

func initSplitContainerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SplitContainer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_split_offset")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnSetSplitOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_split_offset")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnGetSplitOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("clamp_split_offset")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnClampSplitOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_collapsed")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnSetCollapsed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_collapsed")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnIsCollapsed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_dragger_visibility")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnSetDraggerVisibility = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1168273952))
  }
  {
    methodName := StringNameFromStr("get_dragger_visibility")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnGetDraggerVisibility = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 967297479))
  }
  {
    methodName := StringNameFromStr("set_vertical")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnSetVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_vertical")
    defer methodName.Destroy()
    ptrsForSplitContainer.fnIsVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type SplitContainer struct {
  Container
}

func (me *SplitContainer) BaseClass() string {
  return "SplitContainer"
}

func NewSplitContainer() *SplitContainer {
  str := StringNameFromStr("SplitContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SplitContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type SplitContainerDraggerVisibility int
const (
  SplitContainerDraggerVisibilityDraggerVisible SplitContainerDraggerVisibility = 0
  SplitContainerDraggerVisibilityDraggerHidden SplitContainerDraggerVisibility = 1
  SplitContainerDraggerVisibilityDraggerHiddenCollapsed SplitContainerDraggerVisibility = 2
)

func (me *SplitContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SplitContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SplitContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SplitContainer) SetSplitOffset(offset int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnSetSplitOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) GetSplitOffset() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnGetSplitOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SplitContainer) ClampSplitOffset()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnClampSplitOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) SetCollapsed(collapsed bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collapsed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnSetCollapsed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) IsCollapsed() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnIsCollapsed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SplitContainer) SetDraggerVisibility(mode SplitContainerDraggerVisibility, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnSetDraggerVisibility), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) GetDraggerVisibility() SplitContainerDraggerVisibility {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SplitContainerDraggerVisibility

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnGetDraggerVisibility), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SplitContainer) SetVertical(vertical bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnSetVertical), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) IsVertical() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSplitContainer.fnIsVertical), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SplitContainerDraggedSignalFn func(offset int, )

func (me *SplitContainer) ConnectDragged(subs SignalSubscribers, fn SplitContainerDraggedSignalFn) {
  sig := StringNameFromStr("dragged")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SplitContainer) DisconnectDragged(subs SignalSubscribers, fn SplitContainerDraggedSignalFn) {
  sig := StringNameFromStr("dragged")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
