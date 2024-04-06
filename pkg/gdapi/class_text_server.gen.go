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

type ptrsForTextServerList struct {
	fnHasFeature                             gdc.MethodBindPtr
	fnGetName                                gdc.MethodBindPtr
	fnGetFeatures                            gdc.MethodBindPtr
	fnLoadSupportData                        gdc.MethodBindPtr
	fnGetSupportDataFilename                 gdc.MethodBindPtr
	fnGetSupportDataInfo                     gdc.MethodBindPtr
	fnSaveSupportData                        gdc.MethodBindPtr
	fnIsLocaleRightToLeft                    gdc.MethodBindPtr
	fnNameToTag                              gdc.MethodBindPtr
	fnTagToName                              gdc.MethodBindPtr
	fnHas                                    gdc.MethodBindPtr
	fnFreeRid                                gdc.MethodBindPtr
	fnCreateFont                             gdc.MethodBindPtr
	fnCreateFontLinkedVariation              gdc.MethodBindPtr
	fnFontSetData                            gdc.MethodBindPtr
	fnFontSetFaceIndex                       gdc.MethodBindPtr
	fnFontGetFaceIndex                       gdc.MethodBindPtr
	fnFontGetFaceCount                       gdc.MethodBindPtr
	fnFontSetStyle                           gdc.MethodBindPtr
	fnFontGetStyle                           gdc.MethodBindPtr
	fnFontSetName                            gdc.MethodBindPtr
	fnFontGetName                            gdc.MethodBindPtr
	fnFontGetOtNameStrings                   gdc.MethodBindPtr
	fnFontSetStyleName                       gdc.MethodBindPtr
	fnFontGetStyleName                       gdc.MethodBindPtr
	fnFontSetWeight                          gdc.MethodBindPtr
	fnFontGetWeight                          gdc.MethodBindPtr
	fnFontSetStretch                         gdc.MethodBindPtr
	fnFontGetStretch                         gdc.MethodBindPtr
	fnFontSetAntialiasing                    gdc.MethodBindPtr
	fnFontGetAntialiasing                    gdc.MethodBindPtr
	fnFontSetGenerateMipmaps                 gdc.MethodBindPtr
	fnFontGetGenerateMipmaps                 gdc.MethodBindPtr
	fnFontSetMultichannelSignedDistanceField gdc.MethodBindPtr
	fnFontIsMultichannelSignedDistanceField  gdc.MethodBindPtr
	fnFontSetMsdfPixelRange                  gdc.MethodBindPtr
	fnFontGetMsdfPixelRange                  gdc.MethodBindPtr
	fnFontSetMsdfSize                        gdc.MethodBindPtr
	fnFontGetMsdfSize                        gdc.MethodBindPtr
	fnFontSetFixedSize                       gdc.MethodBindPtr
	fnFontGetFixedSize                       gdc.MethodBindPtr
	fnFontSetFixedSizeScaleMode              gdc.MethodBindPtr
	fnFontGetFixedSizeScaleMode              gdc.MethodBindPtr
	fnFontSetAllowSystemFallback             gdc.MethodBindPtr
	fnFontIsAllowSystemFallback              gdc.MethodBindPtr
	fnFontSetForceAutohinter                 gdc.MethodBindPtr
	fnFontIsForceAutohinter                  gdc.MethodBindPtr
	fnFontSetHinting                         gdc.MethodBindPtr
	fnFontGetHinting                         gdc.MethodBindPtr
	fnFontSetSubpixelPositioning             gdc.MethodBindPtr
	fnFontGetSubpixelPositioning             gdc.MethodBindPtr
	fnFontSetEmbolden                        gdc.MethodBindPtr
	fnFontGetEmbolden                        gdc.MethodBindPtr
	fnFontSetSpacing                         gdc.MethodBindPtr
	fnFontGetSpacing                         gdc.MethodBindPtr
	fnFontSetTransform                       gdc.MethodBindPtr
	fnFontGetTransform                       gdc.MethodBindPtr
	fnFontSetVariationCoordinates            gdc.MethodBindPtr
	fnFontGetVariationCoordinates            gdc.MethodBindPtr
	fnFontSetOversampling                    gdc.MethodBindPtr
	fnFontGetOversampling                    gdc.MethodBindPtr
	fnFontGetSizeCacheList                   gdc.MethodBindPtr
	fnFontClearSizeCache                     gdc.MethodBindPtr
	fnFontRemoveSizeCache                    gdc.MethodBindPtr
	fnFontSetAscent                          gdc.MethodBindPtr
	fnFontGetAscent                          gdc.MethodBindPtr
	fnFontSetDescent                         gdc.MethodBindPtr
	fnFontGetDescent                         gdc.MethodBindPtr
	fnFontSetUnderlinePosition               gdc.MethodBindPtr
	fnFontGetUnderlinePosition               gdc.MethodBindPtr
	fnFontSetUnderlineThickness              gdc.MethodBindPtr
	fnFontGetUnderlineThickness              gdc.MethodBindPtr
	fnFontSetScale                           gdc.MethodBindPtr
	fnFontGetScale                           gdc.MethodBindPtr
	fnFontGetTextureCount                    gdc.MethodBindPtr
	fnFontClearTextures                      gdc.MethodBindPtr
	fnFontRemoveTexture                      gdc.MethodBindPtr
	fnFontSetTextureImage                    gdc.MethodBindPtr
	fnFontGetTextureImage                    gdc.MethodBindPtr
	fnFontSetTextureOffsets                  gdc.MethodBindPtr
	fnFontGetTextureOffsets                  gdc.MethodBindPtr
	fnFontGetGlyphList                       gdc.MethodBindPtr
	fnFontClearGlyphs                        gdc.MethodBindPtr
	fnFontRemoveGlyph                        gdc.MethodBindPtr
	fnFontGetGlyphAdvance                    gdc.MethodBindPtr
	fnFontSetGlyphAdvance                    gdc.MethodBindPtr
	fnFontGetGlyphOffset                     gdc.MethodBindPtr
	fnFontSetGlyphOffset                     gdc.MethodBindPtr
	fnFontGetGlyphSize                       gdc.MethodBindPtr
	fnFontSetGlyphSize                       gdc.MethodBindPtr
	fnFontGetGlyphUvRect                     gdc.MethodBindPtr
	fnFontSetGlyphUvRect                     gdc.MethodBindPtr
	fnFontGetGlyphTextureIdx                 gdc.MethodBindPtr
	fnFontSetGlyphTextureIdx                 gdc.MethodBindPtr
	fnFontGetGlyphTextureRid                 gdc.MethodBindPtr
	fnFontGetGlyphTextureSize                gdc.MethodBindPtr
	fnFontGetGlyphContours                   gdc.MethodBindPtr
	fnFontGetKerningList                     gdc.MethodBindPtr
	fnFontClearKerningMap                    gdc.MethodBindPtr
	fnFontRemoveKerning                      gdc.MethodBindPtr
	fnFontSetKerning                         gdc.MethodBindPtr
	fnFontGetKerning                         gdc.MethodBindPtr
	fnFontGetGlyphIndex                      gdc.MethodBindPtr
	fnFontGetCharFromGlyphIndex              gdc.MethodBindPtr
	fnFontHasChar                            gdc.MethodBindPtr
	fnFontGetSupportedChars                  gdc.MethodBindPtr
	fnFontRenderRange                        gdc.MethodBindPtr
	fnFontRenderGlyph                        gdc.MethodBindPtr
	fnFontDrawGlyph                          gdc.MethodBindPtr
	fnFontDrawGlyphOutline                   gdc.MethodBindPtr
	fnFontIsLanguageSupported                gdc.MethodBindPtr
	fnFontSetLanguageSupportOverride         gdc.MethodBindPtr
	fnFontGetLanguageSupportOverride         gdc.MethodBindPtr
	fnFontRemoveLanguageSupportOverride      gdc.MethodBindPtr
	fnFontGetLanguageSupportOverrides        gdc.MethodBindPtr
	fnFontIsScriptSupported                  gdc.MethodBindPtr
	fnFontSetScriptSupportOverride           gdc.MethodBindPtr
	fnFontGetScriptSupportOverride           gdc.MethodBindPtr
	fnFontRemoveScriptSupportOverride        gdc.MethodBindPtr
	fnFontGetScriptSupportOverrides          gdc.MethodBindPtr
	fnFontSetOpentypeFeatureOverrides        gdc.MethodBindPtr
	fnFontGetOpentypeFeatureOverrides        gdc.MethodBindPtr
	fnFontSupportedFeatureList               gdc.MethodBindPtr
	fnFontSupportedVariationList             gdc.MethodBindPtr
	fnFontGetGlobalOversampling              gdc.MethodBindPtr
	fnFontSetGlobalOversampling              gdc.MethodBindPtr
	fnGetHexCodeBoxSize                      gdc.MethodBindPtr
	fnDrawHexCodeBox                         gdc.MethodBindPtr
	fnCreateShapedText                       gdc.MethodBindPtr
	fnShapedTextClear                        gdc.MethodBindPtr
	fnShapedTextSetDirection                 gdc.MethodBindPtr
	fnShapedTextGetDirection                 gdc.MethodBindPtr
	fnShapedTextGetInferredDirection         gdc.MethodBindPtr
	fnShapedTextSetBidiOverride              gdc.MethodBindPtr
	fnShapedTextSetCustomPunctuation         gdc.MethodBindPtr
	fnShapedTextGetCustomPunctuation         gdc.MethodBindPtr
	fnShapedTextSetOrientation               gdc.MethodBindPtr
	fnShapedTextGetOrientation               gdc.MethodBindPtr
	fnShapedTextSetPreserveInvalid           gdc.MethodBindPtr
	fnShapedTextGetPreserveInvalid           gdc.MethodBindPtr
	fnShapedTextSetPreserveControl           gdc.MethodBindPtr
	fnShapedTextGetPreserveControl           gdc.MethodBindPtr
	fnShapedTextSetSpacing                   gdc.MethodBindPtr
	fnShapedTextGetSpacing                   gdc.MethodBindPtr
	fnShapedTextAddString                    gdc.MethodBindPtr
	fnShapedTextAddObject                    gdc.MethodBindPtr
	fnShapedTextResizeObject                 gdc.MethodBindPtr
	fnShapedGetSpanCount                     gdc.MethodBindPtr
	fnShapedGetSpanMeta                      gdc.MethodBindPtr
	fnShapedSetSpanUpdateFont                gdc.MethodBindPtr
	fnShapedTextSubstr                       gdc.MethodBindPtr
	fnShapedTextGetParent                    gdc.MethodBindPtr
	fnShapedTextFitToWidth                   gdc.MethodBindPtr
	fnShapedTextTabAlign                     gdc.MethodBindPtr
	fnShapedTextShape                        gdc.MethodBindPtr
	fnShapedTextIsReady                      gdc.MethodBindPtr
	fnShapedTextHasVisibleChars              gdc.MethodBindPtr
	fnShapedTextGetGlyphs                    gdc.MethodBindPtr
	fnShapedTextSortLogical                  gdc.MethodBindPtr
	fnShapedTextGetGlyphCount                gdc.MethodBindPtr
	fnShapedTextGetRange                     gdc.MethodBindPtr
	fnShapedTextGetLineBreaksAdv             gdc.MethodBindPtr
	fnShapedTextGetLineBreaks                gdc.MethodBindPtr
	fnShapedTextGetWordBreaks                gdc.MethodBindPtr
	fnShapedTextGetTrimPos                   gdc.MethodBindPtr
	fnShapedTextGetEllipsisPos               gdc.MethodBindPtr
	fnShapedTextGetEllipsisGlyphs            gdc.MethodBindPtr
	fnShapedTextGetEllipsisGlyphCount        gdc.MethodBindPtr
	fnShapedTextOverrunTrimToWidth           gdc.MethodBindPtr
	fnShapedTextGetObjects                   gdc.MethodBindPtr
	fnShapedTextGetObjectRect                gdc.MethodBindPtr
	fnShapedTextGetSize                      gdc.MethodBindPtr
	fnShapedTextGetAscent                    gdc.MethodBindPtr
	fnShapedTextGetDescent                   gdc.MethodBindPtr
	fnShapedTextGetWidth                     gdc.MethodBindPtr
	fnShapedTextGetUnderlinePosition         gdc.MethodBindPtr
	fnShapedTextGetUnderlineThickness        gdc.MethodBindPtr
	fnShapedTextGetCarets                    gdc.MethodBindPtr
	fnShapedTextGetSelection                 gdc.MethodBindPtr
	fnShapedTextHitTestGrapheme              gdc.MethodBindPtr
	fnShapedTextHitTestPosition              gdc.MethodBindPtr
	fnShapedTextGetGraphemeBounds            gdc.MethodBindPtr
	fnShapedTextNextGraphemePos              gdc.MethodBindPtr
	fnShapedTextPrevGraphemePos              gdc.MethodBindPtr
	fnShapedTextGetCharacterBreaks           gdc.MethodBindPtr
	fnShapedTextNextCharacterPos             gdc.MethodBindPtr
	fnShapedTextPrevCharacterPos             gdc.MethodBindPtr
	fnShapedTextClosestCharacterPos          gdc.MethodBindPtr
	fnShapedTextDraw                         gdc.MethodBindPtr
	fnShapedTextDrawOutline                  gdc.MethodBindPtr
	fnShapedTextGetDominantDirectionInRange  gdc.MethodBindPtr
	fnFormatNumber                           gdc.MethodBindPtr
	fnParseNumber                            gdc.MethodBindPtr
	fnPercentSign                            gdc.MethodBindPtr
	fnStringGetWordBreaks                    gdc.MethodBindPtr
	fnStringGetCharacterBreaks               gdc.MethodBindPtr
	fnIsConfusable                           gdc.MethodBindPtr
	fnSpoofCheck                             gdc.MethodBindPtr
	fnStripDiacritics                        gdc.MethodBindPtr
	fnIsValidIdentifier                      gdc.MethodBindPtr
	fnStringToUpper                          gdc.MethodBindPtr
	fnStringToLower                          gdc.MethodBindPtr
	fnParseStructuredText                    gdc.MethodBindPtr
}

