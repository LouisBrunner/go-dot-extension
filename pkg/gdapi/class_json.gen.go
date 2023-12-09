// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JSON struct {
  obj gdc.ObjectPtr
}

func (me *JSON) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JSON) BaseClass() string {
  return "JSON"
}



// Enums

func (me *JSON) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JSON) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JSON) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  JSONStringify(data Variant, indent String, sort_keys bool, full_precision bool, )  {
  panic("TODO: implement")
}

func  JSONParseString(json_string String, )  {
  panic("TODO: implement")
}

func  (me *JSON) Parse(json_text String, keep_text bool, )  {
  panic("TODO: implement")
}

func  (me *JSON) GetData()  {
  panic("TODO: implement")
}

func  (me *JSON) SetData(data Variant, )  {
  panic("TODO: implement")
}

func  (me *JSON) GetParsedText()  {
  panic("TODO: implement")
}

func  (me *JSON) GetErrorLine()  {
  panic("TODO: implement")
}

func  (me *JSON) GetErrorMessage()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
