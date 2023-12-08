// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextServerExtension struct {
  obj gdc.ObjectPtr
}

func (me *TextServerExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextServerExtension) BaseClass() string {
  return "TextServerExtension"
}

func  (me *TextServerExtension) XHasFeature(feature TextServerFeature, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XGetName() { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XGetFeatures() { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFreeRid(rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XHas(rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XLoadSupportData(filename String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XGetSupportDataFilename() { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XGetSupportDataInfo() { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XSaveSupportData(filename String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XIsLocaleRightToLeft(locale String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XNameToTag(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XTagToName(tag int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XCreateFont() { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetData(font_rid RID, data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetDataPtr(font_rid RID, data_ptr *uint8, data_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetFaceIndex(font_rid RID, face_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetFaceIndex(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetFaceCount(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetStyle(font_rid RID, style TextServerFontStyle, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetStyle(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetName(font_rid RID, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetName(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetOtNameStrings(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetStyleName(font_rid RID, name_style String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetStyleName(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetWeight(font_rid RID, weight int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetWeight(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetStretch(font_rid RID, stretch int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetStretch(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetAntialiasing(font_rid RID, antialiasing TextServerFontAntialiasing, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetAntialiasing(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetGenerateMipmaps(font_rid RID, generate_mipmaps bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGenerateMipmaps(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetMultichannelSignedDistanceField(font_rid RID, msdf bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontIsMultichannelSignedDistanceField(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetMsdfPixelRange(font_rid RID, msdf_pixel_range int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetMsdfPixelRange(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetMsdfSize(font_rid RID, msdf_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetMsdfSize(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetFixedSize(font_rid RID, fixed_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetFixedSize(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetAllowSystemFallback(font_rid RID, allow_system_fallback bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontIsAllowSystemFallback(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetForceAutohinter(font_rid RID, force_autohinter bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontIsForceAutohinter(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetHinting(font_rid RID, hinting TextServerHinting, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetHinting(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetSubpixelPositioning(font_rid RID, subpixel_positioning TextServerSubpixelPositioning, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetSubpixelPositioning(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetEmbolden(font_rid RID, strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetEmbolden(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetTransform(font_rid RID, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetTransform(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetVariationCoordinates(font_rid RID, variation_coordinates Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetVariationCoordinates(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetOversampling(font_rid RID, oversampling float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetOversampling(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetSizeCacheList(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontClearSizeCache(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRemoveSizeCache(font_rid RID, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetAscent(font_rid RID, size int, ascent float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetAscent(font_rid RID, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetDescent(font_rid RID, size int, descent float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetDescent(font_rid RID, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetUnderlinePosition(font_rid RID, size int, underline_position float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetUnderlinePosition(font_rid RID, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetUnderlineThickness(font_rid RID, size int, underline_thickness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetUnderlineThickness(font_rid RID, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetScale(font_rid RID, size int, scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetScale(font_rid RID, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetTextureCount(font_rid RID, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontClearTextures(font_rid RID, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRemoveTexture(font_rid RID, size Vector2i, texture_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetTextureImage(font_rid RID, size Vector2i, texture_index int, image Image, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetTextureImage(font_rid RID, size Vector2i, texture_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetTextureOffsets(font_rid RID, size Vector2i, texture_index int, offset PackedInt32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetTextureOffsets(font_rid RID, size Vector2i, texture_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphList(font_rid RID, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontClearGlyphs(font_rid RID, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRemoveGlyph(font_rid RID, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphAdvance(font_rid RID, size int, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetGlyphAdvance(font_rid RID, size int, glyph int, advance Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphOffset(font_rid RID, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetGlyphOffset(font_rid RID, size Vector2i, glyph int, offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphSize(font_rid RID, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetGlyphSize(font_rid RID, size Vector2i, glyph int, gl_size Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphUvRect(font_rid RID, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetGlyphUvRect(font_rid RID, size Vector2i, glyph int, uv_rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int, texture_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphTextureRid(font_rid RID, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphTextureSize(font_rid RID, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphContours(font_rid RID, size int, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetKerningList(font_rid RID, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontClearKerningMap(font_rid RID, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRemoveKerning(font_rid RID, size int, glyph_pair Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetKerning(font_rid RID, size int, glyph_pair Vector2i, kerning Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetKerning(font_rid RID, size int, glyph_pair Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlyphIndex(font_rid RID, size int, char int, variation_selector int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetCharFromGlyphIndex(font_rid RID, size int, glyph_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontHasChar(font_rid RID, char int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetSupportedChars(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRenderRange(font_rid RID, size Vector2i, start int, end int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRenderGlyph(font_rid RID, size Vector2i, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontDrawGlyph(font_rid RID, canvas RID, size int, pos Vector2, index int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontDrawGlyphOutline(font_rid RID, canvas RID, size int, outline_size int, pos Vector2, index int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontIsLanguageSupported(font_rid RID, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetLanguageSupportOverride(font_rid RID, language String, supported bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetLanguageSupportOverride(font_rid RID, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRemoveLanguageSupportOverride(font_rid RID, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetLanguageSupportOverrides(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontIsScriptSupported(font_rid RID, script String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetScriptSupportOverride(font_rid RID, script String, supported bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetScriptSupportOverride(font_rid RID, script String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontRemoveScriptSupportOverride(font_rid RID, script String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetScriptSupportOverrides(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetOpentypeFeatureOverrides(font_rid RID, overrides Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetOpentypeFeatureOverrides(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSupportedFeatureList(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSupportedVariationList(font_rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontGetGlobalOversampling() { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFontSetGlobalOversampling(oversampling float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XGetHexCodeBoxSize(size int, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XDrawHexCodeBox(canvas RID, size int, pos Vector2, index int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XCreateShapedText(direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextClear(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSetDirection(shaped RID, direction TextServerDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetDirection(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetInferredDirection(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSetBidiOverride(shaped RID, override Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSetCustomPunctuation(shaped RID, punct String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetCustomPunctuation(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSetOrientation(shaped RID, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetOrientation(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSetPreserveInvalid(shaped RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetPreserveInvalid(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSetPreserveControl(shaped RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetPreserveControl(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSetSpacing(shaped RID, spacing TextServerSpacingType, value int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetSpacing(shaped RID, spacing TextServerSpacingType, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextAddString(shaped RID, text String, fonts RID, size int, opentype_features Dictionary, language String, meta Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextAddObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, length int, baseline float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextResizeObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, baseline float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedGetSpanCount(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedGetSpanMeta(shaped RID, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedSetSpanUpdateFont(shaped RID, index int, fonts RID, size int, opentype_features Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSubstr(shaped RID, start int, length int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetParent(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextFitToWidth(shaped RID, width float32, justification_flags TextServerJustificationFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextTabAlign(shaped RID, tab_stops PackedFloat32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextShape(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextUpdateBreaks(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextUpdateJustificationOps(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextIsReady(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetGlyphs(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextSortLogical(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetGlyphCount(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetRange(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetLineBreaksAdv(shaped RID, width PackedFloat32Array, start int, once bool, break_flags TextServerLineBreakFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetLineBreaks(shaped RID, width float32, start int, break_flags TextServerLineBreakFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetWordBreaks(shaped RID, grapheme_flags TextServerGraphemeFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetTrimPos(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetEllipsisPos(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetEllipsisGlyphCount(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetEllipsisGlyphs(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextOverrunTrimToWidth(shaped RID, width float32, trim_flags TextServerTextOverrunFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetObjects(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetObjectRect(shaped RID, key Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetSize(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetAscent(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetDescent(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetWidth(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetUnderlinePosition(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetUnderlineThickness(shaped RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetDominantDirectionInRange(shaped RID, start int, end int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetCarets(shaped RID, position int, caret *CaretInfo, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetSelection(shaped RID, start int, end int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextHitTestGrapheme(shaped RID, coord float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextHitTestPosition(shaped RID, coord float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextDraw(shaped RID, canvas RID, pos Vector2, clip_l float32, clip_r float32, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextDrawOutline(shaped RID, canvas RID, pos Vector2, clip_l float32, clip_r float32, outline_size int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextGetGraphemeBounds(shaped RID, pos int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextNextGraphemePos(shaped RID, pos int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XShapedTextPrevGraphemePos(shaped RID, pos int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XFormatNumber(string String, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XParseNumber(string String, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XPercentSign(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XStripDiacritics(string String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XIsValidIdentifier(string String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XStringGetWordBreaks(string String, language String, chars_per_line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XIsConfusable(string String, dict PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XSpoofCheck(string String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XStringToUpper(string String, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XStringToLower(string String, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XParseStructuredText(parser_type TextServerStructuredTextParser, args Array, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextServerExtension) XCleanup() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
