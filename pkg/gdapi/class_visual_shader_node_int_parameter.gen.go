// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *VisualShaderNodeIntParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeIntParameter) SetHint(hint VisualShaderNodeIntParameterHint, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) GetHint()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) SetMin(value int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) GetMin()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) SetMax(value int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) GetMax()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) SetStep(value int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) GetStep()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) SetDefaultValueEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) IsDefaultValueEnabled()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) SetDefaultValue(value int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntParameter) GetDefaultValue()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
