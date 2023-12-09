// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) GetHint()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) SetMin(value float32, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) GetMin()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) SetMax(value float32, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) GetMax()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) SetStep(value float32, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) GetStep()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValueEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) IsDefaultValueEnabled()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) SetDefaultValue(value float32, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeFloatParameter) GetDefaultValue()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
