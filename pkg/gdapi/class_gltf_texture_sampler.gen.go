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

type ptrsForGLTFTextureSamplerList struct {
  fnGetMagFilter gdc.MethodBindPtr
  fnSetMagFilter gdc.MethodBindPtr
  fnGetMinFilter gdc.MethodBindPtr
  fnSetMinFilter gdc.MethodBindPtr
  fnGetWrapS gdc.MethodBindPtr
  fnSetWrapS gdc.MethodBindPtr
  fnGetWrapT gdc.MethodBindPtr
  fnSetWrapT gdc.MethodBindPtr
}

var ptrsForGLTFTextureSampler ptrsForGLTFTextureSamplerList

func initGLTFTextureSamplerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GLTFTextureSampler")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_mag_filter")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnGetMagFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_mag_filter")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnSetMagFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_min_filter")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnGetMinFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_min_filter")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnSetMinFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_wrap_s")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnGetWrapS = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_wrap_s")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnSetWrapS = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_wrap_t")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnGetWrapT = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_wrap_t")
    defer methodName.Destroy()
    ptrsForGLTFTextureSampler.fnSetWrapT = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnGetMagFilter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetMagFilter(filter_mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnSetMagFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFTextureSampler) GetMinFilter() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnGetMinFilter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetMinFilter(filter_mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnSetMinFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFTextureSampler) GetWrapS() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnGetWrapS), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetWrapS(wrap_mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnSetWrapS), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFTextureSampler) GetWrapT() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnGetWrapT), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFTextureSampler) SetWrapT(wrap_mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFTextureSampler.fnSetWrapT), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
