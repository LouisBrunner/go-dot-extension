// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeIntParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIntParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIntParameter) BaseClass() string {
  return "VisualShaderNodeIntParameter"
}



// Enums

type VisualShaderNodeIntParameterHint int
const (
  VisualShaderNodeIntParameterHintHintNone VisualShaderNodeIntParameterHint = 0
  VisualShaderNodeIntParameterHintHintRange VisualShaderNodeIntParameterHint = 1
  VisualShaderNodeIntParameterHintHintRangeStep VisualShaderNodeIntParameterHint = 2
  VisualShaderNodeIntParameterHintHintMax VisualShaderNodeIntParameterHint = 3
)

func (me *VisualShaderNodeIntParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIntParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeIntParameter) SetHint(hint VisualShaderNodeIntParameterHint, )  {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2540512075) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntParameter) GetHint() VisualShaderNodeIntParameterHint {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4250814924) // FIXME: should cache?
  var ret VisualShaderNodeIntParameterHint
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeIntParameter) SetMin(value int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntParameter) GetMin() int {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeIntParameter) SetMax(value int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntParameter) GetMax() int {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeIntParameter) SetStep(value int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntParameter) GetStep() int {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeIntParameter) SetDefaultValueEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntParameter) IsDefaultValueEnabled() bool {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeIntParameter) SetDefaultValue(value int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntParameter) GetDefaultValue() int {
  classNameV := StringNameFromStr("VisualShaderNodeIntParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeIntParameter) GetPropHint() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) SetPropHint(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) GetPropMin() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) SetPropMin(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) GetPropMax() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) SetPropMax(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) GetPropStep() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) SetPropStep(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) GetPropDefaultValueEnabled() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) SetPropDefaultValueEnabled(value bool) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) GetPropDefaultValue() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntParameter) SetPropDefaultValue(value int) {
  panic("TODO: implement")
}