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

type ScrollBar struct {
  Range
}

func (me *ScrollBar) BaseClass() string {
  return "ScrollBar"
}

func NewScrollBar() *ScrollBar {
  str := StringNameFromStr("ScrollBar") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ScrollBar{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ScrollBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScrollBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScrollBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ScrollBar) SetCustomStep(step float64, )  {
  classNameV := StringNameFromStr("ScrollBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&step) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScrollBar) GetCustomStep() float64 {
  classNameV := StringNameFromStr("ScrollBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ScrollBarScrollingSignalFn func()

func (me *ScrollBar) ConnectScrolling(subs SignalSubscribers, fn ScrollBarScrollingSignalFn) {
  sig := StringNameFromStr("scrolling")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScrollBar) DisconnectScrolling(subs SignalSubscribers, fn ScrollBarScrollingSignalFn) {
  sig := StringNameFromStr("scrolling")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
