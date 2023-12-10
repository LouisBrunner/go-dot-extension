// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("VisualShaderNodeIntFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 424195284) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntFunc) GetFunction() VisualShaderNodeIntFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeIntFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2753496911) // FIXME: should cache?
  var ret VisualShaderNodeIntFuncFunction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeIntFunc) GetPropFunction() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeIntFunc) SetPropFunction(value int) {
  panic("TODO: implement")
}