// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("VisualShaderNodeUVFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 765791915) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeUVFunc) GetFunction() VisualShaderNodeUVFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeUVFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3772902164) // FIXME: should cache?
  var ret VisualShaderNodeUVFuncFunction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeUVFunc) GetPropFunction() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeUVFunc) SetPropFunction(value int) {
  panic("TODO: implement")
}