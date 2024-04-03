// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeIntFunc struct {
  VisualShaderNode
}

func (me *VisualShaderNodeIntFunc) BaseClass() string {
  return "VisualShaderNodeIntFunc"
}

func NewVisualShaderNodeIntFunc() *VisualShaderNodeIntFunc {
  str := StringNameFromStr("VisualShaderNodeIntFunc") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeIntFunc{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  var ret VisualShaderNodeIntFuncFunction

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
