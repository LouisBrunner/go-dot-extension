// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ClassDB struct {
  obj gdc.ObjectPtr
}

func (me *ClassDB) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ClassDB) BaseClass() string {
  return "ClassDB"
}



// Enums

func (me *ClassDB) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ClassDB) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ClassDB) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ClassDB) GetClassList()  {
  panic("TODO: implement")
}

func  (me *ClassDB) GetInheritersFromClass(class StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) GetParentClass(class StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassExists(class StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) IsParentClass(class StringName, inherits StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) CanInstantiate(class StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) Instantiate(class StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassHasSignal(class StringName, signal StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetSignal(class StringName, signal StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetSignalList(class StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetPropertyList(class StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetProperty(object Object, property StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassSetProperty(object Object, property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassHasMethod(class StringName, method StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetMethodList(class StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetIntegerConstantList(class StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassHasIntegerConstant(class StringName, name StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetIntegerConstant(class StringName, name StringName, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassHasEnum(class StringName, name StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetEnumList(class StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetEnumConstants(class StringName, enum StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) ClassGetIntegerConstantEnum(class StringName, name StringName, no_inheritance bool, )  {
  panic("TODO: implement")
}

func  (me *ClassDB) IsClassEnabled(class StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
