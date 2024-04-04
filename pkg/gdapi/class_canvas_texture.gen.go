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

type CanvasTexture struct {
  Texture2D
}

func (me *CanvasTexture) BaseClass() string {
  return "CanvasTexture"
}

func NewCanvasTexture() *CanvasTexture {
  str := StringNameFromStr("CanvasTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CanvasTexture{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasTexture) GetDiffuseTexture() Texture2D {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_diffuse_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasTexture) SetNormalTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasTexture) GetNormalTexture() Texture2D {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasTexture) SetSpecularTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasTexture) GetSpecularTexture() Texture2D {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasTexture) SetSpecularColor(color Color, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasTexture) GetSpecularColor() Color {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasTexture) SetSpecularShininess(shininess float64, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_specular_shininess")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shininess) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasTexture) GetSpecularShininess() float64 {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_specular_shininess")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasTexture) SetTextureFilter(filter CanvasItemTextureFilter, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1037999706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasTexture) GetTextureFilter() CanvasItemTextureFilter {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121960042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CanvasItemTextureFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CanvasTexture) SetTextureRepeat(repeat CanvasItemTextureRepeat, )  {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1716472974) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasTexture) GetTextureRepeat() CanvasItemTextureRepeat {
  classNameV := StringNameFromStr("CanvasTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2667158319) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CanvasItemTextureRepeat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