var ptrsForTextServer ptrsForTextServerList

func initTextServerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextServer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_feature")
		defer methodName.Destroy()
		ptrsForTextServer.fnHasFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3967367083))
	}
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForTextServer.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_features")
		defer methodName.Destroy()
		ptrsForTextServer.fnGetFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("load_support_data")
		defer methodName.Destroy()
		ptrsForTextServer.fnLoadSupportData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
	}
	{
		methodName := StringNameFromStr("get_support_data_filename")
		defer methodName.Destroy()
		ptrsForTextServer.fnGetSupportDataFilename = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_support_data_info")
		defer methodName.Destroy()
		ptrsForTextServer.fnGetSupportDataInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("save_support_data")
		defer methodName.Destroy()
		ptrsForTextServer.fnSaveSupportData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("is_locale_right_to_left")
		defer methodName.Destroy()
		ptrsForTextServer.fnIsLocaleRightToLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("name_to_tag")
		defer methodName.Destroy()
		ptrsForTextServer.fnNameToTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("tag_to_name")
		defer methodName.Destroy()
		ptrsForTextServer.fnTagToName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("has")
		defer methodName.Destroy()
		ptrsForTextServer.fnHas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
	}
	{
		methodName := StringNameFromStr("free_rid")
		defer methodName.Destroy()
		ptrsForTextServer.fnFreeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("create_font")
		defer methodName.Destroy()
		ptrsForTextServer.fnCreateFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("create_font_linked_variation")
		defer methodName.Destroy()
		ptrsForTextServer.fnCreateFontLinkedVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 41030802))
	}
	{
		methodName := StringNameFromStr("font_set_data")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1355495400))
	}
	{
		methodName := StringNameFromStr("font_set_face_index")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetFaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("font_get_face_index")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetFaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("font_get_face_count")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetFaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("font_set_style")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetStyle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 898466325))
	}
	{
		methodName := StringNameFromStr("font_get_style")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetStyle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3082502592))
	}
	{
		methodName := StringNameFromStr("font_set_name")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
	}
	{
		methodName := StringNameFromStr("font_get_name")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642473191))
	}
	{
		methodName := StringNameFromStr("font_get_ot_name_strings")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetOtNameStrings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1882737106))
	}
	{
		methodName := StringNameFromStr("font_set_style_name")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetStyleName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
	}
	{
		methodName := StringNameFromStr("font_get_style_name")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetStyleName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642473191))
	}
	{
		methodName := StringNameFromStr("font_set_weight")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("font_get_weight")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("font_set_stretch")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("font_get_stretch")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("font_set_antialiasing")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 958337235))
	}
	{
		methodName := StringNameFromStr("font_get_antialiasing")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3389420495))
	}
	{
		methodName := StringNameFromStr("font_set_generate_mipmaps")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("font_get_generate_mipmaps")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("font_set_multichannel_signed_distance_field")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetMultichannelSignedDistanceField = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("font_is_multichannel_signed_distance_field")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontIsMultichannelSignedDistanceField = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("font_set_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("font_get_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("font_set_msdf_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetMsdfSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("font_get_msdf_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetMsdfSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("font_set_fixed_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetFixedSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("font_get_fixed_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetFixedSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("font_set_fixed_size_scale_mode")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetFixedSizeScaleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1029390307))
	}
	{
		methodName := StringNameFromStr("font_get_fixed_size_scale_mode")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetFixedSizeScaleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4113120379))
	}
	{
		methodName := StringNameFromStr("font_set_allow_system_fallback")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetAllowSystemFallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("font_is_allow_system_fallback")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontIsAllowSystemFallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("font_set_force_autohinter")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetForceAutohinter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("font_is_force_autohinter")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontIsForceAutohinter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("font_set_hinting")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetHinting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1520010864))
	}
	{
		methodName := StringNameFromStr("font_get_hinting")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetHinting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3971592737))
	}
	{
		methodName := StringNameFromStr("font_set_subpixel_positioning")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetSubpixelPositioning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3830459669))
	}
	{
		methodName := StringNameFromStr("font_get_subpixel_positioning")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetSubpixelPositioning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2752233671))
	}
	{
		methodName := StringNameFromStr("font_set_embolden")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetEmbolden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("font_get_embolden")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetEmbolden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("font_set_spacing")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1307259930))
	}
	{
		methodName := StringNameFromStr("font_get_spacing")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1213653558))
	}
	{
		methodName := StringNameFromStr("font_set_transform")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
	}
	{
		methodName := StringNameFromStr("font_get_transform")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 213527486))
	}
	{
		methodName := StringNameFromStr("font_set_variation_coordinates")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetVariationCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1217542888))
	}
	{
		methodName := StringNameFromStr("font_get_variation_coordinates")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetVariationCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1882737106))
	}
	{
		methodName := StringNameFromStr("font_set_oversampling")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("font_get_oversampling")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("font_get_size_cache_list")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetSizeCacheList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("font_clear_size_cache")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontClearSizeCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("font_remove_size_cache")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRemoveSizeCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2450610377))
	}
	{
		methodName := StringNameFromStr("font_set_ascent")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1892459533))
	}
	{
		methodName := StringNameFromStr("font_get_ascent")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 755457166))
	}
	{
		methodName := StringNameFromStr("font_set_descent")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1892459533))
	}
	{
		methodName := StringNameFromStr("font_get_descent")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 755457166))
	}
	{
		methodName := StringNameFromStr("font_set_underline_position")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1892459533))
	}
	{
		methodName := StringNameFromStr("font_get_underline_position")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 755457166))
	}
	{
		methodName := StringNameFromStr("font_set_underline_thickness")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1892459533))
	}
	{
		methodName := StringNameFromStr("font_get_underline_thickness")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 755457166))
	}
	{
		methodName := StringNameFromStr("font_set_scale")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1892459533))
	}
	{
		methodName := StringNameFromStr("font_get_scale")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 755457166))
	}
	{
		methodName := StringNameFromStr("font_get_texture_count")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetTextureCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1311001310))
	}
	{
		methodName := StringNameFromStr("font_clear_textures")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontClearTextures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2450610377))
	}
	{
		methodName := StringNameFromStr("font_remove_texture")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRemoveTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3810512262))
	}
	{
		methodName := StringNameFromStr("font_set_texture_image")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetTextureImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2354485091))
	}
	{
		methodName := StringNameFromStr("font_get_texture_image")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetTextureImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2451761155))
	}
	{
		methodName := StringNameFromStr("font_set_texture_offsets")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetTextureOffsets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3005398047))
	}
	{
		methodName := StringNameFromStr("font_get_texture_offsets")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetTextureOffsets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3420028887))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_list")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 46086620))
	}
	{
		methodName := StringNameFromStr("font_clear_glyphs")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontClearGlyphs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2450610377))
	}
	{
		methodName := StringNameFromStr("font_remove_glyph")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRemoveGlyph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3810512262))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_advance")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2555689501))
	}
	{
		methodName := StringNameFromStr("font_set_glyph_advance")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetGlyphAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3219397315))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_offset")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 513728628))
	}
	{
		methodName := StringNameFromStr("font_set_glyph_offset")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetGlyphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1812632090))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 513728628))
	}
	{
		methodName := StringNameFromStr("font_set_glyph_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetGlyphSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1812632090))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_uv_rect")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphUvRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2274268786))
	}
	{
		methodName := StringNameFromStr("font_set_glyph_uv_rect")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetGlyphUvRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1973324081))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_texture_idx")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphTextureIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4292800474))
	}
	{
		methodName := StringNameFromStr("font_set_glyph_texture_idx")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetGlyphTextureIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4254580980))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_texture_rid")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphTextureRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1451696141))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_texture_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphTextureSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 513728628))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_contours")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphContours = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2903964473))
	}
	{
		methodName := StringNameFromStr("font_get_kerning_list")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetKerningList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1778388067))
	}
	{
		methodName := StringNameFromStr("font_clear_kerning_map")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontClearKerningMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("font_remove_kerning")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRemoveKerning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2141860016))
	}
	{
		methodName := StringNameFromStr("font_set_kerning")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetKerning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3630965883))
	}
	{
		methodName := StringNameFromStr("font_get_kerning")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetKerning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1019980169))
	}
	{
		methodName := StringNameFromStr("font_get_glyph_index")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlyphIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1765635060))
	}
	{
		methodName := StringNameFromStr("font_get_char_from_glyph_index")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetCharFromGlyphIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2156738276))
	}
	{
		methodName := StringNameFromStr("font_has_char")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontHasChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3120086654))
	}
	{
		methodName := StringNameFromStr("font_get_supported_chars")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetSupportedChars = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642473191))
	}
	{
		methodName := StringNameFromStr("font_render_range")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRenderRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4254580980))
	}
	{
		methodName := StringNameFromStr("font_render_glyph")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRenderGlyph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3810512262))
	}
	{
		methodName := StringNameFromStr("font_draw_glyph")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontDrawGlyph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1339057948))
	}
	{
		methodName := StringNameFromStr("font_draw_glyph_outline")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontDrawGlyphOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2626165733))
	}
	{
		methodName := StringNameFromStr("font_is_language_supported")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontIsLanguageSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3199320846))
	}
	{
		methodName := StringNameFromStr("font_set_language_support_override")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetLanguageSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2313957094))
	}
	{
		methodName := StringNameFromStr("font_get_language_support_override")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetLanguageSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2829184646))
	}
	{
		methodName := StringNameFromStr("font_remove_language_support_override")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRemoveLanguageSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
	}
	{
		methodName := StringNameFromStr("font_get_language_support_overrides")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetLanguageSupportOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2801473409))
	}
	{
		methodName := StringNameFromStr("font_is_script_supported")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontIsScriptSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3199320846))
	}
	{
		methodName := StringNameFromStr("font_set_script_support_override")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetScriptSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2313957094))
	}
	{
		methodName := StringNameFromStr("font_get_script_support_override")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetScriptSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2829184646))
	}
	{
		methodName := StringNameFromStr("font_remove_script_support_override")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontRemoveScriptSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
	}
	{
		methodName := StringNameFromStr("font_get_script_support_overrides")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetScriptSupportOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2801473409))
	}
	{
		methodName := StringNameFromStr("font_set_opentype_feature_overrides")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetOpentypeFeatureOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1217542888))
	}
	{
		methodName := StringNameFromStr("font_get_opentype_feature_overrides")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetOpentypeFeatureOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1882737106))
	}
	{
		methodName := StringNameFromStr("font_supported_feature_list")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSupportedFeatureList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1882737106))
	}
	{
		methodName := StringNameFromStr("font_supported_variation_list")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSupportedVariationList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1882737106))
	}
	{
		methodName := StringNameFromStr("font_get_global_oversampling")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontGetGlobalOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("font_set_global_oversampling")
		defer methodName.Destroy()
		ptrsForTextServer.fnFontSetGlobalOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_hex_code_box_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnGetHexCodeBoxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3016396712))
	}
	{
		methodName := StringNameFromStr("draw_hex_code_box")
		defer methodName.Destroy()
		ptrsForTextServer.fnDrawHexCodeBox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602046441))
	}
	{
		methodName := StringNameFromStr("create_shaped_text")
		defer methodName.Destroy()
		ptrsForTextServer.fnCreateShapedText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1231398698))
	}
	{
		methodName := StringNameFromStr("shaped_text_clear")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("shaped_text_set_direction")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1551430183))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_direction")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3065904362))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_inferred_direction")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetInferredDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3065904362))
	}
	{
		methodName := StringNameFromStr("shaped_text_set_bidi_override")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSetBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 684822712))
	}
	{
		methodName := StringNameFromStr("shaped_text_set_custom_punctuation")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSetCustomPunctuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_custom_punctuation")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetCustomPunctuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 642473191))
	}
	{
		methodName := StringNameFromStr("shaped_text_set_orientation")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3019609126))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_orientation")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3142708106))
	}
	{
		methodName := StringNameFromStr("shaped_text_set_preserve_invalid")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSetPreserveInvalid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_preserve_invalid")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetPreserveInvalid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("shaped_text_set_preserve_control")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSetPreserveControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_preserve_control")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetPreserveControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("shaped_text_set_spacing")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSetSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1307259930))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_spacing")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1213653558))
	}
	{
		methodName := StringNameFromStr("shaped_text_add_string")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextAddString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 623473029))
	}
	{
		methodName := StringNameFromStr("shaped_text_add_object")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextAddObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3664424789))
	}
	{
		methodName := StringNameFromStr("shaped_text_resize_object")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextResizeObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 790361552))
	}
	{
		methodName := StringNameFromStr("shaped_get_span_count")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedGetSpanCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("shaped_get_span_meta")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedGetSpanMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4069510997))
	}
	{
		methodName := StringNameFromStr("shaped_set_span_update_font")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedSetSpanUpdateFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2022725822))
	}
	{
		methodName := StringNameFromStr("shaped_text_substr")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSubstr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1937682086))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_parent")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("shaped_text_fit_to_width")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextFitToWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 530670926))
	}
	{
		methodName := StringNameFromStr("shaped_text_tab_align")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextTabAlign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1283669550))
	}
	{
		methodName := StringNameFromStr("shaped_text_shape")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
	}
	{
		methodName := StringNameFromStr("shaped_text_is_ready")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextIsReady = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("shaped_text_has_visible_chars")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextHasVisibleChars = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_glyphs")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetGlyphs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("shaped_text_sort_logical")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextSortLogical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2670461153))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_glyph_count")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetGlyphCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_range")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 733700038))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_line_breaks_adv")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetLineBreaksAdv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2376991424))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_line_breaks")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetLineBreaks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2651359741))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_word_breaks")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetWordBreaks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 185957063))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_trim_pos")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetTrimPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_ellipsis_pos")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetEllipsisPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_ellipsis_glyphs")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetEllipsisGlyphs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_ellipsis_glyph_count")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetEllipsisGlyphCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("shaped_text_overrun_trim_to_width")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextOverrunTrimToWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2723146520))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_objects")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_object_rect")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetObjectRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 447978354))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_size")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2440833711))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_ascent")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_descent")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_width")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_underline_position")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_underline_thickness")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_carets")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetCarets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1574219346))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_selection")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3714187733))
	}
	{
		methodName := StringNameFromStr("shaped_text_hit_test_grapheme")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextHitTestGrapheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3149310417))
	}
	{
		methodName := StringNameFromStr("shaped_text_hit_test_position")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextHitTestPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3149310417))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_grapheme_bounds")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetGraphemeBounds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2546185844))
	}
	{
		methodName := StringNameFromStr("shaped_text_next_grapheme_pos")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextNextGraphemePos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("shaped_text_prev_grapheme_pos")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextPrevGraphemePos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_character_breaks")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetCharacterBreaks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 788230395))
	}
	{
		methodName := StringNameFromStr("shaped_text_next_character_pos")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextNextCharacterPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("shaped_text_prev_character_pos")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextPrevCharacterPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("shaped_text_closest_character_pos")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextClosestCharacterPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1120910005))
	}
	{
		methodName := StringNameFromStr("shaped_text_draw")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 880389142))
	}
	{
		methodName := StringNameFromStr("shaped_text_draw_outline")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextDrawOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2559184194))
	}
	{
		methodName := StringNameFromStr("shaped_text_get_dominant_direction_in_range")
		defer methodName.Destroy()
		ptrsForTextServer.fnShapedTextGetDominantDirectionInRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3326907668))
	}
	{
		methodName := StringNameFromStr("format_number")
		defer methodName.Destroy()
		ptrsForTextServer.fnFormatNumber = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2664628024))
	}
	{
		methodName := StringNameFromStr("parse_number")
		defer methodName.Destroy()
		ptrsForTextServer.fnParseNumber = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2664628024))
	}
	{
		methodName := StringNameFromStr("percent_sign")
		defer methodName.Destroy()
		ptrsForTextServer.fnPercentSign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 993269549))
	}
	{
		methodName := StringNameFromStr("string_get_word_breaks")
		defer methodName.Destroy()
		ptrsForTextServer.fnStringGetWordBreaks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 581857818))
	}
	{
		methodName := StringNameFromStr("string_get_character_breaks")
		defer methodName.Destroy()
		ptrsForTextServer.fnStringGetCharacterBreaks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2333794773))
	}
	{
		methodName := StringNameFromStr("is_confusable")
		defer methodName.Destroy()
		ptrsForTextServer.fnIsConfusable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1433197768))
	}
	{
		methodName := StringNameFromStr("spoof_check")
		defer methodName.Destroy()
		ptrsForTextServer.fnSpoofCheck = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("strip_diacritics")
		defer methodName.Destroy()
		ptrsForTextServer.fnStripDiacritics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
	}
	{
		methodName := StringNameFromStr("is_valid_identifier")
		defer methodName.Destroy()
		ptrsForTextServer.fnIsValidIdentifier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("string_to_upper")
		defer methodName.Destroy()
		ptrsForTextServer.fnStringToUpper = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2664628024))
	}
	{
		methodName := StringNameFromStr("string_to_lower")
		defer methodName.Destroy()
		ptrsForTextServer.fnStringToLower = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2664628024))
	}
	{
		methodName := StringNameFromStr("parse_structured_text")
		defer methodName.Destroy()
		ptrsForTextServer.fnParseStructuredText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3310685015))
	}
}

