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

func  TLSOptionsClient(trusted_chain X509Certificate, common_name_override String, ) { // TODO: return value
  // TODO: implement
}

func  TLSOptionsClientUnsafe(trusted_chain X509Certificate, ) { // TODO: return value
  // TODO: implement
}

func  TLSOptionsServer(key CryptoKey, certificate X509Certificate, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
