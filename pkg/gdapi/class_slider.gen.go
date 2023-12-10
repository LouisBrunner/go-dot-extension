// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Slider struct {
  obj gdc.ObjectPtr
}

func (me *Slider) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *Slider) GetPropEditable() bool {
  panic("TODO: implement")
}

func (me *Slider) SetPropEditable(value bool) {
  panic("TODO: implement")
}

func (me *Slider) GetPropScrollable() bool {
  panic("TODO: implement")
}

func (me *Slider) SetPropScrollable(value bool) {
  panic("TODO: implement")
}

func (me *Slider) GetPropTickCount() int {
  panic("TODO: implement")
}

func (me *Slider) SetPropTickCount(value int) {
  panic("TODO: implement")
}

func (me *Slider) GetPropTicksOnBorders() bool {
  panic("TODO: implement")
}

func (me *Slider) SetPropTicksOnBorders(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API