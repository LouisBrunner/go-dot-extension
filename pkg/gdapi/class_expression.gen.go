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

type ptrsForExpressionList struct {
  fnParse gdc.MethodBindPtr
  fnExecute gdc.MethodBindPtr
  fnHasExecuteFailed gdc.MethodBindPtr
  fnGetErrorText gdc.MethodBindPtr
}

var ptrsForExpression ptrsForExpressionList

func initExpressionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Expression")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("parse")
    defer methodName.Destroy()
    ptrsForExpression.fnParse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3069722906))
  }
  {
    methodName := StringNameFromStr("execute")
    defer methodName.Destroy()
    ptrsForExpression.fnExecute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3712471238))
  }
  {
    methodName := StringNameFromStr("has_execute_failed")
    defer methodName.Destroy()
    ptrsForExpression.fnHasExecuteFailed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_error_text")
    defer methodName.Destroy()
    ptrsForExpression.fnGetErrorText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
}

type Expression struct {
  RefCounted
}

func (me *Expression) BaseClass() string {
  return "Expression"
}

func NewExpression() *Expression {
  str := StringNameFromStr("Expression") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Expression{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Expression) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Expression) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Expression) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Expression) Parse(expression String, input_names PackedStringArray, ) Error {
  cargs := []gdc.ConstTypePtr{expression.AsCTypePtr(), input_names.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForExpression.fnParse), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Expression) Execute(inputs Array, base_instance Object, show_error bool, const_calls_only bool, ) Variant {
  cargs := []gdc.ConstTypePtr{inputs.AsCTypePtr(), base_instance.AsCTypePtr(), gdc.ConstTypePtr(&show_error) , gdc.ConstTypePtr(&const_calls_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&show_error)
  pinner.Pin(&const_calls_only)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForExpression.fnExecute), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Expression) HasExecuteFailed() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForExpression.fnHasExecuteFailed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Expression) GetErrorText() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForExpression.fnGetErrorText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
