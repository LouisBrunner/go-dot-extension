// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformFunc) BaseClass() string {
  return "VisualShaderNodeTransformFunc"
}



// Enums

type VisualShaderNodeTransformFuncFunction int
const (
  VisualShaderNodeTransformFuncFunctionFuncInverse VisualShaderNodeTransformFuncFunction = 0
  VisualShaderNodeTransformFuncFunctionFuncTranspose VisualShaderNodeTransformFuncFunction = 1
  VisualShaderNodeTransformFuncFunctionFuncMax VisualShaderNodeTransformFuncFunction = 2
)

func (me *VisualShaderNodeTransformFunc) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformFunc) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformFunc) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTransformFunc) SetFunction(func_ VisualShaderNodeTransformFuncFunction, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTransformFunc) GetFunction()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
