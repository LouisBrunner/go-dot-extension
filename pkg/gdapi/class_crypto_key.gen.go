// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CryptoKey struct {
  obj gdc.ObjectPtr
}

func (me *CryptoKey) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CryptoKey) BaseClass() string {
  return "CryptoKey"
}



// Enums

func (me *CryptoKey) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CryptoKey) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CryptoKey) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CryptoKey) Save(path String, public_only bool, )  {
  panic("TODO: implement")
}

func  (me *CryptoKey) Load(path String, public_only bool, )  {
  panic("TODO: implement")
}

func  (me *CryptoKey) IsPublicOnly()  {
  panic("TODO: implement")
}

func  (me *CryptoKey) SaveToString(public_only bool, )  {
  panic("TODO: implement")
}

func  (me *CryptoKey) LoadFromString(string_key String, public_only bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
