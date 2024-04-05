// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForCompressedTextureLayeredList struct {
  fnLoad gdc.MethodBindPtr
  fnGetLoadPath gdc.MethodBindPtr
}

var ptrsForCompressedTextureLayered ptrsForCompressedTextureLayeredList

func initCompressedTextureLayeredPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CompressedTextureLayered")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("load")
    defer methodName.Destroy()
    ptrsForCompressedTextureLayered.fnLoad = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
  }
  {
    methodName := StringNameFromStr("get_load_path")
    defer methodName.Destroy()
    ptrsForCompressedTextureLayered.fnGetLoadPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
}

type CompressedTextureLayered struct {
  TextureLayered
}

func (me *CompressedTextureLayered) BaseClass() string {
  return "CompressedTextureLayered"
}

func NewCompressedTextureLayered() *CompressedTextureLayered {
  str := StringNameFromStr("CompressedTextureLayered") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CompressedTextureLayered{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CompressedTextureLayered) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedTextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedTextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CompressedTextureLayered) Load(path String, ) Error {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompressedTextureLayered.fnLoad), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CompressedTextureLayered) GetLoadPath() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompressedTextureLayered.fnGetLoadPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
