// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeUVFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeUVFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeUVFunc) BaseClass() string {
  return "VisualShaderNodeUVFunc"
}



// Enums

type VisualShaderNodeUVFuncFunction int
const (
  VisualShaderNodeUVFuncFunctionFuncPanning VisualShaderNodeUVFuncFunction = 0
  VisualShaderNodeUVFuncFunctionFuncScaling VisualShaderNodeUVFuncFunction = 1
  VisualShaderNodeUVFuncFunctionFuncMax VisualShaderNodeUVFuncFunction = 2
)

func (me *VisualShaderNodeUVFunc) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeUVFunc) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUVFunc) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeUVFunc) SetFunction(func_ VisualShaderNodeUVFuncFunction, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeUVFunc) GetFunction()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
