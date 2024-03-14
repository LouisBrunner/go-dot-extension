// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeColorFunc struct {
  VisualShaderNode
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
  classNameV := StringNameFromStr("VisualShaderNodeColorFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3973396138) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeColorFunc) GetFunction() VisualShaderNodeColorFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeColorFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 554863321) // FIXME: should cache?
  var ret VisualShaderNodeColorFuncFunction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
