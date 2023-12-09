// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextServer struct {
  obj gdc.ObjectPtr
}

func (me *TextServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextServer) BaseClass() string {
  return "TextServer"
}



// Enums

type TextServerFontAntialiasing int
const (
  TextServerFontAntialiasingFontAntialiasingNone TextServerFontAntialiasing = 0
  TextServerFontAntialiasingFontAntialiasingGray TextServerFontAntialiasing = 1
  TextServerFontAntialiasingFontAntialiasingLcd TextServerFontAntialiasing = 2
)

type TextServerFontLCDSubpixelLayout int
const (
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutNone TextServerFontLCDSubpixelLayout = 0
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutHrgb TextServerFontLCDSubpixelLayout = 1
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutHbgr TextServerFontLCDSubpixelLayout = 2
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutVrgb TextServerFontLCDSubpixelLayout = 3
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutVbgr TextServerFontLCDSubpixelLayout = 4
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutMax TextServerFontLCDSubpixelLayout = 5
)

type TextServerDirection int
const (
  TextServerDirectionDirectionAuto TextServerDirection = 0
  TextServerDirectionDirectionLtr TextServerDirection = 1
  TextServerDirectionDirectionRtl TextServerDirection = 2
  TextServerDirectionDirectionInherited TextServerDirection = 3
)

type TextServerOrientation int
const (
  TextServerOrientationOrientationHorizontal TextServerOrientation = 0
  TextServerOrientationOrientationVertical TextServerOrientation = 1
)

type TextServerJustificationFlag int
const (
  TextServerJustificationFlagJustificationNone TextServerJustificationFlag = 0
  TextServerJustificationFlagJustificationKashida TextServerJustificationFlag = 1
  TextServerJustificationFlagJustificationWordBound TextServerJustificationFlag = 2
  TextServerJustificationFlagJustificationTrimEdgeSpaces TextServerJustificationFlag = 4
  TextServerJustificationFlagJustificationAfterLastTab TextServerJustificationFlag = 8
  TextServerJustificationFlagJustificationConstrainEllipsis TextServerJustificationFlag = 16
  TextServerJustificationFlagJustificationSkipLastLine TextServerJustificationFlag = 32
  TextServerJustificationFlagJustificationSkipLastLineWithVisibleChars TextServerJustificationFlag = 64
  TextServerJustificationFlagJustificationDoNotSkipSingleLine TextServerJustificationFlag = 128
)

type TextServerAutowrapMode int
const (
  TextServerAutowrapModeAutowrapOff TextServerAutowrapMode = 0
  TextServerAutowrapModeAutowrapArbitrary TextServerAutowrapMode = 1
  TextServerAutowrapModeAutowrapWord TextServerAutowrapMode = 2
  TextServerAutowrapModeAutowrapWordSmart TextServerAutowrapMode = 3
)

type TextServerLineBreakFlag int
const (
  TextServerLineBreakFlagBreakNone TextServerLineBreakFlag = 0
  TextServerLineBreakFlagBreakMandatory TextServerLineBreakFlag = 1
  TextServerLineBreakFlagBreakWordBound TextServerLineBreakFlag = 2
  TextServerLineBreakFlagBreakGraphemeBound TextServerLineBreakFlag = 4
  TextServerLineBreakFlagBreakAdaptive TextServerLineBreakFlag = 8
  TextServerLineBreakFlagBreakTrimEdgeSpaces TextServerLineBreakFlag = 16
)

type TextServerVisibleCharactersBehavior int
const (
  TextServerVisibleCharactersBehaviorVcCharsBeforeShaping TextServerVisibleCharactersBehavior = 0
  TextServerVisibleCharactersBehaviorVcCharsAfterShaping TextServerVisibleCharactersBehavior = 1
  TextServerVisibleCharactersBehaviorVcGlyphsAuto TextServerVisibleCharactersBehavior = 2
  TextServerVisibleCharactersBehaviorVcGlyphsLtr TextServerVisibleCharactersBehavior = 3
  TextServerVisibleCharactersBehaviorVcGlyphsRtl TextServerVisibleCharactersBehavior = 4
)

