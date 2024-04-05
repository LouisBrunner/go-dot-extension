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

type ptrsForTextureProgressBarList struct {
  fnSetUnderTexture gdc.MethodBindPtr
  fnGetUnderTexture gdc.MethodBindPtr
  fnSetProgressTexture gdc.MethodBindPtr
  fnGetProgressTexture gdc.MethodBindPtr
  fnSetOverTexture gdc.MethodBindPtr
  fnGetOverTexture gdc.MethodBindPtr
  fnSetFillMode gdc.MethodBindPtr
  fnGetFillMode gdc.MethodBindPtr
  fnSetTintUnder gdc.MethodBindPtr
  fnGetTintUnder gdc.MethodBindPtr
  fnSetTintProgress gdc.MethodBindPtr
  fnGetTintProgress gdc.MethodBindPtr
  fnSetTintOver gdc.MethodBindPtr
  fnGetTintOver gdc.MethodBindPtr
  fnSetTextureProgressOffset gdc.MethodBindPtr
  fnGetTextureProgressOffset gdc.MethodBindPtr
  fnSetRadialInitialAngle gdc.MethodBindPtr
  fnGetRadialInitialAngle gdc.MethodBindPtr
  fnSetRadialCenterOffset gdc.MethodBindPtr
  fnGetRadialCenterOffset gdc.MethodBindPtr
  fnSetFillDegrees gdc.MethodBindPtr
  fnGetFillDegrees gdc.MethodBindPtr
  fnSetStretchMargin gdc.MethodBindPtr
  fnGetStretchMargin gdc.MethodBindPtr
  fnSetNinePatchStretch gdc.MethodBindPtr
  fnGetNinePatchStretch gdc.MethodBindPtr
}

var ptrsForTextureProgressBar ptrsForTextureProgressBarList

func initTextureProgressBarPtrs(iface gdc.Interface) {

  className := StringNameFromStr("TextureProgressBar")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_under_texture")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetUnderTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_under_texture")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetUnderTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_progress_texture")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetProgressTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_progress_texture")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetProgressTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_over_texture")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetOverTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_over_texture")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetOverTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_fill_mode")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetFillMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_fill_mode")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetFillMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_tint_under")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetTintUnder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_tint_under")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetTintUnder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_tint_progress")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetTintProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_tint_progress")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetTintProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_tint_over")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetTintOver = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_tint_over")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetTintOver = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_texture_progress_offset")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetTextureProgressOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_texture_progress_offset")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetTextureProgressOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_radial_initial_angle")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetRadialInitialAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_radial_initial_angle")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetRadialInitialAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_radial_center_offset")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetRadialCenterOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_radial_center_offset")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetRadialCenterOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
  }
  {
    methodName := StringNameFromStr("set_fill_degrees")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetFillDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_fill_degrees")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetFillDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_stretch_margin")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetStretchMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 437707142))
  }
  {
    methodName := StringNameFromStr("get_stretch_margin")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetStretchMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1983885014))
  }
  {
    methodName := StringNameFromStr("set_nine_patch_stretch")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnSetNinePatchStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_nine_patch_stretch")
    defer methodName.Destroy()
    ptrsForTextureProgressBar.fnGetNinePatchStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{tex.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetUnderTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetUnderTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetUnderTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetProgressTexture(tex Texture2D, )  {
  cargs := []gdc.ConstTypePtr{tex.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetProgressTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetProgressTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetProgressTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetOverTexture(tex Texture2D, )  {
  cargs := []gdc.ConstTypePtr{tex.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetOverTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetOverTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetOverTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetFillMode(mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetFillMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetFillMode() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetFillMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetTintUnder(tint Color, )  {
  cargs := []gdc.ConstTypePtr{tint.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetTintUnder), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTintUnder() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetTintUnder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetTintProgress(tint Color, )  {
  cargs := []gdc.ConstTypePtr{tint.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetTintProgress), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTintProgress() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetTintProgress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetTintOver(tint Color, )  {
  cargs := []gdc.ConstTypePtr{tint.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetTintOver), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTintOver() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetTintOver), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetTextureProgressOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetTextureProgressOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetTextureProgressOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetTextureProgressOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetRadialInitialAngle(mode float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetRadialInitialAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetRadialInitialAngle() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetRadialInitialAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetRadialCenterOffset(mode Vector2, )  {
  cargs := []gdc.ConstTypePtr{mode.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetRadialCenterOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetRadialCenterOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetRadialCenterOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextureProgressBar) SetFillDegrees(mode float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetFillDegrees), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetFillDegrees() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetFillDegrees), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetStretchMargin(margin Side, value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetStretchMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetStretchMargin(margin Side, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&margin)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetStretchMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextureProgressBar) SetNinePatchStretch(stretch bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnSetNinePatchStretch), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextureProgressBar) GetNinePatchStretch() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextureProgressBar.fnGetNinePatchStretch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
