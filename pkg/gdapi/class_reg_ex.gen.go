// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RegEx struct {
  obj gdc.ObjectPtr
}

func (me *RegEx) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RegEx) BaseClass() string {
  return "RegEx"
}



// Enums

func (me *RegEx) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RegEx) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RegEx) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  RegExCreateFromString(pattern String, )  {
  panic("TODO: implement")
}

func  (me *RegEx) Clear()  {
  panic("TODO: implement")
}

func  (me *RegEx) Compile(pattern String, )  {
  panic("TODO: implement")
}

func  (me *RegEx) Search(subject String, offset int, end int, )  {
  panic("TODO: implement")
}

func  (me *RegEx) SearchAll(subject String, offset int, end int, )  {
  panic("TODO: implement")
}

func  (me *RegEx) Sub(subject String, replacement String, all bool, offset int, end int, )  {
  panic("TODO: implement")
}

func  (me *RegEx) IsValid()  {
  panic("TODO: implement")
}

func  (me *RegEx) GetPattern()  {
  panic("TODO: implement")
}

func  (me *RegEx) GetGroupCount()  {
  panic("TODO: implement")
}

func  (me *RegEx) GetNames()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
