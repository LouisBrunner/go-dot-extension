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
  classNameV := StringNameFromStr("Expression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3069722906) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{expression.AsCTypePtr(), input_names.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Expression) Execute(inputs Array, base_instance Object, show_error bool, const_calls_only bool, ) Variant {
  classNameV := StringNameFromStr("Expression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("execute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3712471238) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{inputs.AsCTypePtr(), base_instance.AsCTypePtr(), gdc.ConstTypePtr(&show_error) , gdc.ConstTypePtr(&const_calls_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&show_error)
  pinner.Pin(&const_calls_only)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Expression) HasExecuteFailed() bool {
  classNameV := StringNameFromStr("Expression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_execute_failed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Expression) GetErrorText() String {
  classNameV := StringNameFromStr("Expression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_error_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
