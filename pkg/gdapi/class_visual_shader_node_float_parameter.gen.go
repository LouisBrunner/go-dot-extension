// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeFloatParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFloatParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFloatParameter) BaseClass() string {
  return "VisualShaderNodeFloatParameter"
}



// Enums

type VisualShaderNodeFloatParameterHint int
const (
  VisualShaderNodeFloatParameterHintHintNone VisualShaderNodeFloatParameterHint = 0
  VisualShaderNodeFloatParameterHintHintRange VisualShaderNodeFloatParameterHint = 1
  VisualShaderNodeFloatParameterHintHintRangeStep VisualShaderNodeFloatParameterHint = 2
  VisualShaderNodeFloatParameterHintHintMax VisualShaderNodeFloatParameterHint = 3
)

func (me *VisualShaderNodeFloatParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFloatParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFloatParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeFloatParameter) SetHint(hint VisualShaderNodeFloatParameterHint, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3712586466) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatParameter) GetHint() VisualShaderNodeFloatParameterHint {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3042240429) // FIXME: should cache?
  var ret VisualShaderNodeFloatParameterHint
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeFloatParameter) SetMin(value float32, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatParameter) GetMin() float32 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeFloatParameter) SetMax(value float32, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatParameter) GetMax() float32 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeFloatParameter) SetStep(value float32, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatParameter) GetStep() float32 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValueEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatParameter) IsDefaultValueEnabled() bool {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValue(value float32, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatParameter) GetDefaultValue() float32 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeFloatParameter) GetPropHint() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) SetPropHint(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) GetPropMin() float32 {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) SetPropMin(value float32) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) GetPropMax() float32 {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) SetPropMax(value float32) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) GetPropStep() float32 {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) SetPropStep(value float32) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) GetPropDefaultValueEnabled() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) SetPropDefaultValueEnabled(value bool) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) GetPropDefaultValue() float32 {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatParameter) SetPropDefaultValue(value float32) {
  panic("TODO: implement")
}