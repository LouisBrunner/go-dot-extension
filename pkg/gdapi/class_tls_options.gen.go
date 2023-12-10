// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  TLSOptionsClient(trusted_chain X509Certificate, common_name_override String, ) TLSOptions {
  classNameV := StringNameFromStr("TLSOptions")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("client")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3565000357) // FIXME: should cache?
  var ret TLSOptions
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(trusted_chain.AsCTypePtr()), gdc.ConstTypePtr(common_name_override.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  TLSOptionsClientUnsafe(trusted_chain X509Certificate, ) TLSOptions {
  classNameV := StringNameFromStr("TLSOptions")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("client_unsafe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2090251749) // FIXME: should cache?
  var ret TLSOptions
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(trusted_chain.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  TLSOptionsServer(key CryptoKey, certificate X509Certificate, ) TLSOptions {
  classNameV := StringNameFromStr("TLSOptions")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36969539) // FIXME: should cache?
  var ret TLSOptions
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(certificate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties