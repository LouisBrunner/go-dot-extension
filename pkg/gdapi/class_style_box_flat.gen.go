// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StyleBoxFlat struct {
  StyleBox
}

func (me *StyleBoxFlat) BaseClass() string {
  return "StyleBoxFlat"
}



// Enums

func (me *StyleBoxFlat) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StyleBoxFlat) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StyleBoxFlat) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StyleBoxFlat) SetBgColor(color Color, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetBgColor() Color {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetBorderColor(color Color, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetBorderColor() Color {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetBorderWidthAll(width int, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_width_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetBorderWidthMin() int {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_width_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetBorderWidth(margin Side, width int, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 437707142) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetBorderWidth(margin Side, ) int {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1983885014) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetBorderBlend(blend bool, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetBorderBlend() bool {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetCornerRadiusAll(radius int, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_corner_radius_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) SetCornerRadius(corner Corner, radius int, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_corner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2696158768) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&corner), gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetCornerRadius(corner Corner, ) int {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_corner_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3982397690) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&corner), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetExpandMargin(margin Side, size float32, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4290182280) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) SetExpandMarginAll(size float32, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_margin_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetExpandMargin(margin Side, ) float32 {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_expand_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869120046) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetDrawCenter(draw_center bool, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_center), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) IsDrawCenterEnabled() bool {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_center_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetSkew(skew Vector2, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skew")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skew.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetSkew() Vector2 {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skew")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetShadowColor(color Color, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetShadowColor() Color {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetShadowSize(size int, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetShadowSize() int {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetShadowOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetShadowOffset() Vector2 {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetAntiAliased(anti_aliased bool, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_anti_aliased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anti_aliased), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) IsAntiAliased() bool {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_anti_aliased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetAaSize(size float32, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aa_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetAaSize() float32 {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aa_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StyleBoxFlat) SetCornerDetail(detail int, )  {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_corner_detail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StyleBoxFlat) GetCornerDetail() int {
  classNameV := StringNameFromStr("StyleBoxFlat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_corner_detail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
