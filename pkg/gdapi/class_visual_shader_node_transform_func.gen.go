// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("VisualShaderNodeTransformFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2900990409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTransformFunc) GetFunction() VisualShaderNodeTransformFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeTransformFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2839926569) // FIXME: should cache?
  var ret VisualShaderNodeTransformFuncFunction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeTransformFunc) GetPropFunction() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTransformFunc) SetPropFunction(value int) {
  panic("TODO: implement")
}