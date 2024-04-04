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

type GLTFTextureSampler struct {
  Resource
}

func (me *GLTFTextureSampler) BaseClass() string {
  return "GLTFTextureSampler"
}

func NewGLTFTextureSampler() *GLTFTextureSampler {
  str := StringNameFromStr("GLTFTextureSampler") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFTextureSampler{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GLTFTextureSampler) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFTextureSampler) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFTextureSampler) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFTextureSampler) GetMagFilter() int64 {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mag_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetMagFilter(filter_mode int64, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mag_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFTextureSampler) GetMinFilter() int64 {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetMinFilter(filter_mode int64, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFTextureSampler) GetWrapS() int64 {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wrap_s")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetWrapS(wrap_mode int64, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wrap_s")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFTextureSampler) GetWrapT() int64 {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wrap_t")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetWrapT(wrap_mode int64, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wrap_t")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
