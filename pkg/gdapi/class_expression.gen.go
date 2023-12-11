// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Expression struct {
  obj gdc.ObjectPtr
}

func (me *Expression) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Expression) BaseClass() string {
  return "Expression"
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3658149758) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(expression.AsCTypePtr()), gdc.ConstTypePtr(input_names.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Expression) Execute(inputs Array, base_instance Object, show_error bool, const_calls_only bool, ) Variant {
  classNameV := StringNameFromStr("Expression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("execute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3712471238) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(inputs.AsCTypePtr()), gdc.ConstTypePtr(base_instance.AsCTypePtr()), gdc.ConstTypePtr(&show_error), gdc.ConstTypePtr(&const_calls_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Expression) HasExecuteFailed() bool {
  classNameV := StringNameFromStr("Expression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_execute_failed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Expression) GetErrorText() String {
  classNameV := StringNameFromStr("Expression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_error_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
