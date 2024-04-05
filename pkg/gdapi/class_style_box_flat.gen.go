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

type ptrsForStyleBoxFlatList struct {
  fnSetBgColor gdc.MethodBindPtr
  fnGetBgColor gdc.MethodBindPtr
  fnSetBorderColor gdc.MethodBindPtr
  fnGetBorderColor gdc.MethodBindPtr
  fnSetBorderWidthAll gdc.MethodBindPtr
  fnGetBorderWidthMin gdc.MethodBindPtr
  fnSetBorderWidth gdc.MethodBindPtr
  fnGetBorderWidth gdc.MethodBindPtr
  fnSetBorderBlend gdc.MethodBindPtr
  fnGetBorderBlend gdc.MethodBindPtr
  fnSetCornerRadiusAll gdc.MethodBindPtr
  fnSetCornerRadius gdc.MethodBindPtr
  fnGetCornerRadius gdc.MethodBindPtr
  fnSetExpandMargin gdc.MethodBindPtr
  fnSetExpandMarginAll gdc.MethodBindPtr
  fnGetExpandMargin gdc.MethodBindPtr
  fnSetDrawCenter gdc.MethodBindPtr
  fnIsDrawCenterEnabled gdc.MethodBindPtr
  fnSetSkew gdc.MethodBindPtr
  fnGetSkew gdc.MethodBindPtr
  fnSetShadowColor gdc.MethodBindPtr
  fnGetShadowColor gdc.MethodBindPtr
  fnSetShadowSize gdc.MethodBindPtr
  fnGetShadowSize gdc.MethodBindPtr
  fnSetShadowOffset gdc.MethodBindPtr
  fnGetShadowOffset gdc.MethodBindPtr
  fnSetAntiAliased gdc.MethodBindPtr
  fnIsAntiAliased gdc.MethodBindPtr
  fnSetAaSize gdc.MethodBindPtr
  fnGetAaSize gdc.MethodBindPtr
  fnSetCornerDetail gdc.MethodBindPtr
  fnGetCornerDetail gdc.MethodBindPtr
}

var ptrsForStyleBoxFlat ptrsForStyleBoxFlatList

func initStyleBoxFlatPtrs(iface gdc.Interface) {

  className := StringNameFromStr("StyleBoxFlat")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_bg_color")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_bg_color")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_border_color")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetBorderColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_border_color")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetBorderColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_border_width_all")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetBorderWidthAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_border_width_min")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetBorderWidthMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_border_width")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetBorderWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 437707142))
  }
  {
    methodName := StringNameFromStr("get_border_width")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetBorderWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1983885014))
  }
  {
    methodName := StringNameFromStr("set_border_blend")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetBorderBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_border_blend")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetBorderBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_corner_radius_all")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetCornerRadiusAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_corner_radius")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetCornerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2696158768))
  }
  {
    methodName := StringNameFromStr("get_corner_radius")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetCornerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3982397690))
  }
  {
    methodName := StringNameFromStr("set_expand_margin")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetExpandMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4290182280))
  }
  {
    methodName := StringNameFromStr("set_expand_margin_all")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetExpandMarginAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_expand_margin")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetExpandMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
  }
  {
    methodName := StringNameFromStr("set_draw_center")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetDrawCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_draw_center_enabled")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnIsDrawCenterEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_skew")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetSkew = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_skew")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetSkew = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_shadow_color")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetShadowColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_shadow_color")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetShadowColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_shadow_size")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetShadowSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_shadow_size")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetShadowSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_shadow_offset")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetShadowOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_shadow_offset")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetShadowOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_anti_aliased")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetAntiAliased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_anti_aliased")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnIsAntiAliased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_aa_size")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetAaSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_aa_size")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetAaSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_corner_detail")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnSetCornerDetail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_corner_detail")
    defer methodName.Destroy()
    ptrsForStyleBoxFlat.fnGetCornerDetail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type StyleBoxFlat struct {
  StyleBox
}

func (me *StyleBoxFlat) BaseClass() string {
  return "StyleBoxFlat"
}

func NewStyleBoxFlat() *StyleBoxFlat {
  str := StringNameFromStr("StyleBoxFlat") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StyleBoxFlat{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetBgColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetBgColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetBgColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxFlat) SetBorderColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetBorderColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetBorderColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetBorderColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxFlat) SetBorderWidthAll(width int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetBorderWidthAll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetBorderWidthMin() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetBorderWidthMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetBorderWidth(margin Side, width int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetBorderWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetBorderWidth(margin Side, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&margin)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetBorderWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetBorderBlend(blend bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetBorderBlend), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetBorderBlend() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetBorderBlend), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetCornerRadiusAll(radius int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetCornerRadiusAll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) SetCornerRadius(corner Corner, radius int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&corner) , gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetCornerRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetCornerRadius(corner Corner, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&corner) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&corner)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetCornerRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetExpandMargin(margin Side, size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetExpandMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) SetExpandMarginAll(size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetExpandMarginAll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetExpandMargin(margin Side, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&margin)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetExpandMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetDrawCenter(draw_center bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_center) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetDrawCenter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) IsDrawCenterEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnIsDrawCenterEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetSkew(skew Vector2, )  {
  cargs := []gdc.ConstTypePtr{skew.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetSkew), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetSkew() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetSkew), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxFlat) SetShadowColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetShadowColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetShadowColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetShadowColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxFlat) SetShadowSize(size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetShadowSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetShadowSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetShadowSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetShadowOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetShadowOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetShadowOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetShadowOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxFlat) SetAntiAliased(anti_aliased bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&anti_aliased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetAntiAliased), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) IsAntiAliased() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnIsAntiAliased), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetAaSize(size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetAaSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetAaSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetAaSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxFlat) SetCornerDetail(detail int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnSetCornerDetail), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxFlat) GetCornerDetail() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxFlat.fnGetCornerDetail), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
