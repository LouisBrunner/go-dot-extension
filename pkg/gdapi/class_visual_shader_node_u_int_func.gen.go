// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeUIntFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeUIntFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeUIntFunc) BaseClass() string {
  return "VisualShaderNodeUIntFunc"
}



// Enums

type VisualShaderNodeUIntFuncFunction int
const (
  VisualShaderNodeUIntFuncFunctionFuncNegate VisualShaderNodeUIntFuncFunction = 0
  VisualShaderNodeUIntFuncFunctionFuncBitwiseNot VisualShaderNodeUIntFuncFunction = 1
  VisualShaderNodeUIntFuncFunctionFuncMax VisualShaderNodeUIntFuncFunction = 2
)

func (me *VisualShaderNodeUIntFunc) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeUIntFunc) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUIntFunc) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeUIntFunc) SetFunction(func_ VisualShaderNodeUIntFuncFunction, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeUIntFunc) GetFunction()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
