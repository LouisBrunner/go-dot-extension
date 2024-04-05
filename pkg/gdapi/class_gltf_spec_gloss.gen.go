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

type ptrsForGLTFSpecGlossList struct {
  fnGetDiffuseImg gdc.MethodBindPtr
  fnSetDiffuseImg gdc.MethodBindPtr
  fnGetDiffuseFactor gdc.MethodBindPtr
  fnSetDiffuseFactor gdc.MethodBindPtr
  fnGetGlossFactor gdc.MethodBindPtr
  fnSetGlossFactor gdc.MethodBindPtr
  fnGetSpecularFactor gdc.MethodBindPtr
  fnSetSpecularFactor gdc.MethodBindPtr
  fnGetSpecGlossImg gdc.MethodBindPtr
  fnSetSpecGlossImg gdc.MethodBindPtr
}

var ptrsForGLTFSpecGloss ptrsForGLTFSpecGlossList

func initGLTFSpecGlossPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GLTFSpecGloss")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_diffuse_img")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnGetDiffuseImg = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 564927088))
  }
  {
    methodName := StringNameFromStr("set_diffuse_img")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnSetDiffuseImg = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 532598488))
  }
  {
    methodName := StringNameFromStr("get_diffuse_factor")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnGetDiffuseFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200896285))
  }
  {
    methodName := StringNameFromStr("set_diffuse_factor")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnSetDiffuseFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_gloss_factor")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnGetGlossFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_gloss_factor")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnSetGlossFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_specular_factor")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnGetSpecularFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200896285))
  }
  {
    methodName := StringNameFromStr("set_specular_factor")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnSetSpecularFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_spec_gloss_img")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnGetSpecGlossImg = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 564927088))
  }
  {
    methodName := StringNameFromStr("set_spec_gloss_img")
    defer methodName.Destroy()
    ptrsForGLTFSpecGloss.fnSetSpecGlossImg = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 532598488))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnGetDiffuseImg), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetDiffuseImg(diffuse_img Image, )  {
  cargs := []gdc.ConstTypePtr{diffuse_img.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnSetDiffuseImg), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetDiffuseFactor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnGetDiffuseFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetDiffuseFactor(diffuse_factor Color, )  {
  cargs := []gdc.ConstTypePtr{diffuse_factor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnSetDiffuseFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetGlossFactor() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnGetGlossFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFSpecGloss) SetGlossFactor(gloss_factor float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gloss_factor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnSetGlossFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetSpecularFactor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnGetSpecularFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetSpecularFactor(specular_factor Color, )  {
  cargs := []gdc.ConstTypePtr{specular_factor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnSetSpecularFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFSpecGloss) GetSpecGlossImg() Image {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnGetSpecGlossImg), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFSpecGloss) SetSpecGlossImg(spec_gloss_img Image, )  {
  cargs := []gdc.ConstTypePtr{spec_gloss_img.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFSpecGloss.fnSetSpecGlossImg), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
