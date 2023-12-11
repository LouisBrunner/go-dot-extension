// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFTexture struct {
  obj gdc.ObjectPtr
}

func (me *GLTFTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFTexture) BaseClass() string {
  return "GLTFTexture"
}



// Enums

func (me *GLTFTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFTexture) GetSrcImage() int {
  classNameV := StringNameFromStr("GLTFTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_src_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFTexture) SetSrcImage(src_image int, )  {
  classNameV := StringNameFromStr("GLTFTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_src_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&src_image), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFTexture) GetSampler() int {
  classNameV := StringNameFromStr("GLTFTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sampler")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFTexture) SetSampler(sampler int, )  {
  classNameV := StringNameFromStr("GLTFTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sampler")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sampler), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
