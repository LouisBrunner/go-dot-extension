// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeIs struct {
  VisualShaderNode
}

func (me *VisualShaderNodeIs) BaseClass() string {
  return "VisualShaderNodeIs"
}



// Enums

type VisualShaderNodeIsFunction int
const (
  VisualShaderNodeIsFunctionFuncIsInf VisualShaderNodeIsFunction = 0
  VisualShaderNodeIsFunctionFuncIsNan VisualShaderNodeIsFunction = 1
  VisualShaderNodeIsFunctionFuncMax VisualShaderNodeIsFunction = 2
)

func (me *VisualShaderNodeIs) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIs) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIs) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeIs) SetFunction(func_ VisualShaderNodeIsFunction, )  {
  classNameV := StringNameFromStr("VisualShaderNodeIs")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1438374690) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIs) GetFunction() VisualShaderNodeIsFunction {
  classNameV := StringNameFromStr("VisualShaderNodeIs")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 580678557) // FIXME: should cache?
  var ret VisualShaderNodeIsFunction
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
