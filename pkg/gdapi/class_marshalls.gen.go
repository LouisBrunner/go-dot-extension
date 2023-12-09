// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Marshalls struct {
  obj gdc.ObjectPtr
}

func (me *Marshalls) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Marshalls) BaseClass() string {
  return "Marshalls"
}



// Enums

func (me *Marshalls) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Marshalls) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Marshalls) VariantToBase64(variant Variant, full_objects bool, )  {
  panic("TODO: implement")
}

func  (me *Marshalls) Base64ToVariant(base64_str String, allow_objects bool, )  {
  panic("TODO: implement")
}

func  (me *Marshalls) RawToBase64(array PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *Marshalls) Base64ToRaw(base64_str String, )  {
  panic("TODO: implement")
}

func  (me *Marshalls) Utf8ToBase64(utf8_str String, )  {
  panic("TODO: implement")
}

func  (me *Marshalls) Base64ToUtf8(base64_str String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