type TextServerOverrunBehavior int
const (
  TextServerOverrunBehaviorOverrunNoTrimming TextServerOverrunBehavior = 0
  TextServerOverrunBehaviorOverrunTrimChar TextServerOverrunBehavior = 1
  TextServerOverrunBehaviorOverrunTrimWord TextServerOverrunBehavior = 2
  TextServerOverrunBehaviorOverrunTrimEllipsis TextServerOverrunBehavior = 3
  TextServerOverrunBehaviorOverrunTrimWordEllipsis TextServerOverrunBehavior = 4
)

type TextServerTextOverrunFlag int
const (
  TextServerTextOverrunFlagOverrunNoTrim TextServerTextOverrunFlag = 0
  TextServerTextOverrunFlagOverrunTrim TextServerTextOverrunFlag = 1
  TextServerTextOverrunFlagOverrunTrimWordOnly TextServerTextOverrunFlag = 2
  TextServerTextOverrunFlagOverrunAddEllipsis TextServerTextOverrunFlag = 4
  TextServerTextOverrunFlagOverrunEnforceEllipsis TextServerTextOverrunFlag = 8
  TextServerTextOverrunFlagOverrunJustificationAware TextServerTextOverrunFlag = 16
)

type TextServerGraphemeFlag int
const (
  TextServerGraphemeFlagGraphemeIsValid TextServerGraphemeFlag = 1
  TextServerGraphemeFlagGraphemeIsRtl TextServerGraphemeFlag = 2
  TextServerGraphemeFlagGraphemeIsVirtual TextServerGraphemeFlag = 4
  TextServerGraphemeFlagGraphemeIsSpace TextServerGraphemeFlag = 8
  TextServerGraphemeFlagGraphemeIsBreakHard TextServerGraphemeFlag = 16
  TextServerGraphemeFlagGraphemeIsBreakSoft TextServerGraphemeFlag = 32
  TextServerGraphemeFlagGraphemeIsTab TextServerGraphemeFlag = 64
  TextServerGraphemeFlagGraphemeIsElongation TextServerGraphemeFlag = 128
  TextServerGraphemeFlagGraphemeIsPunctuation TextServerGraphemeFlag = 256
  TextServerGraphemeFlagGraphemeIsUnderscore TextServerGraphemeFlag = 512
  TextServerGraphemeFlagGraphemeIsConnected TextServerGraphemeFlag = 1024
  TextServerGraphemeFlagGraphemeIsSafeToInsertTatweel TextServerGraphemeFlag = 2048
  TextServerGraphemeFlagGraphemeIsEmbeddedObject TextServerGraphemeFlag = 4096
)

type TextServerHinting int
const (
  TextServerHintingHintingNone TextServerHinting = 0
  TextServerHintingHintingLight TextServerHinting = 1
  TextServerHintingHintingNormal TextServerHinting = 2
)

type TextServerSubpixelPositioning int
const (
  TextServerSubpixelPositioningSubpixelPositioningDisabled TextServerSubpixelPositioning = 0
  TextServerSubpixelPositioningSubpixelPositioningAuto TextServerSubpixelPositioning = 1
  TextServerSubpixelPositioningSubpixelPositioningOneHalf TextServerSubpixelPositioning = 2
  TextServerSubpixelPositioningSubpixelPositioningOneQuarter TextServerSubpixelPositioning = 3
  TextServerSubpixelPositioningSubpixelPositioningOneHalfMaxSize TextServerSubpixelPositioning = 20
  TextServerSubpixelPositioningSubpixelPositioningOneQuarterMaxSize TextServerSubpixelPositioning = 16
)

