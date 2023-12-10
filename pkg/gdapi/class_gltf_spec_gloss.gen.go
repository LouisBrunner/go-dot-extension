// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFSpecGloss struct {
  obj gdc.ObjectPtr
}

func (me *GLTFSpecGloss) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFSpecGloss) BaseClass() string {
  return "GLTFSpecGloss"
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
  var ret Image
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSpecGloss) SetDiffuseImg(diffuse_img Image, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_diffuse_img")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 532598488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(diffuse_img.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSpecGloss) GetDiffuseFactor() Color {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_diffuse_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200896285) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSpecGloss) SetDiffuseFactor(diffuse_factor Color, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_diffuse_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(diffuse_factor.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSpecGloss) GetGlossFactor() float32 {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gloss_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSpecGloss) SetGlossFactor(gloss_factor float32, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gloss_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gloss_factor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSpecGloss) GetSpecularFactor() Color {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200896285) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSpecGloss) SetSpecularFactor(specular_factor Color, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(specular_factor.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSpecGloss) GetSpecGlossImg() Image {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spec_gloss_img")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 564927088) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSpecGloss) SetSpecGlossImg(spec_gloss_img Image, )  {
  classNameV := StringNameFromStr("GLTFSpecGloss")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_spec_gloss_img")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 532598488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(spec_gloss_img.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *GLTFSpecGloss) GetPropDiffuseImg() Object {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) SetPropDiffuseImg(value Object) {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) GetPropDiffuseFactor() Color {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) SetPropDiffuseFactor(value Color) {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) GetPropGlossFactor() float32 {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) SetPropGlossFactor(value float32) {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) GetPropSpecularFactor() Color {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) SetPropSpecularFactor(value Color) {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) GetPropSpecGlossImg() Object {
  panic("TODO: implement")
}

func (me *GLTFSpecGloss) SetPropSpecGlossImg(value Object) {
  panic("TODO: implement")
}