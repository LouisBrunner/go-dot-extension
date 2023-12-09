// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type X509Certificate struct {
  obj gdc.ObjectPtr
}

func (me *X509Certificate) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *X509Certificate) BaseClass() string {
  return "X509Certificate"
}



// Enums

func (me *X509Certificate) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *X509Certificate) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *X509Certificate) Save(path String, )  {
  panic("TODO: implement")
}

func  (me *X509Certificate) Load(path String, )  {
  panic("TODO: implement")
}

func  (me *X509Certificate) SaveToString()  {
  panic("TODO: implement")
}

func  (me *X509Certificate) LoadFromString(string_ String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
