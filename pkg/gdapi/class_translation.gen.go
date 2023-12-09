// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Translation struct {
  obj gdc.ObjectPtr
}

func (me *Translation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Translation) BaseClass() string {
  return "Translation"
}



// Enums

func (me *Translation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Translation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Translation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Translation) XGetPluralMessage(src_message StringName, src_plural_message StringName, n int, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Translation) XGetMessage(src_message StringName, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Translation) SetLocale(locale String, )  {
  panic("TODO: implement")
}

func  (me *Translation) GetLocale()  {
  panic("TODO: implement")
}

func  (me *Translation) AddMessage(src_message StringName, xlated_message StringName, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Translation) AddPluralMessage(src_message StringName, xlated_messages PackedStringArray, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Translation) GetMessage(src_message StringName, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Translation) GetPluralMessage(src_message StringName, src_plural_message StringName, n int, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Translation) EraseMessage(src_message StringName, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Translation) GetMessageList()  {
  panic("TODO: implement")
}

func  (me *Translation) GetTranslatedMessageList()  {
  panic("TODO: implement")
}

func  (me *Translation) GetMessageCount()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