type TextServerFeature int
const (
  TextServerFeatureFeatureSimpleLayout TextServerFeature = 1
  TextServerFeatureFeatureBidiLayout TextServerFeature = 2
  TextServerFeatureFeatureVerticalLayout TextServerFeature = 4
  TextServerFeatureFeatureShaping TextServerFeature = 8
  TextServerFeatureFeatureKashidaJustification TextServerFeature = 16
  TextServerFeatureFeatureBreakIterators TextServerFeature = 32
  TextServerFeatureFeatureFontBitmap TextServerFeature = 64
  TextServerFeatureFeatureFontDynamic TextServerFeature = 128
  TextServerFeatureFeatureFontMsdf TextServerFeature = 256
  TextServerFeatureFeatureFontSystem TextServerFeature = 512
  TextServerFeatureFeatureFontVariable TextServerFeature = 1024
  TextServerFeatureFeatureContextSensitiveCaseConversion TextServerFeature = 2048
  TextServerFeatureFeatureUseSupportData TextServerFeature = 4096
  TextServerFeatureFeatureUnicodeIdentifiers TextServerFeature = 8192
  TextServerFeatureFeatureUnicodeSecurity TextServerFeature = 16384
)

type TextServerContourPointTag int
const (
  TextServerContourPointTagContourCurveTagOn TextServerContourPointTag = 1
  TextServerContourPointTagContourCurveTagOffConic TextServerContourPointTag = 0
  TextServerContourPointTagContourCurveTagOffCubic TextServerContourPointTag = 2
)

type TextServerSpacingType int
const (
  TextServerSpacingTypeSpacingGlyph TextServerSpacingType = 0
  TextServerSpacingTypeSpacingSpace TextServerSpacingType = 1
  TextServerSpacingTypeSpacingTop TextServerSpacingType = 2
  TextServerSpacingTypeSpacingBottom TextServerSpacingType = 3
  TextServerSpacingTypeSpacingMax TextServerSpacingType = 4
)

type TextServerFontStyle int
const (
  TextServerFontStyleFontBold TextServerFontStyle = 1
  TextServerFontStyleFontItalic TextServerFontStyle = 2
  TextServerFontStyleFontFixedWidth TextServerFontStyle = 4
)

type TextServerStructuredTextParser int
const (
  TextServerStructuredTextParserStructuredTextDefault TextServerStructuredTextParser = 0
  TextServerStructuredTextParserStructuredTextUri TextServerStructuredTextParser = 1
  TextServerStructuredTextParserStructuredTextFile TextServerStructuredTextParser = 2
  TextServerStructuredTextParserStructuredTextEmail TextServerStructuredTextParser = 3
  TextServerStructuredTextParserStructuredTextList TextServerStructuredTextParser = 4
  TextServerStructuredTextParserStructuredTextGdscript TextServerStructuredTextParser = 5
  TextServerStructuredTextParserStructuredTextCustom TextServerStructuredTextParser = 6
)

func (me *TextServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TextServer) HasFeature(feature TextServerFeature, )  {
  panic("TODO: implement")
}

func  (me *TextServer) GetName()  {
  panic("TODO: implement")
}

func  (me *TextServer) GetFeatures()  {
  panic("TODO: implement")
}

func  (me *TextServer) LoadSupportData(filename String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) GetSupportDataFilename()  {
  panic("TODO: implement")
}

func  (me *TextServer) GetSupportDataInfo()  {
  panic("TODO: implement")
}

