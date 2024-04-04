// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeUIntFunc struct {
  VisualShaderNode
}

func (me *VisualShaderNodeUIntFunc) BaseClass() string {
  return "VisualShaderNodeUIntFunc"
}

func NewVisualShaderNodeUIntFunc() *VisualShaderNodeUIntFunc {
  str := StringNameFromStr("VisualShaderNodeUIntFunc") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeUIntFunc{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("VisualShaderNodeUIntFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2273148961) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeUIntFunc) GetFunction() VisualShaderNodeUIntFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeUIntFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4187123296) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeUIntFuncFunction

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
