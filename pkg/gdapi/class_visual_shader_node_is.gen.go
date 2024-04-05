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

type ptrsForVisualShaderNodeIsList struct {
  fnSetFunction gdc.MethodBindPtr
  fnGetFunction gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeIs ptrsForVisualShaderNodeIsList

func initVisualShaderNodeIsPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeIs")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_function")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeIs.fnSetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1438374690))
  }
  {
    methodName := StringNameFromStr("get_function")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeIs.fnGetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 580678557))
  }
}

type VisualShaderNodeIs struct {
  VisualShaderNode
}

func (me *VisualShaderNodeIs) BaseClass() string {
  return "VisualShaderNodeIs"
}

func NewVisualShaderNodeIs() *VisualShaderNodeIs {
  str := StringNameFromStr("VisualShaderNodeIs") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeIs{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIs.fnSetFunction), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeIs) GetFunction() VisualShaderNodeIsFunction {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeIsFunction

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIs.fnGetFunction), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
