// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Font struct {
  obj gdc.ObjectPtr
}

func (me *Font) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Font) BaseClass() string {
  return "Font"
}



// Enums

func (me *Font) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Font) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Font) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Font) SetFallbacks(fallbacks Font, )  {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fallbacks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(fallbacks.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Font) GetFallbacks() Font {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fallbacks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) FindVariation(variation_coordinates Dictionary, face_index int, strength float32, transform Transform2D, spacing_top int, spacing_bottom int, spacing_space int, spacing_glyph int, ) RID {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_variation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3344325384) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(variation_coordinates.AsCTypePtr()), gdc.ConstTypePtr(&face_index), gdc.ConstTypePtr(&strength), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(&spacing_top), gdc.ConstTypePtr(&spacing_bottom), gdc.ConstTypePtr(&spacing_space), gdc.ConstTypePtr(&spacing_glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetRids() RID {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetHeight(font_size int, ) float32 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 378113874) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetAscent(font_size int, ) float32 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 378113874) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetDescent(font_size int, ) float32 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 378113874) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetUnderlinePosition(font_size int, ) float32 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 378113874) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetUnderlineThickness(font_size int, ) float32 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 378113874) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetFontName() String {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetFontStyleName() String {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_style_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetOtNameStrings() Dictionary {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ot_name_strings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetFontStyle() TextServerFontStyle {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_style")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2520224254) // FIXME: should cache?
  var ret TextServerFontStyle
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetFontWeight() int {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetFontStretch() int {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetSpacing(spacing TextServerSpacingType, ) int {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1310880908) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spacing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetOpentypeFeatures() Dictionary {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_opentype_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) SetCacheCapacity(single_line int, multi_line int, )  {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cache_capacity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&single_line), gdc.ConstTypePtr(&multi_line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Font) GetStringSize(text String, alignment HorizontalAlignment, width float32, font_size int, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) Vector2 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_string_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1868866121) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetMultilineStringSize(text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) Vector2 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multiline_string_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 519636710) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) DrawString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1983721962) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Font) DrawMultilineString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1171506176) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Font) DrawStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_string_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 623754045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Font) DrawMultilineStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, size int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_string_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3206388178) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Font) GetCharSize(char int, font_size int, ) Vector2 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_char_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3016396712) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) DrawChar(canvas_item RID, pos Vector2, char int, font_size int, modulate Color, ) float32 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_char")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3815617597) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) DrawCharOutline(canvas_item RID, pos Vector2, char int, font_size int, size int, modulate Color, ) float32 {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_char_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 209525354) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) HasChar(char int, ) bool {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_char")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&char), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetSupportedChars() String {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_supported_chars")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) IsLanguageSupported(language String, ) bool {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_language_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) IsScriptSupported(script String, ) bool {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_script_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(script.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetSupportedFeatureList() Dictionary {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_supported_feature_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetSupportedVariationList() Dictionary {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_supported_variation_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Font) GetFaceCount() int {
  classNameV := StringNameFromStr("Font")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_face_count")
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
