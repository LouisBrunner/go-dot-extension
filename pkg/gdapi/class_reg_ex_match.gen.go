// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RegExMatch struct {
  obj gdc.ObjectPtr
}

func (me *RegExMatch) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RegExMatch) BaseClass() string {
  return "RegExMatch"
}



// Enums

func (me *RegExMatch) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RegExMatch) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RegExMatch) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RegExMatch) GetSubject()  {
  panic("TODO: implement")
}

func  (me *RegExMatch) GetGroupCount()  {
  panic("TODO: implement")
}

func  (me *RegExMatch) GetNames()  {
  panic("TODO: implement")
}

func  (me *RegExMatch) GetStrings()  {
  panic("TODO: implement")
}

func  (me *RegExMatch) GetString(name Variant, )  {
  panic("TODO: implement")
}

func  (me *RegExMatch) GetStart(name Variant, )  {
  panic("TODO: implement")
}

func  (me *RegExMatch) GetEnd(name Variant, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
