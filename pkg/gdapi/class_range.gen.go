// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Range struct {
  obj gdc.ObjectPtr
}

func (me *Range) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Range) BaseClass() string {
  return "Range"
}



// Enums

func (me *Range) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Range) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Range) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Range) GetValue() float32 {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) GetMin() float32 {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) GetMax() float32 {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) GetStep() float32 {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) GetPage() float32 {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_page")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) GetAsRatio() float32 {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_as_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) SetValue(value float32, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) SetValueNoSignal(value float32, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_value_no_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) SetMin(minimum float32, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minimum), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) SetMax(maximum float32, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&maximum), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) SetStep(step float32, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&step), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) SetPage(pagesize float32, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_page")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pagesize), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) SetAsRatio(value float32, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) SetUseRoundedValues(enabled bool, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_rounded_values")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) IsUsingRoundedValues() bool {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_rounded_values")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) SetExpRatio(enabled bool, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exp_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) IsRatioExp() bool {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ratio_exp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) SetAllowGreater(allow bool, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_greater")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) IsGreaterAllowed() bool {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_greater_allowed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) SetAllowLesser(allow bool, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_lesser")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) IsLesserAllowed() bool {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_lesser_allowed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Range) Share(with Node, )  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("share")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(with.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Range) Unshare()  {
  classNameV := StringNameFromStr("Range")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unshare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *Range) GetPropMinValue() float32 {
  panic("TODO: implement")
}

func (me *Range) SetPropMinValue(value float32) {
  panic("TODO: implement")
}

func (me *Range) GetPropMaxValue() float32 {
  panic("TODO: implement")
}

func (me *Range) SetPropMaxValue(value float32) {
  panic("TODO: implement")
}

func (me *Range) GetPropStep() float32 {
  panic("TODO: implement")
}

func (me *Range) SetPropStep(value float32) {
  panic("TODO: implement")
}

func (me *Range) GetPropPage() float32 {
  panic("TODO: implement")
}

func (me *Range) SetPropPage(value float32) {
  panic("TODO: implement")
}

func (me *Range) GetPropValue() float32 {
  panic("TODO: implement")
}

func (me *Range) SetPropValue(value float32) {
  panic("TODO: implement")
}

func (me *Range) GetPropRatio() float32 {
  panic("TODO: implement")
}

func (me *Range) SetPropRatio(value float32) {
  panic("TODO: implement")
}

func (me *Range) GetPropExpEdit() bool {
  panic("TODO: implement")
}

func (me *Range) SetPropExpEdit(value bool) {
  panic("TODO: implement")
}

func (me *Range) GetPropRounded() bool {
  panic("TODO: implement")
}

func (me *Range) SetPropRounded(value bool) {
  panic("TODO: implement")
}

func (me *Range) GetPropAllowGreater() bool {
  panic("TODO: implement")
}

func (me *Range) SetPropAllowGreater(value bool) {
  panic("TODO: implement")
}

func (me *Range) GetPropAllowLesser() bool {
  panic("TODO: implement")
}

func (me *Range) SetPropAllowLesser(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API