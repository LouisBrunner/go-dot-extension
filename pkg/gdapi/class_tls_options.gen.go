// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TLSOptions struct {
  obj gdc.ObjectPtr
}

func (me *TLSOptions) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TLSOptions) BaseClass() string {
  return "TLSOptions"
}



// Enums

func (me *TLSOptions) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TLSOptions) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TLSOptions) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  TLSOptionsClient(trusted_chain X509Certificate, common_name_override String, )  {
  panic("TODO: implement")
}

func  TLSOptionsClientUnsafe(trusted_chain X509Certificate, )  {
  panic("TODO: implement")
}

func  TLSOptionsServer(key CryptoKey, certificate X509Certificate, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
