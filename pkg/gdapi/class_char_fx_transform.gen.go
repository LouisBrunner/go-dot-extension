// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CharFXTransform struct {
  obj gdc.ObjectPtr
}

func (me *CharFXTransform) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CharFXTransform) BaseClass() string {
  return "CharFXTransform"
}



// Enums

func (me *CharFXTransform) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CharFXTransform) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CharFXTransform) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CharFXTransform) GetRange() Vector2i {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2741790807) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetRange(range_ Vector2i, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(range_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetElapsedTime() float32 {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_elapsed_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetElapsedTime(time float32, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_elapsed_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) IsVisible() bool {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetVisibility(visibility bool, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visibility), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) IsOutline() bool {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetOutline(outline bool, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outline), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetOffset() Vector2 {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1497962370) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetColor() Color {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200896285) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetColor(color Color, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetEnvironment() Dictionary {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetEnvironment(environment Dictionary, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_environment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(environment.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetGlyphIndex() int {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetGlyphIndex(glyph_index int, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&glyph_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetRelativeIndex() int {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_relative_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetRelativeIndex(relative_index int, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_relative_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&relative_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetGlyphCount() int {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetGlyphCount(glyph_count int, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&glyph_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetGlyphFlags() int {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_glyph_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetGlyphFlags(glyph_flags int, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_glyph_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&glyph_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CharFXTransform) GetFont() RID {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CharFXTransform) SetFont(font RID, )  {
  classNameV := StringNameFromStr("CharFXTransform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *CharFXTransform) GetPropRange() Vector2i {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropRange(value Vector2i) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropElapsedTime() float32 {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropElapsedTime(value float32) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropVisible() bool {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropVisible(value bool) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropOutline() bool {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropOutline(value bool) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropOffset() Vector2 {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropEnv() Dictionary {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropEnv(value Dictionary) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropGlyphIndex() int {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropGlyphIndex(value int) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropGlyphCount() int {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropGlyphCount(value int) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropGlyphFlags() int {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropGlyphFlags(value int) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropRelativeIndex() int {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropRelativeIndex(value int) {
  panic("TODO: implement")
}

func (me *CharFXTransform) GetPropFont() RID {
  panic("TODO: implement")
}

func (me *CharFXTransform) SetPropFont(value RID) {
  panic("TODO: implement")
}