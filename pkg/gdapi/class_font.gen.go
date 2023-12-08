// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Font struct {
  obj gdc.ObjectPtr
}

func (me *Font) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Font) BaseClass() string {
  return "Font"
}

func  (me *Font) SetFallbacks(fallbacks Font, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetFallbacks() { // TODO: return value
  // TODO: implement
}

func  (me *Font) FindVariation(variation_coordinates Dictionary, face_index int, strength float32, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetRids() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetHeight(font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetAscent(font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetDescent(font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetUnderlinePosition(font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetUnderlineThickness(font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetFontName() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetFontStyleName() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetOtNameStrings() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetFontStyle() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetFontWeight() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetFontStretch() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetSpacing(spacing TextServerSpacingType, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetOpentypeFeatures() { // TODO: return value
  // TODO: implement
}

func  (me *Font) SetCacheCapacity(single_line int, multi_line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetStringSize(text String, alignment HorizontalAlignment, width float32, font_size int, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetMultilineStringSize(text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) DrawString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) DrawMultilineString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) DrawStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) DrawMultilineStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, size int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetCharSize(char int, font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) DrawChar(canvas_item RID, pos Vector2, char int, font_size int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) DrawCharOutline(canvas_item RID, pos Vector2, char int, font_size int, size int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) HasChar(char int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetSupportedChars() { // TODO: return value
  // TODO: implement
}

func  (me *Font) IsLanguageSupported(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) IsScriptSupported(script String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetSupportedFeatureList() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetSupportedVariationList() { // TODO: return value
  // TODO: implement
}

func  (me *Font) GetFaceCount() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
