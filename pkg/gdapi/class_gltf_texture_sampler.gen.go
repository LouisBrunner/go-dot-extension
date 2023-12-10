// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFTextureSampler struct {
  obj gdc.ObjectPtr
}

func (me *GLTFTextureSampler) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFTextureSampler) BaseClass() string {
  return "GLTFTextureSampler"
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

func  (me *GLTFTextureSampler) GetMagFilter() int {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mag_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFTextureSampler) SetMagFilter(filter_mode int, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mag_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFTextureSampler) GetMinFilter() int {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFTextureSampler) SetMinFilter(filter_mode int, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFTextureSampler) GetWrapS() int {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wrap_s")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFTextureSampler) SetWrapS(wrap_mode int, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wrap_s")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFTextureSampler) GetWrapT() int {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wrap_t")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFTextureSampler) SetWrapT(wrap_mode int, )  {
  classNameV := StringNameFromStr("GLTFTextureSampler")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wrap_t")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *GLTFTextureSampler) GetPropMagFilter() int {
  panic("TODO: implement")
}

func (me *GLTFTextureSampler) SetPropMagFilter(value int) {
  panic("TODO: implement")
}

func (me *GLTFTextureSampler) GetPropMinFilter() int {
  panic("TODO: implement")
}

func (me *GLTFTextureSampler) SetPropMinFilter(value int) {
  panic("TODO: implement")
}

func (me *GLTFTextureSampler) GetPropWrapS() int {
  panic("TODO: implement")
}

func (me *GLTFTextureSampler) SetPropWrapS(value int) {
  panic("TODO: implement")
}

func (me *GLTFTextureSampler) GetPropWrapT() int {
  panic("TODO: implement")
}

func (me *GLTFTextureSampler) SetPropWrapT(value int) {
  panic("TODO: implement")
}