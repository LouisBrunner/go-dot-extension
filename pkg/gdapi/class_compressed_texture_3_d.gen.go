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

type CompressedTexture3D struct {
  Texture3D
}

func (me *CompressedTexture3D) BaseClass() string {
  return "CompressedTexture3D"
}

func NewCompressedTexture3D() *CompressedTexture3D {
  str := StringNameFromStr("CompressedTexture3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CompressedTexture3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CompressedTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CompressedTexture3D) Load(path String, ) Error {
  classNameV := StringNameFromStr("CompressedTexture3D")
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

func  (me *CompressedTexture3D) GetLoadPath() String {
  classNameV := StringNameFromStr("CompressedTexture3D")
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