type TextServer struct {
	RefCounted
}

func (me *TextServer) BaseClass() string {
	return "TextServer"
}

func NewTextServer() *TextServer {
	str := StringNameFromStr("TextServer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextServer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TextServerFontAntialiasing int

const (
	TextServerFontAntialiasingFontAntialiasingNone TextServerFontAntialiasing = 0
	TextServerFontAntialiasingFontAntialiasingGray TextServerFontAntialiasing = 1
	TextServerFontAntialiasingFontAntialiasingLcd  TextServerFontAntialiasing = 2
)

type TextServerFontLCDSubpixelLayout int

const (
	TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutNone TextServerFontLCDSubpixelLayout = 0
	TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutHrgb TextServerFontLCDSubpixelLayout = 1
	TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutHbgr TextServerFontLCDSubpixelLayout = 2
	TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutVrgb TextServerFontLCDSubpixelLayout = 3
	TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutVbgr TextServerFontLCDSubpixelLayout = 4
	TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutMax  TextServerFontLCDSubpixelLayout = 5
)

type TextServerDirection int

const (
	TextServerDirectionDirectionAuto      TextServerDirection = 0
	TextServerDirectionDirectionLtr       TextServerDirection = 1
	TextServerDirectionDirectionRtl       TextServerDirection = 2
	TextServerDirectionDirectionInherited TextServerDirection = 3
)

type TextServerOrientation int

const (
	TextServerOrientationOrientationHorizontal TextServerOrientation = 0
	TextServerOrientationOrientationVertical   TextServerOrientation = 1
)

type TextServerJustificationFlag int

const (
	TextServerJustificationFlagJustificationNone                         TextServerJustificationFlag = 0
	TextServerJustificationFlagJustificationKashida                      TextServerJustificationFlag = 1
	TextServerJustificationFlagJustificationWordBound                    TextServerJustificationFlag = 2
	TextServerJustificationFlagJustificationTrimEdgeSpaces               TextServerJustificationFlag = 4
	TextServerJustificationFlagJustificationAfterLastTab                 TextServerJustificationFlag = 8
	TextServerJustificationFlagJustificationConstrainEllipsis            TextServerJustificationFlag = 16
	TextServerJustificationFlagJustificationSkipLastLine                 TextServerJustificationFlag = 32
	TextServerJustificationFlagJustificationSkipLastLineWithVisibleChars TextServerJustificationFlag = 64
	TextServerJustificationFlagJustificationDoNotSkipSingleLine          TextServerJustificationFlag = 128
)

type TextServerAutowrapMode int

const (
	TextServerAutowrapModeAutowrapOff       TextServerAutowrapMode = 0
	TextServerAutowrapModeAutowrapArbitrary TextServerAutowrapMode = 1
	TextServerAutowrapModeAutowrapWord      TextServerAutowrapMode = 2
	TextServerAutowrapModeAutowrapWordSmart TextServerAutowrapMode = 3
)

type TextServerLineBreakFlag int

const (
	TextServerLineBreakFlagBreakNone           TextServerLineBreakFlag = 0
	TextServerLineBreakFlagBreakMandatory      TextServerLineBreakFlag = 1
	TextServerLineBreakFlagBreakWordBound      TextServerLineBreakFlag = 2
	TextServerLineBreakFlagBreakGraphemeBound  TextServerLineBreakFlag = 4
	TextServerLineBreakFlagBreakAdaptive       TextServerLineBreakFlag = 8
	TextServerLineBreakFlagBreakTrimEdgeSpaces TextServerLineBreakFlag = 16
)

type TextServerVisibleCharactersBehavior int

const (
	TextServerVisibleCharactersBehaviorVcCharsBeforeShaping TextServerVisibleCharactersBehavior = 0
	TextServerVisibleCharactersBehaviorVcCharsAfterShaping  TextServerVisibleCharactersBehavior = 1
	TextServerVisibleCharactersBehaviorVcGlyphsAuto         TextServerVisibleCharactersBehavior = 2
	TextServerVisibleCharactersBehaviorVcGlyphsLtr          TextServerVisibleCharactersBehavior = 3
	TextServerVisibleCharactersBehaviorVcGlyphsRtl          TextServerVisibleCharactersBehavior = 4
)

type TextServerOverrunBehavior int

const (
	TextServerOverrunBehaviorOverrunNoTrimming       TextServerOverrunBehavior = 0
	TextServerOverrunBehaviorOverrunTrimChar         TextServerOverrunBehavior = 1
	TextServerOverrunBehaviorOverrunTrimWord         TextServerOverrunBehavior = 2
	TextServerOverrunBehaviorOverrunTrimEllipsis     TextServerOverrunBehavior = 3
	TextServerOverrunBehaviorOverrunTrimWordEllipsis TextServerOverrunBehavior = 4
)

type TextServerTextOverrunFlag int

const (
	TextServerTextOverrunFlagOverrunNoTrim             TextServerTextOverrunFlag = 0
	TextServerTextOverrunFlagOverrunTrim               TextServerTextOverrunFlag = 1
	TextServerTextOverrunFlagOverrunTrimWordOnly       TextServerTextOverrunFlag = 2
	TextServerTextOverrunFlagOverrunAddEllipsis        TextServerTextOverrunFlag = 4
	TextServerTextOverrunFlagOverrunEnforceEllipsis    TextServerTextOverrunFlag = 8
	TextServerTextOverrunFlagOverrunJustificationAware TextServerTextOverrunFlag = 16
)

type TextServerGraphemeFlag int

const (
	TextServerGraphemeFlagGraphemeIsValid               TextServerGraphemeFlag = 1
	TextServerGraphemeFlagGraphemeIsRtl                 TextServerGraphemeFlag = 2
	TextServerGraphemeFlagGraphemeIsVirtual             TextServerGraphemeFlag = 4
	TextServerGraphemeFlagGraphemeIsSpace               TextServerGraphemeFlag = 8
	TextServerGraphemeFlagGraphemeIsBreakHard           TextServerGraphemeFlag = 16
	TextServerGraphemeFlagGraphemeIsBreakSoft           TextServerGraphemeFlag = 32
	TextServerGraphemeFlagGraphemeIsTab                 TextServerGraphemeFlag = 64
	TextServerGraphemeFlagGraphemeIsElongation          TextServerGraphemeFlag = 128
	TextServerGraphemeFlagGraphemeIsPunctuation         TextServerGraphemeFlag = 256
	TextServerGraphemeFlagGraphemeIsUnderscore          TextServerGraphemeFlag = 512
	TextServerGraphemeFlagGraphemeIsConnected           TextServerGraphemeFlag = 1024
	TextServerGraphemeFlagGraphemeIsSafeToInsertTatweel TextServerGraphemeFlag = 2048
	TextServerGraphemeFlagGraphemeIsEmbeddedObject      TextServerGraphemeFlag = 4096
)

type TextServerHinting int

const (
	TextServerHintingHintingNone   TextServerHinting = 0
	TextServerHintingHintingLight  TextServerHinting = 1
	TextServerHintingHintingNormal TextServerHinting = 2
)

type TextServerSubpixelPositioning int

const (
	TextServerSubpixelPositioningSubpixelPositioningDisabled          TextServerSubpixelPositioning = 0
	TextServerSubpixelPositioningSubpixelPositioningAuto              TextServerSubpixelPositioning = 1
	TextServerSubpixelPositioningSubpixelPositioningOneHalf           TextServerSubpixelPositioning = 2
	TextServerSubpixelPositioningSubpixelPositioningOneQuarter        TextServerSubpixelPositioning = 3
	TextServerSubpixelPositioningSubpixelPositioningOneHalfMaxSize    TextServerSubpixelPositioning = 20
	TextServerSubpixelPositioningSubpixelPositioningOneQuarterMaxSize TextServerSubpixelPositioning = 16
)

type TextServerFeature int

const (
	TextServerFeatureFeatureSimpleLayout                   TextServerFeature = 1
	TextServerFeatureFeatureBidiLayout                     TextServerFeature = 2
	TextServerFeatureFeatureVerticalLayout                 TextServerFeature = 4
	TextServerFeatureFeatureShaping                        TextServerFeature = 8
	TextServerFeatureFeatureKashidaJustification           TextServerFeature = 16
	TextServerFeatureFeatureBreakIterators                 TextServerFeature = 32
	TextServerFeatureFeatureFontBitmap                     TextServerFeature = 64
	TextServerFeatureFeatureFontDynamic                    TextServerFeature = 128
	TextServerFeatureFeatureFontMsdf                       TextServerFeature = 256
	TextServerFeatureFeatureFontSystem                     TextServerFeature = 512
	TextServerFeatureFeatureFontVariable                   TextServerFeature = 1024
	TextServerFeatureFeatureContextSensitiveCaseConversion TextServerFeature = 2048
	TextServerFeatureFeatureUseSupportData                 TextServerFeature = 4096
	TextServerFeatureFeatureUnicodeIdentifiers             TextServerFeature = 8192
	TextServerFeatureFeatureUnicodeSecurity                TextServerFeature = 16384
)

type TextServerContourPointTag int

const (
	TextServerContourPointTagContourCurveTagOn       TextServerContourPointTag = 1
	TextServerContourPointTagContourCurveTagOffConic TextServerContourPointTag = 0
	TextServerContourPointTagContourCurveTagOffCubic TextServerContourPointTag = 2
)

type TextServerSpacingType int

const (
	TextServerSpacingTypeSpacingGlyph  TextServerSpacingType = 0
	TextServerSpacingTypeSpacingSpace  TextServerSpacingType = 1
	TextServerSpacingTypeSpacingTop    TextServerSpacingType = 2
	TextServerSpacingTypeSpacingBottom TextServerSpacingType = 3
	TextServerSpacingTypeSpacingMax    TextServerSpacingType = 4
)

type TextServerFontStyle int

const (
	TextServerFontStyleFontBold       TextServerFontStyle = 1
	TextServerFontStyleFontItalic     TextServerFontStyle = 2
	TextServerFontStyleFontFixedWidth TextServerFontStyle = 4
)

type TextServerStructuredTextParser int

const (
	TextServerStructuredTextParserStructuredTextDefault  TextServerStructuredTextParser = 0
	TextServerStructuredTextParserStructuredTextUri      TextServerStructuredTextParser = 1
	TextServerStructuredTextParserStructuredTextFile     TextServerStructuredTextParser = 2
	TextServerStructuredTextParserStructuredTextEmail    TextServerStructuredTextParser = 3
	TextServerStructuredTextParserStructuredTextList     TextServerStructuredTextParser = 4
	TextServerStructuredTextParserStructuredTextGdscript TextServerStructuredTextParser = 5
	TextServerStructuredTextParserStructuredTextCustom   TextServerStructuredTextParser = 6
)

type TextServerFixedSizeScaleMode int

const (
	TextServerFixedSizeScaleModeFixedSizeScaleDisable     TextServerFixedSizeScaleMode = 0
	TextServerFixedSizeScaleModeFixedSizeScaleIntegerOnly TextServerFixedSizeScaleMode = 1
	TextServerFixedSizeScaleModeFixedSizeScaleEnabled     TextServerFixedSizeScaleMode = 2
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

func (me *TextServer) HasFeature(feature TextServerFeature) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&feature)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnHasFeature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) GetName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) GetFeatures() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnGetFeatures), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) LoadSupportData(filename String) bool {
	cargs := []gdc.ConstTypePtr{filename.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnLoadSupportData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) GetSupportDataFilename() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnGetSupportDataFilename), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) GetSupportDataInfo() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnGetSupportDataInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) SaveSupportData(filename String) bool {
	cargs := []gdc.ConstTypePtr{filename.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnSaveSupportData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) IsLocaleRightToLeft(locale String) bool {
	cargs := []gdc.ConstTypePtr{locale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnIsLocaleRightToLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) NameToTag(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnNameToTag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) TagToName(tag int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&tag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnTagToName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) Has(rid RID) bool {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnHas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FreeRid(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFreeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) CreateFont() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnCreateFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) CreateFontLinkedVariation(font_rid RID) RID {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnCreateFontLinkedVariation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetData(font_rid RID, data PackedByteArray) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontSetFaceIndex(font_rid RID, face_index int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&face_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetFaceIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetFaceIndex(font_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetFaceIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontGetFaceCount(font_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetFaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetStyle(font_rid RID, style TextServerFontStyle) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&style)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetStyle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetStyle(font_rid RID) TextServerFontStyle {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerFontStyle

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetStyle), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) FontSetName(font_rid RID, name String) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetName(font_rid RID) String {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontGetOtNameStrings(font_rid RID) Dictionary {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetOtNameStrings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetStyleName(font_rid RID, name String) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetStyleName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetStyleName(font_rid RID) String {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetStyleName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetWeight(font_rid RID, weight int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&weight)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetWeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetWeight(font_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetWeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetStretch(font_rid RID, weight int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&weight)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetStretch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetStretch(font_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetStretch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetAntialiasing(font_rid RID, antialiasing TextServerFontAntialiasing) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&antialiasing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetAntialiasing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetAntialiasing(font_rid RID) TextServerFontAntialiasing {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerFontAntialiasing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetAntialiasing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) FontSetGenerateMipmaps(font_rid RID, generate_mipmaps bool) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&generate_mipmaps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetGenerateMipmaps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetGenerateMipmaps(font_rid RID) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGenerateMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetMultichannelSignedDistanceField(font_rid RID, msdf bool) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&msdf)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetMultichannelSignedDistanceField), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontIsMultichannelSignedDistanceField(font_rid RID) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontIsMultichannelSignedDistanceField), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetMsdfPixelRange(font_rid RID, msdf_pixel_range int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&msdf_pixel_range)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetMsdfPixelRange(font_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetMsdfSize(font_rid RID, msdf_size int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&msdf_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetMsdfSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetMsdfSize(font_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetMsdfSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetFixedSize(font_rid RID, fixed_size int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&fixed_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetFixedSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetFixedSize(font_rid RID) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetFixedSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetFixedSizeScaleMode(font_rid RID, fixed_size_scale_mode TextServerFixedSizeScaleMode) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&fixed_size_scale_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetFixedSizeScaleMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetFixedSizeScaleMode(font_rid RID) TextServerFixedSizeScaleMode {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerFixedSizeScaleMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetFixedSizeScaleMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) FontSetAllowSystemFallback(font_rid RID, allow_system_fallback bool) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&allow_system_fallback)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetAllowSystemFallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontIsAllowSystemFallback(font_rid RID) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontIsAllowSystemFallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetForceAutohinter(font_rid RID, force_autohinter bool) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&force_autohinter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetForceAutohinter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontIsForceAutohinter(font_rid RID) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontIsForceAutohinter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetHinting(font_rid RID, hinting TextServerHinting) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&hinting)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetHinting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetHinting(font_rid RID) TextServerHinting {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerHinting

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetHinting), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) FontSetSubpixelPositioning(font_rid RID, subpixel_positioning TextServerSubpixelPositioning) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&subpixel_positioning)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetSubpixelPositioning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetSubpixelPositioning(font_rid RID) TextServerSubpixelPositioning {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerSubpixelPositioning

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetSubpixelPositioning), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) FontSetEmbolden(font_rid RID, strength float64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetEmbolden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetEmbolden(font_rid RID) float64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetEmbolden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetSpacing(font_rid RID, spacing TextServerSpacingType, value int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&spacing), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetSpacing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetSpacing(font_rid RID, spacing TextServerSpacingType) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&spacing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&spacing)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetSpacing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetTransform(font_rid RID, transform Transform2D) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetTransform(font_rid RID) Transform2D {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetVariationCoordinates(font_rid RID, variation_coordinates Dictionary) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), variation_coordinates.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetVariationCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetVariationCoordinates(font_rid RID) Dictionary {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetVariationCoordinates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetOversampling(font_rid RID, oversampling float64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&oversampling)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetOversampling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetOversampling(font_rid RID) float64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetOversampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontGetSizeCacheList(font_rid RID) []Vector2i {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetSizeCacheList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TextServer) FontClearSizeCache(font_rid RID) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontClearSizeCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontRemoveSizeCache(font_rid RID, size Vector2i) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRemoveSizeCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontSetAscent(font_rid RID, size int64, ascent float64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&ascent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetAscent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetAscent(font_rid RID, size int64) float64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetAscent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetDescent(font_rid RID, size int64, descent float64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&descent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetDescent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetDescent(font_rid RID, size int64) float64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetDescent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetUnderlinePosition(font_rid RID, size int64, underline_position float64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetUnderlinePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetUnderlinePosition(font_rid RID, size int64) float64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetUnderlinePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetUnderlineThickness(font_rid RID, size int64, underline_thickness float64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_thickness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetUnderlineThickness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetUnderlineThickness(font_rid RID, size int64) float64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetUnderlineThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetScale(font_rid RID, size int64, scale float64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetScale(font_rid RID, size int64) float64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontGetTextureCount(font_rid RID, size Vector2i) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetTextureCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontClearTextures(font_rid RID, size Vector2i) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontClearTextures), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontRemoveTexture(font_rid RID, size Vector2i, texture_index int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRemoveTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontSetTextureImage(font_rid RID, size Vector2i, texture_index int64, image Image) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index), image.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetTextureImage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetTextureImage(font_rid RID, size Vector2i, texture_index int64) Image {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()
	pinner.Pin(&texture_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetTextureImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetTextureOffsets(font_rid RID, size Vector2i, texture_index int64, offset PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index), offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetTextureOffsets), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetTextureOffsets(font_rid RID, size Vector2i, texture_index int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&texture_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetTextureOffsets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontGetGlyphList(font_rid RID, size Vector2i) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontClearGlyphs(font_rid RID, size Vector2i) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontClearGlyphs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontRemoveGlyph(font_rid RID, size Vector2i, glyph int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRemoveGlyph), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetGlyphAdvance(font_rid RID, size int64, glyph int64) Vector2 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&size)
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphAdvance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetGlyphAdvance(font_rid RID, size int64, glyph int64, advance Vector2) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph), advance.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetGlyphAdvance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetGlyphOffset(font_rid RID, size Vector2i, glyph int64) Vector2 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetGlyphOffset(font_rid RID, size Vector2i, glyph int64, offset Vector2) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetGlyphOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetGlyphSize(font_rid RID, size Vector2i, glyph int64) Vector2 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetGlyphSize(font_rid RID, size Vector2i, glyph int64, gl_size Vector2) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), gl_size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetGlyphSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetGlyphUvRect(font_rid RID, size Vector2i, glyph int64) Rect2 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphUvRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetGlyphUvRect(font_rid RID, size Vector2i, glyph int64, uv_rect Rect2) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), uv_rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetGlyphUvRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int64) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphTextureIdx), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int64, texture_idx int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(&texture_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetGlyphTextureIdx), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetGlyphTextureRid(font_rid RID, size Vector2i, glyph int64) RID {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphTextureRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontGetGlyphTextureSize(font_rid RID, size Vector2i, glyph int64) Vector2 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphTextureSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontGetGlyphContours(font RID, size int64, index int64) Dictionary {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&size)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphContours), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontGetKerningList(font_rid RID, size int64) []Vector2i {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetKerningList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TextServer) FontClearKerningMap(font_rid RID, size int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontClearKerningMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontRemoveKerning(font_rid RID, size int64, glyph_pair Vector2i) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), glyph_pair.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRemoveKerning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontSetKerning(font_rid RID, size int64, glyph_pair Vector2i, kerning Vector2) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), glyph_pair.AsCTypePtr(), kerning.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetKerning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetKerning(font_rid RID, size int64, glyph_pair Vector2i) Vector2 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), glyph_pair.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetKerning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontGetGlyphIndex(font_rid RID, size int64, char int64, variation_selector int64) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&variation_selector)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&size)
	pinner.Pin(&char)
	pinner.Pin(&variation_selector)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlyphIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontGetCharFromGlyphIndex(font_rid RID, size int64, glyph_index int64) int64 {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&size)
	pinner.Pin(&glyph_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetCharFromGlyphIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontHasChar(font_rid RID, char int64) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), gdc.ConstTypePtr(&char)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&char)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontHasChar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontGetSupportedChars(font_rid RID) String {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetSupportedChars), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontRenderRange(font_rid RID, size Vector2i, start int64, end int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRenderRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontRenderGlyph(font_rid RID, size Vector2i, index int64) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRenderGlyph), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontDrawGlyph(font_rid RID, canvas RID, size int64, pos Vector2, index int64, color Color) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), canvas.AsCTypePtr(), gdc.ConstTypePtr(&size), pos.AsCTypePtr(), gdc.ConstTypePtr(&index), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontDrawGlyph), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontDrawGlyphOutline(font_rid RID, canvas RID, size int64, outline_size int64, pos Vector2, index int64, color Color) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), canvas.AsCTypePtr(), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&outline_size), pos.AsCTypePtr(), gdc.ConstTypePtr(&index), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontDrawGlyphOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontIsLanguageSupported(font_rid RID, language String) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontIsLanguageSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetLanguageSupportOverride(font_rid RID, language String, supported bool) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), language.AsCTypePtr(), gdc.ConstTypePtr(&supported)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetLanguageSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetLanguageSupportOverride(font_rid RID, language String) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetLanguageSupportOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontRemoveLanguageSupportOverride(font_rid RID, language String) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRemoveLanguageSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetLanguageSupportOverrides(font_rid RID) PackedStringArray {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetLanguageSupportOverrides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontIsScriptSupported(font_rid RID, script String) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontIsScriptSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetScriptSupportOverride(font_rid RID, script String, supported bool) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), script.AsCTypePtr(), gdc.ConstTypePtr(&supported)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetScriptSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetScriptSupportOverride(font_rid RID, script String) bool {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetScriptSupportOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontRemoveScriptSupportOverride(font_rid RID, script String) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontRemoveScriptSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetScriptSupportOverrides(font_rid RID) PackedStringArray {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetScriptSupportOverrides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSetOpentypeFeatureOverrides(font_rid RID, overrides Dictionary) {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr(), overrides.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetOpentypeFeatureOverrides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) FontGetOpentypeFeatureOverrides(font_rid RID) Dictionary {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetOpentypeFeatureOverrides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSupportedFeatureList(font_rid RID) Dictionary {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSupportedFeatureList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontSupportedVariationList(font_rid RID) Dictionary {
	cargs := []gdc.ConstTypePtr{font_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSupportedVariationList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) FontGetGlobalOversampling() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontGetGlobalOversampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) FontSetGlobalOversampling(oversampling float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversampling)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFontSetGlobalOversampling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) GetHexCodeBoxSize(size int64, index int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&size)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnGetHexCodeBoxSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) DrawHexCodeBox(canvas RID, size int64, pos Vector2, index int64, color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), gdc.ConstTypePtr(&size), pos.AsCTypePtr(), gdc.ConstTypePtr(&index), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnDrawHexCodeBox), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) CreateShapedText(direction TextServerDirection, orientation TextServerOrientation) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&direction)
	pinner.Pin(&orientation)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnCreateShapedText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextClear(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextSetDirection(shaped RID, direction TextServerDirection) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSetDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetDirection(shaped RID) TextServerDirection {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) ShapedTextGetInferredDirection(shaped RID) TextServerDirection {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetInferredDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) ShapedTextSetBidiOverride(shaped RID, override Array) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), override.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSetBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextSetCustomPunctuation(shaped RID, punct String) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), punct.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSetCustomPunctuation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetCustomPunctuation(shaped RID) String {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetCustomPunctuation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextSetOrientation(shaped RID, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSetOrientation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetOrientation(shaped RID) TextServerOrientation {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOrientation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetOrientation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) ShapedTextSetPreserveInvalid(shaped RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSetPreserveInvalid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetPreserveInvalid(shaped RID) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetPreserveInvalid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextSetPreserveControl(shaped RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSetPreserveControl), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetPreserveControl(shaped RID) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetPreserveControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextSetSpacing(shaped RID, spacing TextServerSpacingType, value int64) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&spacing), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSetSpacing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetSpacing(shaped RID, spacing TextServerSpacingType) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&spacing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&spacing)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetSpacing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextAddString(shaped RID, text String, fonts []RID, size int64, opentype_features Dictionary, language String, meta Variant) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&fonts), gdc.ConstTypePtr(&size), opentype_features.AsCTypePtr(), language.AsCTypePtr(), meta.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&fonts)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextAddString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextAddObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, length int64, baseline float64) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&length), gdc.ConstTypePtr(&baseline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&inline_align)
	pinner.Pin(&length)
	pinner.Pin(&baseline)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextAddObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextResizeObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, baseline float64) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&baseline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&inline_align)
	pinner.Pin(&baseline)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextResizeObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedGetSpanCount(shaped RID) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedGetSpanCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedGetSpanMeta(shaped RID, index int64) Variant {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedGetSpanMeta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedSetSpanUpdateFont(shaped RID, index int64, fonts []RID, size int64, opentype_features Dictionary) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&fonts), gdc.ConstTypePtr(&size), opentype_features.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedSetSpanUpdateFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextSubstr(shaped RID, start int64, length int64) RID {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&start)
	pinner.Pin(&length)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSubstr), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetParent(shaped RID) RID {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextFitToWidth(shaped RID, width float64, justification_flags TextServerJustificationFlag) float64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&justification_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&width)
	pinner.Pin(&justification_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextFitToWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextTabAlign(shaped RID, tab_stops PackedFloat32Array) float64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), tab_stops.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextTabAlign), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextShape(shaped RID) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextIsReady(shaped RID) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextIsReady), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextHasVisibleChars(shaped RID) bool {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextHasVisibleChars), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetGlyphs(shaped RID) []Dictionary {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetGlyphs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TextServer) ShapedTextSortLogical(shaped RID) []Dictionary {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextSortLogical), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TextServer) ShapedTextGetGlyphCount(shaped RID) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetGlyphCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetRange(shaped RID) Vector2i {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetLineBreaksAdv(shaped RID, width PackedFloat32Array, start int64, once bool, break_flags TextServerLineBreakFlag) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), width.AsCTypePtr(), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&once), gdc.ConstTypePtr(&break_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&start)
	pinner.Pin(&once)
	pinner.Pin(&break_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetLineBreaksAdv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetLineBreaks(shaped RID, width float64, start int64, break_flags TextServerLineBreakFlag) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&break_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&width)
	pinner.Pin(&start)
	pinner.Pin(&break_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetLineBreaks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetWordBreaks(shaped RID, grapheme_flags TextServerGraphemeFlag) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&grapheme_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&grapheme_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetWordBreaks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetTrimPos(shaped RID) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetTrimPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetEllipsisPos(shaped RID) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetEllipsisPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetEllipsisGlyphs(shaped RID) []Dictionary {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetEllipsisGlyphs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TextServer) ShapedTextGetEllipsisGlyphCount(shaped RID) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetEllipsisGlyphCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextOverrunTrimToWidth(shaped RID, width float64, overrun_trim_flags TextServerTextOverrunFlag) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&overrun_trim_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextOverrunTrimToWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetObjects(shaped RID) Array {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetObjects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetObjectRect(shaped RID, key Variant) Rect2 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), key.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetObjectRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetSize(shaped RID) Vector2 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetAscent(shaped RID) float64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetAscent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetDescent(shaped RID) float64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetDescent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetWidth(shaped RID) float64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetUnderlinePosition(shaped RID) float64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetUnderlinePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetUnderlineThickness(shaped RID) float64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetUnderlineThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetCarets(shaped RID, position int64) Dictionary {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&position)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetCarets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextGetSelection(shaped RID, start int64, end int64) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()
	pinner.Pin(&start)
	pinner.Pin(&end)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextHitTestGrapheme(shaped RID, coords float64) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&coords)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&coords)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextHitTestGrapheme), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextHitTestPosition(shaped RID, coords float64) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&coords)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&coords)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextHitTestPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetGraphemeBounds(shaped RID, pos int64) Vector2 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&pos)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetGraphemeBounds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextNextGraphemePos(shaped RID, pos int64) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&pos)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextNextGraphemePos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextPrevGraphemePos(shaped RID, pos int64) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&pos)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextPrevGraphemePos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextGetCharacterBreaks(shaped RID) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetCharacterBreaks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ShapedTextNextCharacterPos(shaped RID, pos int64) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&pos)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextNextCharacterPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextPrevCharacterPos(shaped RID, pos int64) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&pos)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextPrevCharacterPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextClosestCharacterPos(shaped RID, pos int64) int64 {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&pos)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextClosestCharacterPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) ShapedTextDraw(shaped RID, canvas RID, pos Vector2, clip_l float64, clip_r float64, color Color) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&clip_l), gdc.ConstTypePtr(&clip_r), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextDrawOutline(shaped RID, canvas RID, pos Vector2, clip_l float64, clip_r float64, outline_size int64, color Color) {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&clip_l), gdc.ConstTypePtr(&clip_r), gdc.ConstTypePtr(&outline_size), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextDrawOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextServer) ShapedTextGetDominantDirectionInRange(shaped RID, start int64, end int64) TextServerDirection {
	cargs := []gdc.ConstTypePtr{shaped.AsCTypePtr(), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerDirection
	pinner.Pin(&start)
	pinner.Pin(&end)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnShapedTextGetDominantDirectionInRange), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextServer) FormatNumber(number String, language String) String {
	cargs := []gdc.ConstTypePtr{number.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnFormatNumber), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ParseNumber(number String, language String) String {
	cargs := []gdc.ConstTypePtr{number.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnParseNumber), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) PercentSign(language String) String {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnPercentSign), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) StringGetWordBreaks(string_ String, language String, chars_per_line int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), language.AsCTypePtr(), gdc.ConstTypePtr(&chars_per_line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&chars_per_line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnStringGetWordBreaks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) StringGetCharacterBreaks(string_ String, language String) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnStringGetCharacterBreaks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) IsConfusable(string_ String, dict PackedStringArray) int64 {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), dict.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnIsConfusable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) SpoofCheck(string_ String) bool {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnSpoofCheck), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) StripDiacritics(string_ String) String {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnStripDiacritics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) IsValidIdentifier(string_ String) bool {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnIsValidIdentifier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextServer) StringToUpper(string_ String, language String) String {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnStringToUpper), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) StringToLower(string_ String, language String) String {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnStringToLower), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextServer) ParseStructuredText(parser_type TextServerStructuredTextParser, args Array, text String) []Vector3i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser_type), args.AsCTypePtr(), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&parser_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServer.fnParseStructuredText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector3i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Signals
