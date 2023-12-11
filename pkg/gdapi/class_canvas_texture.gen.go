// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CanvasTexture struct {
  obj gdc.ObjectPtr
}

func (me *CanvasTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasTexture) BaseClass() string {
  return "CanvasTexture"
}



// Enums

func (me *CanvasTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CanvasTexture) SetDiffuseTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_diffuse_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasTexture) GetDiffuseTexture() Texture2D {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_diffuse_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasTexture) SetNormalTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasTexture) GetNormalTexture() Texture2D {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasTexture) SetSpecularTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasTexture) GetSpecularTexture() Texture2D {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasTexture) SetSpecularColor(color Color, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasTexture) GetSpecularColor() Color {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasTexture) SetSpecularShininess(shininess float32, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_shininess")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shininess), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasTexture) GetSpecularShininess() float32 {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_shininess")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasTexture) SetTextureFilter(filter CanvasItemTextureFilter, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1037999706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasTexture) GetTextureFilter() CanvasItemTextureFilter {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121960042) // FIXME: should cache?
  var ret CanvasItemTextureFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasTexture) SetTextureRepeat(repeat CanvasItemTextureRepeat, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1716472974) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasTexture) GetTextureRepeat() CanvasItemTextureRepeat {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2667158319) // FIXME: should cache?
  var ret CanvasItemTextureRepeat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