func  (me *TextServer) SaveSupportData(filename String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) IsLocaleRightToLeft(locale String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) NameToTag(name String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) TagToName(tag int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) Has(rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FreeRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) CreateFont()  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetData(font_rid RID, data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetFaceIndex(font_rid RID, face_index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetFaceIndex(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetFaceCount(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetStyle(font_rid RID, style TextServerFontStyle, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetStyle(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetName(font_rid RID, name String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetName(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetOtNameStrings(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetStyleName(font_rid RID, name String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetStyleName(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetWeight(font_rid RID, weight int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetWeight(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetStretch(font_rid RID, weight int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetStretch(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetAntialiasing(font_rid RID, antialiasing TextServerFontAntialiasing, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetAntialiasing(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetGenerateMipmaps(font_rid RID, generate_mipmaps bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGenerateMipmaps(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetMultichannelSignedDistanceField(font_rid RID, msdf bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontIsMultichannelSignedDistanceField(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetMsdfPixelRange(font_rid RID, msdf_pixel_range int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetMsdfPixelRange(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetMsdfSize(font_rid RID, msdf_size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetMsdfSize(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetFixedSize(font_rid RID, fixed_size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetFixedSize(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetAllowSystemFallback(font_rid RID, allow_system_fallback bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontIsAllowSystemFallback(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetForceAutohinter(font_rid RID, force_autohinter bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontIsForceAutohinter(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetHinting(font_rid RID, hinting TextServerHinting, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetHinting(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetSubpixelPositioning(font_rid RID, subpixel_positioning TextServerSubpixelPositioning, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetSubpixelPositioning(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetEmbolden(font_rid RID, strength float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetEmbolden(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetTransform(font_rid RID, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetTransform(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetVariationCoordinates(font_rid RID, variation_coordinates Dictionary, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetVariationCoordinates(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetOversampling(font_rid RID, oversampling float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetOversampling(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetSizeCacheList(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontClearSizeCache(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRemoveSizeCache(font_rid RID, size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetAscent(font_rid RID, size int, ascent float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetAscent(font_rid RID, size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetDescent(font_rid RID, size int, descent float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetDescent(font_rid RID, size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetUnderlinePosition(font_rid RID, size int, underline_position float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetUnderlinePosition(font_rid RID, size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetUnderlineThickness(font_rid RID, size int, underline_thickness float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetUnderlineThickness(font_rid RID, size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetScale(font_rid RID, size int, scale float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetScale(font_rid RID, size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetTextureCount(font_rid RID, size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontClearTextures(font_rid RID, size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRemoveTexture(font_rid RID, size Vector2i, texture_index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetTextureImage(font_rid RID, size Vector2i, texture_index int, image Image, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetTextureImage(font_rid RID, size Vector2i, texture_index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetTextureOffsets(font_rid RID, size Vector2i, texture_index int, offset PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetTextureOffsets(font_rid RID, size Vector2i, texture_index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphList(font_rid RID, size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontClearGlyphs(font_rid RID, size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRemoveGlyph(font_rid RID, size Vector2i, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphAdvance(font_rid RID, size int, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetGlyphAdvance(font_rid RID, size int, glyph int, advance Vector2, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphOffset(font_rid RID, size Vector2i, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetGlyphOffset(font_rid RID, size Vector2i, glyph int, offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphSize(font_rid RID, size Vector2i, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetGlyphSize(font_rid RID, size Vector2i, glyph int, gl_size Vector2, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphUvRect(font_rid RID, size Vector2i, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetGlyphUvRect(font_rid RID, size Vector2i, glyph int, uv_rect Rect2, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int, texture_idx int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphTextureRid(font_rid RID, size Vector2i, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphTextureSize(font_rid RID, size Vector2i, glyph int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphContours(font RID, size int, index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetKerningList(font_rid RID, size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontClearKerningMap(font_rid RID, size int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRemoveKerning(font_rid RID, size int, glyph_pair Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetKerning(font_rid RID, size int, glyph_pair Vector2i, kerning Vector2, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetKerning(font_rid RID, size int, glyph_pair Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlyphIndex(font_rid RID, size int, char int, variation_selector int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetCharFromGlyphIndex(font_rid RID, size int, glyph_index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontHasChar(font_rid RID, char int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetSupportedChars(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRenderRange(font_rid RID, size Vector2i, start int, end int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRenderGlyph(font_rid RID, size Vector2i, index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontDrawGlyph(font_rid RID, canvas RID, size int, pos Vector2, index int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontDrawGlyphOutline(font_rid RID, canvas RID, size int, outline_size int, pos Vector2, index int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontIsLanguageSupported(font_rid RID, language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetLanguageSupportOverride(font_rid RID, language String, supported bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetLanguageSupportOverride(font_rid RID, language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRemoveLanguageSupportOverride(font_rid RID, language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetLanguageSupportOverrides(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontIsScriptSupported(font_rid RID, script String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetScriptSupportOverride(font_rid RID, script String, supported bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetScriptSupportOverride(font_rid RID, script String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontRemoveScriptSupportOverride(font_rid RID, script String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetScriptSupportOverrides(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetOpentypeFeatureOverrides(font_rid RID, overrides Dictionary, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetOpentypeFeatureOverrides(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSupportedFeatureList(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSupportedVariationList(font_rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FontGetGlobalOversampling()  {
  panic("TODO: implement")
}

func  (me *TextServer) FontSetGlobalOversampling(oversampling float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) GetHexCodeBoxSize(size int, index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) DrawHexCodeBox(canvas RID, size int, pos Vector2, index int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextServer) CreateShapedText(direction TextServerDirection, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextClear(rid RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSetDirection(shaped RID, direction TextServerDirection, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetDirection(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetInferredDirection(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSetBidiOverride(shaped RID, override Array, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSetCustomPunctuation(shaped RID, punct String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetCustomPunctuation(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSetOrientation(shaped RID, orientation TextServerOrientation, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetOrientation(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSetPreserveInvalid(shaped RID, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetPreserveInvalid(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSetPreserveControl(shaped RID, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetPreserveControl(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSetSpacing(shaped RID, spacing TextServerSpacingType, value int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetSpacing(shaped RID, spacing TextServerSpacingType, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextAddString(shaped RID, text String, fonts RID, size int, opentype_features Dictionary, language String, meta Variant, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextAddObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, length int, baseline float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextResizeObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, baseline float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedGetSpanCount(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedGetSpanMeta(shaped RID, index int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedSetSpanUpdateFont(shaped RID, index int, fonts RID, size int, opentype_features Dictionary, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSubstr(shaped RID, start int, length int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetParent(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextFitToWidth(shaped RID, width float32, justification_flags TextServerJustificationFlag, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextTabAlign(shaped RID, tab_stops PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextShape(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextIsReady(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextHasVisibleChars(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetGlyphs(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextSortLogical(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetGlyphCount(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetRange(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetLineBreaksAdv(shaped RID, width PackedFloat32Array, start int, once bool, break_flags TextServerLineBreakFlag, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetLineBreaks(shaped RID, width float32, start int, break_flags TextServerLineBreakFlag, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetWordBreaks(shaped RID, grapheme_flags TextServerGraphemeFlag, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetTrimPos(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetEllipsisPos(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetEllipsisGlyphs(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetEllipsisGlyphCount(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextOverrunTrimToWidth(shaped RID, width float32, overrun_trim_flags TextServerTextOverrunFlag, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetObjects(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetObjectRect(shaped RID, key Variant, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetSize(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetAscent(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetDescent(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetWidth(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetUnderlinePosition(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetUnderlineThickness(shaped RID, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetCarets(shaped RID, position int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetSelection(shaped RID, start int, end int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextHitTestGrapheme(shaped RID, coords float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextHitTestPosition(shaped RID, coords float32, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetGraphemeBounds(shaped RID, pos int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextNextGraphemePos(shaped RID, pos int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextPrevGraphemePos(shaped RID, pos int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextDraw(shaped RID, canvas RID, pos Vector2, clip_l float32, clip_r float32, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextDrawOutline(shaped RID, canvas RID, pos Vector2, clip_l float32, clip_r float32, outline_size int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ShapedTextGetDominantDirectionInRange(shaped RID, start int, end int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) FormatNumber(number String, language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ParseNumber(number String, language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) PercentSign(language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) StringGetWordBreaks(string_ String, language String, chars_per_line int, )  {
  panic("TODO: implement")
}

func  (me *TextServer) IsConfusable(string_ String, dict PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *TextServer) SpoofCheck(string_ String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) StripDiacritics(string_ String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) IsValidIdentifier(string_ String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) StringToUpper(string_ String, language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) StringToLower(string_ String, language String, )  {
  panic("TODO: implement")
}

func  (me *TextServer) ParseStructuredText(parser_type TextServerStructuredTextParser, args Array, text String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
