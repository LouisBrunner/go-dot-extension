// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_split_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) GetSplitOffset() int64 {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_split_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SplitContainer) ClampSplitOffset()  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clamp_split_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) SetCollapsed(collapsed bool, )  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collapsed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) IsCollapsed() bool {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collapsed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SplitContainer) SetDraggerVisibility(mode SplitContainerDraggerVisibility, )  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dragger_visibility")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1168273952) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) GetDraggerVisibility() SplitContainerDraggerVisibility {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dragger_visibility")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 967297479) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SplitContainerDraggerVisibility

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SplitContainer) SetVertical(vertical bool, )  {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SplitContainer) IsVertical() bool {
  classNameV := StringNameFromStr("SplitContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
