// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Slider struct {
  Range
}

func (me *Slider) BaseClass() string {
  return "Slider"
}



// Enums

func (me *Slider) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Slider) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Slider) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Slider) SetTicks(count int, )  {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ticks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Slider) GetTicks() int {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ticks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Slider) GetTicksOnBorders() bool {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ticks_on_borders")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Slider) SetTicksOnBorders(ticks_on_border bool, )  {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ticks_on_borders")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ticks_on_border), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Slider) SetEditable(editable bool, )  {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&editable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Slider) IsEditable() bool {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Slider) SetScrollable(scrollable bool, )  {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scrollable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scrollable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Slider) IsScrollable() bool {
  classNameV := StringNameFromStr("Slider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_scrollable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SliderDragStartedSignalFn func()

func (me *Slider) ConnectDragStarted(subs SignalSubscribers, fn SliderDragStartedSignalFn) {
  sig := StringNameFromStr("drag_started")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Slider) DisconnectDragStarted(subs SignalSubscribers, fn SliderDragStartedSignalFn) {
  sig := StringNameFromStr("drag_started")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type SliderDragEndedSignalFn func(value_changed bool, )

func (me *Slider) ConnectDragEnded(subs SignalSubscribers, fn SliderDragEndedSignalFn) {
  sig := StringNameFromStr("drag_ended")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Slider) DisconnectDragEnded(subs SignalSubscribers, fn SliderDragEndedSignalFn) {
  sig := StringNameFromStr("drag_ended")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
