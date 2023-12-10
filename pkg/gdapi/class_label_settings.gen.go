// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type LabelSettings struct {
  obj gdc.ObjectPtr
}

func (me *LabelSettings) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LabelSettings) BaseClass() string {
  return "LabelSettings"
}



// Enums

func (me *LabelSettings) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LabelSettings) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LabelSettings) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *LabelSettings) SetLineSpacing(spacing float32, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spacing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetLineSpacing() float32 {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetFont(font Font, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262170328) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetFont() Font {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229501585) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetFontSize(size int, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetFontSize() int {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetFontColor(color Color, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetFontColor() Color {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetOutlineSize(size int, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetOutlineSize() int {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetOutlineColor(color Color, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetOutlineColor() Color {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetShadowSize(size int, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetShadowSize() int {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetShadowColor(color Color, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetShadowColor() Color {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LabelSettings) SetShadowOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LabelSettings) GetShadowOffset() Vector2 {
  classNameV := StringNameFromStr("LabelSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *LabelSettings) GetPropLineSpacing() float32 {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropLineSpacing(value float32) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropFont() Font {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropFont(value Font) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropFontSize() int {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropFontSize(value int) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropFontColor() Color {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropFontColor(value Color) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropOutlineSize() int {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropOutlineSize(value int) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropOutlineColor() Color {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropOutlineColor(value Color) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropShadowSize() int {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropShadowSize(value int) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropShadowColor() Color {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropShadowColor(value Color) {
  panic("TODO: implement")
}

func (me *LabelSettings) GetPropShadowOffset() Vector2 {
  panic("TODO: implement")
}

func (me *LabelSettings) SetPropShadowOffset(value Vector2) {
  panic("TODO: implement")
}