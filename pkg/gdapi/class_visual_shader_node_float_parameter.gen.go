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

type ptrsForVisualShaderNodeFloatParameterList struct {
  fnSetHint gdc.MethodBindPtr
  fnGetHint gdc.MethodBindPtr
  fnSetMin gdc.MethodBindPtr
  fnGetMin gdc.MethodBindPtr
  fnSetMax gdc.MethodBindPtr
  fnGetMax gdc.MethodBindPtr
  fnSetStep gdc.MethodBindPtr
  fnGetStep gdc.MethodBindPtr
  fnSetDefaultValueEnabled gdc.MethodBindPtr
  fnIsDefaultValueEnabled gdc.MethodBindPtr
  fnSetDefaultValue gdc.MethodBindPtr
  fnGetDefaultValue gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeFloatParameter ptrsForVisualShaderNodeFloatParameterList

func initVisualShaderNodeFloatParameterPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeFloatParameter")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_hint")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnSetHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3712586466))
  }
  {
    methodName := StringNameFromStr("get_hint")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnGetHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3042240429))
  }
  {
    methodName := StringNameFromStr("set_min")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnSetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_min")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnGetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnSetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnGetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_step")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnSetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_step")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnGetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_default_value_enabled")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnSetDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_default_value_enabled")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnIsDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_default_value")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnSetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_default_value")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatParameter.fnGetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hint) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnSetHint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetHint() VisualShaderNodeFloatParameterHint {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeFloatParameterHint

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnGetHint), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShaderNodeFloatParameter) SetMin(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnSetMin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetMin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnGetMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetMax(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnSetMax), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetMax() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnGetMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetStep(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnSetStep), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetStep() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnGetStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValueEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnSetDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) IsDefaultValueEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnIsDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValue(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnSetDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatParameter) GetDefaultValue() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatParameter.fnGetDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
