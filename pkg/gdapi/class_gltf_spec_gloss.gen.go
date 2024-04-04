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

type GLTFSpecGloss struct {
  Resource
}

func (me *GLTFSpecGloss) BaseClass() string {
  return "GLTFSpecGloss"
}

func NewGLTFSpecGloss() *GLTFSpecGloss {
  str := StringNameFromStr("GLTFSpecGloss") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFSpecGloss{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GLTFSpecGloss) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFSpecGloss) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFSpecGloss) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFSpecGloss) GetDiffuseImg() Image {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_diffuse_img")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 564927088) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetDiffuseImg(diffuse_img Image, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_diffuse_img")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 532598488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{diffuse_img.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetDiffuseFactor() Color {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_diffuse_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200896285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetDiffuseFactor(diffuse_factor Color, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_diffuse_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{diffuse_factor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetGlossFactor() float64 {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gloss_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFSpecGloss) SetGlossFactor(gloss_factor float64, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gloss_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gloss_factor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetSpecularFactor() Color {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200896285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetSpecularFactor(specular_factor Color, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{specular_factor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetSpecGlossImg() Image {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spec_gloss_img")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 564927088) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetSpecGlossImg(spec_gloss_img Image, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spec_gloss_img")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 532598488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{spec_gloss_img.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
