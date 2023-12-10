// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureProgressBar struct {
  obj gdc.ObjectPtr
}

func (me *TextureProgressBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureProgressBar) BaseClass() string {
  return "TextureProgressBar"
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
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureProgressBar) SetFillMode(mode int, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureProgressBar) GetFillMode() int {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureProgressBar) SetRadialInitialAngle(mode float32, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radial_initial_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureProgressBar) GetRadialInitialAngle() float32 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radial_initial_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureProgressBar) SetFillDegrees(mode float32, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureProgressBar) GetFillDegrees() float32 {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureProgressBar) SetStretchMargin(margin Side, value int, )  {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 437707142) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextureProgressBar) GetStretchMargin(margin Side, ) int {
  classNameV := StringNameFromStr("TextureProgressBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1983885014) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TextureProgressBar) GetPropFillMode() int {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropFillMode(value int) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropNinePatchStretch() bool {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropNinePatchStretch(value bool) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropStretchMarginLeft() int {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropStretchMarginLeft(value int) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropStretchMarginTop() int {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropStretchMarginTop(value int) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropStretchMarginRight() int {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropStretchMarginRight(value int) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropStretchMarginBottom() int {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropStretchMarginBottom(value int) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropTextureUnder() Texture2D {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropTextureUnder(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropTextureOver() Texture2D {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropTextureOver(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropTextureProgress() Texture2D {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropTextureProgress(value Texture2D) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropTextureProgressOffset() Vector2 {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropTextureProgressOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropTintUnder() Color {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropTintUnder(value Color) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropTintOver() Color {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropTintOver(value Color) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropTintProgress() Color {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropTintProgress(value Color) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropRadialInitialAngle() float32 {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropRadialInitialAngle(value float32) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropRadialFillDegrees() float32 {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropRadialFillDegrees(value float32) {
  panic("TODO: implement")
}

func (me *TextureProgressBar) GetPropRadialCenterOffset() Vector2 {
  panic("TODO: implement")
}

func (me *TextureProgressBar) SetPropRadialCenterOffset(value Vector2) {
  panic("TODO: implement")
}