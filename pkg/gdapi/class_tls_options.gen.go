// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type TLSOptions struct {
  RefCounted
}

func (me *TLSOptions) BaseClass() string {
  return "TLSOptions"
}

func NewTLSOptions() *TLSOptions {
  str := StringNameFromStr("TLSOptions") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TLSOptions{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{trusted_chain.AsCTypePtr(), common_name_override.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTLSOptions()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  TLSOptionsClientUnsafe(trusted_chain X509Certificate, ) TLSOptions {
  classNameV := StringNameFromStr("TLSOptions")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("client_unsafe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2090251749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{trusted_chain.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTLSOptions()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  TLSOptionsServer(key CryptoKey, certificate X509Certificate, ) TLSOptions {
  classNameV := StringNameFromStr("TLSOptions")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36969539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), certificate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTLSOptions()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
