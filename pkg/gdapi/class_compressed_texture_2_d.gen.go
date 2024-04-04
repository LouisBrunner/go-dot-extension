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

type CompressedTexture2D struct {
  Texture2D
}

func (me *CompressedTexture2D) BaseClass() string {
  return "CompressedTexture2D"
}

func NewCompressedTexture2D() *CompressedTexture2D {
  str := StringNameFromStr("CompressedTexture2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CompressedTexture2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CompressedTexture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CompressedTexture2D) Load(path String, ) Error {
  classNameV := StringNameFromStr("CompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CompressedTexture2D) GetLoadPath() String {
  classNameV := StringNameFromStr("CompressedTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_load_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
