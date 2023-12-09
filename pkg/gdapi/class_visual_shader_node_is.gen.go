// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIs struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIs) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIs) BaseClass() string {
  return "VisualShaderNodeIs"
}



// Enums

type VisualShaderNodeIsFunction int
const (
  VisualShaderNodeIsFunctionFuncIsInf VisualShaderNodeIsFunction = 0
  VisualShaderNodeIsFunctionFuncIsNan VisualShaderNodeIsFunction = 1
  VisualShaderNodeIsFunctionFuncMax VisualShaderNodeIsFunction = 2
)

func (me *VisualShaderNodeIs) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIs) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIs) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeIs) SetFunction(func_ VisualShaderNodeIsFunction, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIs) GetFunction()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
