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

type ptrsForTextServerExtensionList struct {
	fnXHasFeature                             gdc.MethodBindPtr
	fnXGetName                                gdc.MethodBindPtr
	fnXGetFeatures                            gdc.MethodBindPtr
	fnXFreeRid                                gdc.MethodBindPtr
	fnXHas                                    gdc.MethodBindPtr
	fnXLoadSupportData                        gdc.MethodBindPtr
	fnXGetSupportDataFilename                 gdc.MethodBindPtr
	fnXGetSupportDataInfo                     gdc.MethodBindPtr
	fnXSaveSupportData                        gdc.MethodBindPtr
	fnXIsLocaleRightToLeft                    gdc.MethodBindPtr
	fnXNameToTag                              gdc.MethodBindPtr
	fnXTagToName                              gdc.MethodBindPtr
	fnXCreateFont                             gdc.MethodBindPtr
	fnXCreateFontLinkedVariation              gdc.MethodBindPtr
	fnXFontSetData                            gdc.MethodBindPtr
	fnXFontSetDataPtr                         gdc.MethodBindPtr
	fnXFontSetFaceIndex                       gdc.MethodBindPtr
	fnXFontGetFaceIndex                       gdc.MethodBindPtr
	fnXFontGetFaceCount                       gdc.MethodBindPtr
	fnXFontSetStyle                           gdc.MethodBindPtr
	fnXFontGetStyle                           gdc.MethodBindPtr
	fnXFontSetName                            gdc.MethodBindPtr
	fnXFontGetName                            gdc.MethodBindPtr
	fnXFontGetOtNameStrings                   gdc.MethodBindPtr
	fnXFontSetStyleName                       gdc.MethodBindPtr
	fnXFontGetStyleName                       gdc.MethodBindPtr
	fnXFontSetWeight                          gdc.MethodBindPtr
	fnXFontGetWeight                          gdc.MethodBindPtr
	fnXFontSetStretch                         gdc.MethodBindPtr
	fnXFontGetStretch                         gdc.MethodBindPtr
	fnXFontSetAntialiasing                    gdc.MethodBindPtr
	fnXFontGetAntialiasing                    gdc.MethodBindPtr
	fnXFontSetGenerateMipmaps                 gdc.MethodBindPtr
	fnXFontGetGenerateMipmaps                 gdc.MethodBindPtr
	fnXFontSetMultichannelSignedDistanceField gdc.MethodBindPtr
	fnXFontIsMultichannelSignedDistanceField  gdc.MethodBindPtr
	fnXFontSetMsdfPixelRange                  gdc.MethodBindPtr
	fnXFontGetMsdfPixelRange                  gdc.MethodBindPtr
	fnXFontSetMsdfSize                        gdc.MethodBindPtr
	fnXFontGetMsdfSize                        gdc.MethodBindPtr
	fnXFontSetFixedSize                       gdc.MethodBindPtr
	fnXFontGetFixedSize                       gdc.MethodBindPtr
	fnXFontSetFixedSizeScaleMode              gdc.MethodBindPtr
	fnXFontGetFixedSizeScaleMode              gdc.MethodBindPtr
	fnXFontSetAllowSystemFallback             gdc.MethodBindPtr
	fnXFontIsAllowSystemFallback              gdc.MethodBindPtr
	fnXFontSetForceAutohinter                 gdc.MethodBindPtr
	fnXFontIsForceAutohinter                  gdc.MethodBindPtr
	fnXFontSetHinting                         gdc.MethodBindPtr
	fnXFontGetHinting                         gdc.MethodBindPtr
	fnXFontSetSubpixelPositioning             gdc.MethodBindPtr
	fnXFontGetSubpixelPositioning             gdc.MethodBindPtr
	fnXFontSetEmbolden                        gdc.MethodBindPtr
	fnXFontGetEmbolden                        gdc.MethodBindPtr
	fnXFontSetSpacing                         gdc.MethodBindPtr
	fnXFontGetSpacing                         gdc.MethodBindPtr
	fnXFontSetTransform                       gdc.MethodBindPtr
	fnXFontGetTransform                       gdc.MethodBindPtr
	fnXFontSetVariationCoordinates            gdc.MethodBindPtr
	fnXFontGetVariationCoordinates            gdc.MethodBindPtr
	fnXFontSetOversampling                    gdc.MethodBindPtr
	fnXFontGetOversampling                    gdc.MethodBindPtr
	fnXFontGetSizeCacheList                   gdc.MethodBindPtr
	fnXFontClearSizeCache                     gdc.MethodBindPtr
	fnXFontRemoveSizeCache                    gdc.MethodBindPtr
	fnXFontSetAscent                          gdc.MethodBindPtr
	fnXFontGetAscent                          gdc.MethodBindPtr
	fnXFontSetDescent                         gdc.MethodBindPtr
	fnXFontGetDescent                         gdc.MethodBindPtr
	fnXFontSetUnderlinePosition               gdc.MethodBindPtr
	fnXFontGetUnderlinePosition               gdc.MethodBindPtr
	fnXFontSetUnderlineThickness              gdc.MethodBindPtr
	fnXFontGetUnderlineThickness              gdc.MethodBindPtr
	fnXFontSetScale                           gdc.MethodBindPtr
	fnXFontGetScale                           gdc.MethodBindPtr
	fnXFontGetTextureCount                    gdc.MethodBindPtr
	fnXFontClearTextures                      gdc.MethodBindPtr
	fnXFontRemoveTexture                      gdc.MethodBindPtr
	fnXFontSetTextureImage                    gdc.MethodBindPtr
	fnXFontGetTextureImage                    gdc.MethodBindPtr
	fnXFontSetTextureOffsets                  gdc.MethodBindPtr
	fnXFontGetTextureOffsets                  gdc.MethodBindPtr
	fnXFontGetGlyphList                       gdc.MethodBindPtr
	fnXFontClearGlyphs                        gdc.MethodBindPtr
	fnXFontRemoveGlyph                        gdc.MethodBindPtr
	fnXFontGetGlyphAdvance                    gdc.MethodBindPtr
	fnXFontSetGlyphAdvance                    gdc.MethodBindPtr
	fnXFontGetGlyphOffset                     gdc.MethodBindPtr
	fnXFontSetGlyphOffset                     gdc.MethodBindPtr
	fnXFontGetGlyphSize                       gdc.MethodBindPtr
	fnXFontSetGlyphSize                       gdc.MethodBindPtr
	fnXFontGetGlyphUvRect                     gdc.MethodBindPtr
	fnXFontSetGlyphUvRect                     gdc.MethodBindPtr
	fnXFontGetGlyphTextureIdx                 gdc.MethodBindPtr
	fnXFontSetGlyphTextureIdx                 gdc.MethodBindPtr
	fnXFontGetGlyphTextureRid                 gdc.MethodBindPtr
	fnXFontGetGlyphTextureSize                gdc.MethodBindPtr
	fnXFontGetGlyphContours                   gdc.MethodBindPtr
	fnXFontGetKerningList                     gdc.MethodBindPtr
	fnXFontClearKerningMap                    gdc.MethodBindPtr
	fnXFontRemoveKerning                      gdc.MethodBindPtr
	fnXFontSetKerning                         gdc.MethodBindPtr
	fnXFontGetKerning                         gdc.MethodBindPtr
	fnXFontGetGlyphIndex                      gdc.MethodBindPtr
	fnXFontGetCharFromGlyphIndex              gdc.MethodBindPtr
	fnXFontHasChar                            gdc.MethodBindPtr
	fnXFontGetSupportedChars                  gdc.MethodBindPtr
	fnXFontRenderRange                        gdc.MethodBindPtr
	fnXFontRenderGlyph                        gdc.MethodBindPtr
	fnXFontDrawGlyph                          gdc.MethodBindPtr
	fnXFontDrawGlyphOutline                   gdc.MethodBindPtr
	fnXFontIsLanguageSupported                gdc.MethodBindPtr
	fnXFontSetLanguageSupportOverride         gdc.MethodBindPtr
	fnXFontGetLanguageSupportOverride         gdc.MethodBindPtr
	fnXFontRemoveLanguageSupportOverride      gdc.MethodBindPtr
	fnXFontGetLanguageSupportOverrides        gdc.MethodBindPtr
	fnXFontIsScriptSupported                  gdc.MethodBindPtr
	fnXFontSetScriptSupportOverride           gdc.MethodBindPtr
	fnXFontGetScriptSupportOverride           gdc.MethodBindPtr
	fnXFontRemoveScriptSupportOverride        gdc.MethodBindPtr
	fnXFontGetScriptSupportOverrides          gdc.MethodBindPtr
	fnXFontSetOpentypeFeatureOverrides        gdc.MethodBindPtr
	fnXFontGetOpentypeFeatureOverrides        gdc.MethodBindPtr
	fnXFontSupportedFeatureList               gdc.MethodBindPtr
	fnXFontSupportedVariationList             gdc.MethodBindPtr
	fnXFontGetGlobalOversampling              gdc.MethodBindPtr
	fnXFontSetGlobalOversampling              gdc.MethodBindPtr
	fnXGetHexCodeBoxSize                      gdc.MethodBindPtr
	fnXDrawHexCodeBox                         gdc.MethodBindPtr
	fnXCreateShapedText                       gdc.MethodBindPtr
	fnXShapedTextClear                        gdc.MethodBindPtr
	fnXShapedTextSetDirection                 gdc.MethodBindPtr
	fnXShapedTextGetDirection                 gdc.MethodBindPtr
	fnXShapedTextGetInferredDirection         gdc.MethodBindPtr
	fnXShapedTextSetBidiOverride              gdc.MethodBindPtr
	fnXShapedTextSetCustomPunctuation         gdc.MethodBindPtr
	fnXShapedTextGetCustomPunctuation         gdc.MethodBindPtr
	fnXShapedTextSetOrientation               gdc.MethodBindPtr
	fnXShapedTextGetOrientation               gdc.MethodBindPtr
	fnXShapedTextSetPreserveInvalid           gdc.MethodBindPtr
	fnXShapedTextGetPreserveInvalid           gdc.MethodBindPtr
	fnXShapedTextSetPreserveControl           gdc.MethodBindPtr
	fnXShapedTextGetPreserveControl           gdc.MethodBindPtr
	fnXShapedTextSetSpacing                   gdc.MethodBindPtr
	fnXShapedTextGetSpacing                   gdc.MethodBindPtr
	fnXShapedTextAddString                    gdc.MethodBindPtr
	fnXShapedTextAddObject                    gdc.MethodBindPtr
	fnXShapedTextResizeObject                 gdc.MethodBindPtr
	fnXShapedGetSpanCount                     gdc.MethodBindPtr
	fnXShapedGetSpanMeta                      gdc.MethodBindPtr
	fnXShapedSetSpanUpdateFont                gdc.MethodBindPtr
	fnXShapedTextSubstr                       gdc.MethodBindPtr
	fnXShapedTextGetParent                    gdc.MethodBindPtr
	fnXShapedTextFitToWidth                   gdc.MethodBindPtr
	fnXShapedTextTabAlign                     gdc.MethodBindPtr
	fnXShapedTextShape                        gdc.MethodBindPtr
	fnXShapedTextUpdateBreaks                 gdc.MethodBindPtr
	fnXShapedTextUpdateJustificationOps       gdc.MethodBindPtr
	fnXShapedTextIsReady                      gdc.MethodBindPtr
	fnXShapedTextGetGlyphs                    gdc.MethodBindPtr
	fnXShapedTextSortLogical                  gdc.MethodBindPtr
	fnXShapedTextGetGlyphCount                gdc.MethodBindPtr
	fnXShapedTextGetRange                     gdc.MethodBindPtr
	fnXShapedTextGetLineBreaksAdv             gdc.MethodBindPtr
	fnXShapedTextGetLineBreaks                gdc.MethodBindPtr
	fnXShapedTextGetWordBreaks                gdc.MethodBindPtr
	fnXShapedTextGetTrimPos                   gdc.MethodBindPtr
	fnXShapedTextGetEllipsisPos               gdc.MethodBindPtr
	fnXShapedTextGetEllipsisGlyphCount        gdc.MethodBindPtr
	fnXShapedTextGetEllipsisGlyphs            gdc.MethodBindPtr
	fnXShapedTextOverrunTrimToWidth           gdc.MethodBindPtr
	fnXShapedTextGetObjects                   gdc.MethodBindPtr
	fnXShapedTextGetObjectRect                gdc.MethodBindPtr
	fnXShapedTextGetSize                      gdc.MethodBindPtr
	fnXShapedTextGetAscent                    gdc.MethodBindPtr
	fnXShapedTextGetDescent                   gdc.MethodBindPtr
	fnXShapedTextGetWidth                     gdc.MethodBindPtr
	fnXShapedTextGetUnderlinePosition         gdc.MethodBindPtr
	fnXShapedTextGetUnderlineThickness        gdc.MethodBindPtr
	fnXShapedTextGetDominantDirectionInRange  gdc.MethodBindPtr
	fnXShapedTextGetCarets                    gdc.MethodBindPtr
	fnXShapedTextGetSelection                 gdc.MethodBindPtr
	fnXShapedTextHitTestGrapheme              gdc.MethodBindPtr
	fnXShapedTextHitTestPosition              gdc.MethodBindPtr
	fnXShapedTextDraw                         gdc.MethodBindPtr
	fnXShapedTextDrawOutline                  gdc.MethodBindPtr
	fnXShapedTextGetGraphemeBounds            gdc.MethodBindPtr
	fnXShapedTextNextGraphemePos              gdc.MethodBindPtr
	fnXShapedTextPrevGraphemePos              gdc.MethodBindPtr
	fnXShapedTextGetCharacterBreaks           gdc.MethodBindPtr
	fnXShapedTextNextCharacterPos             gdc.MethodBindPtr
	fnXShapedTextPrevCharacterPos             gdc.MethodBindPtr
	fnXShapedTextClosestCharacterPos          gdc.MethodBindPtr
	fnXFormatNumber                           gdc.MethodBindPtr
	fnXParseNumber                            gdc.MethodBindPtr
	fnXPercentSign                            gdc.MethodBindPtr
	fnXStripDiacritics                        gdc.MethodBindPtr
	fnXIsValidIdentifier                      gdc.MethodBindPtr
	fnXStringGetWordBreaks                    gdc.MethodBindPtr
	fnXStringGetCharacterBreaks               gdc.MethodBindPtr
	fnXIsConfusable                           gdc.MethodBindPtr
	fnXSpoofCheck                             gdc.MethodBindPtr
	fnXStringToUpper                          gdc.MethodBindPtr
	fnXStringToLower                          gdc.MethodBindPtr
	fnXParseStructuredText                    gdc.MethodBindPtr
	fnXCleanup                                gdc.MethodBindPtr
}

var ptrsForTextServerExtension ptrsForTextServerExtensionList

func initTextServerExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextServerExtension")
	defer className.Destroy()
}

type TextServerExtension struct {
	TextServer
}

func (me *TextServerExtension) BaseClass() string {
	return "TextServerExtension"
}

func NewTextServerExtension() *TextServerExtension {
	str := StringNameFromStr("TextServerExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextServerExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TextServerExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextServerExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextServerExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
