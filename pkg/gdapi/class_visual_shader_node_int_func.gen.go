// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIntFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIntFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIntFunc) BaseClass() string {
  return "VisualShaderNodeIntFunc"
}



// Enums

type VisualShaderNodeIntFuncFunction int
const (
  VisualShaderNodeIntFuncFunctionFuncAbs VisualShaderNodeIntFuncFunction = 0
  VisualShaderNodeIntFuncFunctionFuncNegate VisualShaderNodeIntFuncFunction = 1
  VisualShaderNodeIntFuncFunctionFuncSign VisualShaderNodeIntFuncFunction = 2
  VisualShaderNodeIntFuncFunctionFuncBitwiseNot VisualShaderNodeIntFuncFunction = 3
  VisualShaderNodeIntFuncFunctionFuncMax VisualShaderNodeIntFuncFunction = 4
)

func (me *VisualShaderNodeIntFunc) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIntFunc) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntFunc) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeIntFunc) SetFunction(func_ VisualShaderNodeIntFuncFunction, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntFunc) GetFunction()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
