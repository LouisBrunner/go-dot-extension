// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeColorFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeColorFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeColorFunc) BaseClass() string {
  return "VisualShaderNodeColorFunc"
}



// Enums

type VisualShaderNodeColorFuncFunction int
const (
  VisualShaderNodeColorFuncFunctionFuncGrayscale VisualShaderNodeColorFuncFunction = 0
  VisualShaderNodeColorFuncFunctionFuncHsv2Rgb VisualShaderNodeColorFuncFunction = 1
  VisualShaderNodeColorFuncFunctionFuncRgb2Hsv VisualShaderNodeColorFuncFunction = 2
  VisualShaderNodeColorFuncFunctionFuncSepia VisualShaderNodeColorFuncFunction = 3
  VisualShaderNodeColorFuncFunctionFuncMax VisualShaderNodeColorFuncFunction = 4
)

func (me *VisualShaderNodeColorFunc) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeColorFunc) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeColorFunc) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeColorFunc) SetFunction(func_ VisualShaderNodeColorFuncFunction, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeColorFunc) GetFunction()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
