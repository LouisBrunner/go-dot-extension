// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StyleBoxTexture struct {
  StyleBox
}

func (me *StyleBoxTexture) BaseClass() string {
  return "StyleBoxTexture"
}

func NewStyleBoxTexture() *StyleBoxTexture {
  str := StringNameFromStr("StyleBoxTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StyleBoxTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type StyleBoxTextureAxisStretchMode int
const (
  StyleBoxTextureAxisStretchModeAxisStretchModeStretch StyleBoxTextureAxisStretchMode = 0
  StyleBoxTextureAxisStretchModeAxisStretchModeTile StyleBoxTextureAxisStretchMode = 1
  StyleBoxTextureAxisStretchModeAxisStretchModeTileFit StyleBoxTextureAxisStretchMode = 2
)

func (me *StyleBoxTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StyleBoxTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StyleBoxTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StyleBoxTexture) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) GetTexture() Texture2D {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxTexture) SetTextureMargin(margin Side, size float64, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4290182280) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) SetTextureMarginAll(size float64, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_margin_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) GetTextureMargin(margin Side, ) float64 {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869120046) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxTexture) SetExpandMargin(margin Side, size float64, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4290182280) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) SetExpandMarginAll(size float64, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_margin_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) GetExpandMargin(margin Side, ) float64 {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_expand_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869120046) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxTexture) SetRegionRect(region Rect2, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) GetRegionRect() Rect2 {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxTexture) SetDrawCenter(enable bool, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) IsDrawCenterEnabled() bool {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_center_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxTexture) SetModulate(color Color, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) GetModulate() Color {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxTexture) SetHAxisStretchMode(mode StyleBoxTextureAxisStretchMode, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_h_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2965538783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) GetHAxisStretchMode() StyleBoxTextureAxisStretchMode {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_h_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3807744063) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret StyleBoxTextureAxisStretchMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StyleBoxTexture) SetVAxisStretchMode(mode StyleBoxTextureAxisStretchMode, )  {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_v_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2965538783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxTexture) GetVAxisStretchMode() StyleBoxTextureAxisStretchMode {
  classNameV := StringNameFromStr("StyleBoxTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_axis_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3807744063) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret StyleBoxTextureAxisStretchMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
