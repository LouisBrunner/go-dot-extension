// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *Expression) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Expression) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Expression) Parse(expression String, input_names PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *Expression) Execute(inputs Array, base_instance Object, show_error bool, const_calls_only bool, )  {
  panic("TODO: implement")
}

func  (me *Expression) HasExecuteFailed()  {
  panic("TODO: implement")
}

func  (me *Expression) GetErrorText()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
