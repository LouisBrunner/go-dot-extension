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

type VisualShaderNodeFloatParameter struct {
  VisualShaderNodeParameter
}

func (me *VisualShaderNodeFloatParameter) BaseClass() string {
  return "VisualShaderNodeFloatParameter"
}

func NewVisualShaderNodeFloatParameter() *VisualShaderNodeFloatParameter {
  str := StringNameFromStr("VisualShaderNodeFloatParameter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeFloatParameter{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hint) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetHint() VisualShaderNodeFloatParameterHint {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3042240429) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeFloatParameterHint

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShaderNodeFloatParameter) SetMin(value float64, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetMin() float64 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetMax(value float64, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetMax() float64 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetStep(value float64, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetStep() float64 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValueEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) IsDefaultValueEnabled() bool {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValue(value float64, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetDefaultValue() float64 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_value")
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
