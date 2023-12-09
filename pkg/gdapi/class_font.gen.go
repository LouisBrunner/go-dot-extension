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



// Enums

func (me *Font) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Font) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Font) SetFallbacks(fallbacks Font, )  {
  panic("TODO: implement")
}

func  (me *Font) GetFallbacks()  {
  panic("TODO: implement")
}

func  (me *Font) FindVariation(variation_coordinates Dictionary, face_index int, strength float32, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Font) GetRids()  {
  panic("TODO: implement")
}

func  (me *Font) GetHeight(font_size int, )  {
  panic("TODO: implement")
}

func  (me *Font) GetAscent(font_size int, )  {
  panic("TODO: implement")
}

func  (me *Font) GetDescent(font_size int, )  {
  panic("TODO: implement")
}

func  (me *Font) GetUnderlinePosition(font_size int, )  {
  panic("TODO: implement")
}

func  (me *Font) GetUnderlineThickness(font_size int, )  {
  panic("TODO: implement")
}

func  (me *Font) GetFontName()  {
  panic("TODO: implement")
}

func  (me *Font) GetFontStyleName()  {
  panic("TODO: implement")
}

func  (me *Font) GetOtNameStrings()  {
  panic("TODO: implement")
}

func  (me *Font) GetFontStyle()  {
  panic("TODO: implement")
}

func  (me *Font) GetFontWeight()  {
  panic("TODO: implement")
}

func  (me *Font) GetFontStretch()  {
  panic("TODO: implement")
}

func  (me *Font) GetSpacing(spacing TextServerSpacingType, )  {
  panic("TODO: implement")
}

func  (me *Font) GetOpentypeFeatures()  {
  panic("TODO: implement")
}

func  (me *Font) SetCacheCapacity(single_line int, multi_line int, )  {
  panic("TODO: implement")
}

func  (me *Font) GetStringSize(text String, alignment HorizontalAlignment, width float32, font_size int, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *Font) GetMultilineStringSize(text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *Font) DrawString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *Font) DrawMultilineString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *Font) DrawStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *Font) DrawMultilineStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, size int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *Font) GetCharSize(char int, font_size int, )  {
  panic("TODO: implement")
}

func  (me *Font) DrawChar(canvas_item RID, pos Vector2, char int, font_size int, modulate Color, )  {
  panic("TODO: implement")
}

func  (me *Font) DrawCharOutline(canvas_item RID, pos Vector2, char int, font_size int, size int, modulate Color, )  {
  panic("TODO: implement")
}

func  (me *Font) HasChar(char int, )  {
  panic("TODO: implement")
}

func  (me *Font) GetSupportedChars()  {
  panic("TODO: implement")
}

func  (me *Font) IsLanguageSupported(language String, )  {
  panic("TODO: implement")
}

func  (me *Font) IsScriptSupported(script String, )  {
  panic("TODO: implement")
}

func  (me *Font) GetSupportedFeatureList()  {
  panic("TODO: implement")
}

func  (me *Font) GetSupportedVariationList()  {
  panic("TODO: implement")
}

func  (me *Font) GetFaceCount()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
