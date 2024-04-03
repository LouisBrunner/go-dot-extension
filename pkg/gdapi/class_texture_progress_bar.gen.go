// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureProgressBar struct {
  Range
}

func (me *TextureProgressBar) BaseClass() string {
  return "TextureProgressBar"
}

func NewTextureProgressBar() *TextureProgressBar {
  str := StringNameFromStr("TextureProgressBar") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextureProgressBar{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TextureProgressBarFillMode int
const (
  TextureProgressBarFillModeFillLeftToRight TextureProgressBarFillMode = 0
  TextureProgressBarFillModeFillRightToLeft TextureProgressBarFillMode = 1
  TextureProgressBarFillModeFillTopToBottom TextureProgressBarFillMode = 2
  TextureProgressBarFillModeFillBottomToTop TextureProgressBarFillMode = 3
  TextureProgressBarFillModeFillClockwise TextureProgressBarFillMode = 4
  TextureProgressBarFillModeFillCounterClockwise TextureProgressBarFillMode = 5
  TextureProgressBarFillModeFillBilinearLeftAndRight TextureProgressBarFillMode = 6
  TextureProgressBarFillModeFillBilinearTopAndBottom TextureProgressBarFillMode = 7
  TextureProgressBarFillModeFillClockwiseAndCounterClockwise TextureProgressBarFillMode = 8
)

func (me *TextureProgressBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureProgressBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureProgressBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextureProgressBar) SetUnderTexture(tex Texture2D, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_under_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tex.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetUnderTexture() Texture2D {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_under_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetProgressTexture(tex Texture2D, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tex.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetProgressTexture() Texture2D {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetOverTexture(tex Texture2D, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_over_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tex.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetOverTexture() Texture2D {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_over_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetFillMode(mode int64, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetFillMode() int64 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetTintUnder(tint Color, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tint_under")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tint.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTintUnder() Color {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tint_under")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetTintProgress(tint Color, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tint_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tint.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTintProgress() Color {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tint_progress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetTintOver(tint Color, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tint_over")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tint.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTintOver() Color {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tint_over")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetTextureProgressOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_progress_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTextureProgressOffset() Vector2 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_progress_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetRadialInitialAngle(mode float64, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radial_initial_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetRadialInitialAngle() float64 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radial_initial_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetRadialCenterOffset(mode Vector2, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radial_center_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mode.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetRadialCenterOffset() Vector2 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radial_center_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1497962370) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetFillDegrees(mode float64, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetFillDegrees() float64 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetStretchMargin(margin Side, value int64, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 437707142) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&value), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetStretchMargin(margin Side, ) int64 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1983885014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetNinePatchStretch(stretch bool, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_nine_patch_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetNinePatchStretch() bool {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_nine_patch_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
