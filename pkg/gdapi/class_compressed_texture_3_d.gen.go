// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CompressedTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *CompressedTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedTexture3D) BaseClass() string {
  return "CompressedTexture3D"
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
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CompressedTexture3D) GetLoadPath() String {
  classNameV := StringNameFromStr("CompressedTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_load_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CompressedTexture3D) GetPropLoadPath() String {
  panic("TODO: implement")
}

func (me *CompressedTexture3D) SetPropLoadPath(value String) {
  panic("TODO: implement")
}