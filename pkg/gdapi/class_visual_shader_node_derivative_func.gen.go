// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeDerivativeFunc struct {
  VisualShaderNode
}

func (me *VisualShaderNodeDerivativeFunc) BaseClass() string {
  return "VisualShaderNodeDerivativeFunc"
}

func NewVisualShaderNodeDerivativeFunc() *VisualShaderNodeDerivativeFunc {
  str := StringNameFromStr("VisualShaderNodeDerivativeFunc") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeDerivativeFunc{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeDerivativeFuncOpType int
const (
  VisualShaderNodeDerivativeFuncOpTypeOpTypeScalar VisualShaderNodeDerivativeFuncOpType = 0
  VisualShaderNodeDerivativeFuncOpTypeOpTypeVector2D VisualShaderNodeDerivativeFuncOpType = 1
  VisualShaderNodeDerivativeFuncOpTypeOpTypeVector3D VisualShaderNodeDerivativeFuncOpType = 2
  VisualShaderNodeDerivativeFuncOpTypeOpTypeVector4D VisualShaderNodeDerivativeFuncOpType = 3
  VisualShaderNodeDerivativeFuncOpTypeOpTypeMax VisualShaderNodeDerivativeFuncOpType = 4
)

type VisualShaderNodeDerivativeFuncFunction int
const (
  VisualShaderNodeDerivativeFuncFunctionFuncSum VisualShaderNodeDerivativeFuncFunction = 0
  VisualShaderNodeDerivativeFuncFunctionFuncX VisualShaderNodeDerivativeFuncFunction = 1
  VisualShaderNodeDerivativeFuncFunctionFuncY VisualShaderNodeDerivativeFuncFunction = 2
  VisualShaderNodeDerivativeFuncFunctionFuncMax VisualShaderNodeDerivativeFuncFunction = 3
)

type VisualShaderNodeDerivativeFuncPrecision int
const (
  VisualShaderNodeDerivativeFuncPrecisionPrecisionNone VisualShaderNodeDerivativeFuncPrecision = 0
  VisualShaderNodeDerivativeFuncPrecisionPrecisionCoarse VisualShaderNodeDerivativeFuncPrecision = 1
  VisualShaderNodeDerivativeFuncPrecisionPrecisionFine VisualShaderNodeDerivativeFuncPrecision = 2
  VisualShaderNodeDerivativeFuncPrecisionPrecisionMax VisualShaderNodeDerivativeFuncPrecision = 3
)

func (me *VisualShaderNodeDerivativeFunc) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeDerivativeFunc) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeDerivativeFunc) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeDerivativeFunc) SetOpType(type_ VisualShaderNodeDerivativeFuncOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeDerivativeFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 377800221) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeDerivativeFunc) GetOpType() VisualShaderNodeDerivativeFuncOpType {
  classNameV := StringNameFromStr("VisualShaderNodeDerivativeFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3997800514) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeDerivativeFuncOpType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShaderNodeDerivativeFunc) SetFunction(func_ VisualShaderNodeDerivativeFuncFunction, )  {
  classNameV := StringNameFromStr("VisualShaderNodeDerivativeFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1944704156) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeDerivativeFunc) GetFunction() VisualShaderNodeDerivativeFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeDerivativeFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2389093396) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeDerivativeFuncFunction

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShaderNodeDerivativeFunc) SetPrecision(precision VisualShaderNodeDerivativeFuncPrecision, )  {
  classNameV := StringNameFromStr("VisualShaderNodeDerivativeFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_precision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797270566) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&precision) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeDerivativeFunc) GetPrecision() VisualShaderNodeDerivativeFuncPrecision {
  classNameV := StringNameFromStr("VisualShaderNodeDerivativeFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_precision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3822547323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeDerivativeFuncPrecision

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
