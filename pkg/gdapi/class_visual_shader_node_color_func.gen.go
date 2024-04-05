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

type ptrsForVisualShaderNodeColorFuncList struct {
  fnSetFunction gdc.MethodBindPtr
  fnGetFunction gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeColorFunc ptrsForVisualShaderNodeColorFuncList

func initVisualShaderNodeColorFuncPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeColorFunc")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_function")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeColorFunc.fnSetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3973396138))
  }
  {
    methodName := StringNameFromStr("get_function")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeColorFunc.fnGetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 554863321))
  }
}

type VisualShaderNodeColorFunc struct {
  VisualShaderNode
}

func (me *VisualShaderNodeColorFunc) BaseClass() string {
  return "VisualShaderNodeColorFunc"
}

func NewVisualShaderNodeColorFunc() *VisualShaderNodeColorFunc {
  str := StringNameFromStr("VisualShaderNodeColorFunc") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeColorFunc{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorFunc.fnSetFunction), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeColorFunc) GetFunction() VisualShaderNodeColorFuncFunction {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeColorFuncFunction

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorFunc.fnGetFunction), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
