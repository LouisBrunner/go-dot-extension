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

type VisualShaderNodeTransformFunc struct {
  VisualShaderNode
}

func (me *VisualShaderNodeTransformFunc) BaseClass() string {
  return "VisualShaderNodeTransformFunc"
}

func NewVisualShaderNodeTransformFunc() *VisualShaderNodeTransformFunc {
  str := StringNameFromStr("VisualShaderNodeTransformFunc") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeTransformFunc{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeTransformFunc) GetFunction() VisualShaderNodeTransformFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeTransformFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2839926569) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeTransformFuncFunction

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
