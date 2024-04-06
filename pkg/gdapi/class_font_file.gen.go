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

type ptrsForFontFileList struct {
	fnLoadBitmapFont                     gdc.MethodBindPtr
	fnLoadDynamicFont                    gdc.MethodBindPtr
	fnSetData                            gdc.MethodBindPtr
	fnGetData                            gdc.MethodBindPtr
	fnSetFontName                        gdc.MethodBindPtr
	fnSetFontStyleName                   gdc.MethodBindPtr
	fnSetFontStyle                       gdc.MethodBindPtr
	fnSetFontWeight                      gdc.MethodBindPtr
	fnSetFontStretch                     gdc.MethodBindPtr
	fnSetAntialiasing                    gdc.MethodBindPtr
	fnGetAntialiasing                    gdc.MethodBindPtr
	fnSetGenerateMipmaps                 gdc.MethodBindPtr
	fnGetGenerateMipmaps                 gdc.MethodBindPtr
	fnSetMultichannelSignedDistanceField gdc.MethodBindPtr
	fnIsMultichannelSignedDistanceField  gdc.MethodBindPtr
	fnSetMsdfPixelRange                  gdc.MethodBindPtr
	fnGetMsdfPixelRange                  gdc.MethodBindPtr
	fnSetMsdfSize                        gdc.MethodBindPtr
	fnGetMsdfSize                        gdc.MethodBindPtr
	fnSetFixedSize                       gdc.MethodBindPtr
	fnGetFixedSize                       gdc.MethodBindPtr
	fnSetFixedSizeScaleMode              gdc.MethodBindPtr
	fnGetFixedSizeScaleMode              gdc.MethodBindPtr
	fnSetAllowSystemFallback             gdc.MethodBindPtr
	fnIsAllowSystemFallback              gdc.MethodBindPtr
	fnSetForceAutohinter                 gdc.MethodBindPtr
	fnIsForceAutohinter                  gdc.MethodBindPtr
	fnSetHinting                         gdc.MethodBindPtr
	fnGetHinting                         gdc.MethodBindPtr
	fnSetSubpixelPositioning             gdc.MethodBindPtr
	fnGetSubpixelPositioning             gdc.MethodBindPtr
	fnSetOversampling                    gdc.MethodBindPtr
	fnGetOversampling                    gdc.MethodBindPtr
	fnGetCacheCount                      gdc.MethodBindPtr
	fnClearCache                         gdc.MethodBindPtr
	fnRemoveCache                        gdc.MethodBindPtr
	fnGetSizeCacheList                   gdc.MethodBindPtr
	fnClearSizeCache                     gdc.MethodBindPtr
	fnRemoveSizeCache                    gdc.MethodBindPtr
	fnSetVariationCoordinates            gdc.MethodBindPtr
	fnGetVariationCoordinates            gdc.MethodBindPtr
	fnSetEmbolden                        gdc.MethodBindPtr
	fnGetEmbolden                        gdc.MethodBindPtr
	fnSetTransform                       gdc.MethodBindPtr
	fnGetTransform                       gdc.MethodBindPtr
	fnSetExtraSpacing                    gdc.MethodBindPtr
	fnGetExtraSpacing                    gdc.MethodBindPtr
	fnSetFaceIndex                       gdc.MethodBindPtr
	fnGetFaceIndex                       gdc.MethodBindPtr
	fnSetCacheAscent                     gdc.MethodBindPtr
	fnGetCacheAscent                     gdc.MethodBindPtr
	fnSetCacheDescent                    gdc.MethodBindPtr
	fnGetCacheDescent                    gdc.MethodBindPtr
	fnSetCacheUnderlinePosition          gdc.MethodBindPtr
	fnGetCacheUnderlinePosition          gdc.MethodBindPtr
	fnSetCacheUnderlineThickness         gdc.MethodBindPtr
	fnGetCacheUnderlineThickness         gdc.MethodBindPtr
	fnSetCacheScale                      gdc.MethodBindPtr
	fnGetCacheScale                      gdc.MethodBindPtr
	fnGetTextureCount                    gdc.MethodBindPtr
	fnClearTextures                      gdc.MethodBindPtr
	fnRemoveTexture                      gdc.MethodBindPtr
	fnSetTextureImage                    gdc.MethodBindPtr
	fnGetTextureImage                    gdc.MethodBindPtr
	fnSetTextureOffsets                  gdc.MethodBindPtr
	fnGetTextureOffsets                  gdc.MethodBindPtr
	fnGetGlyphList                       gdc.MethodBindPtr
	fnClearGlyphs                        gdc.MethodBindPtr
	fnRemoveGlyph                        gdc.MethodBindPtr
	fnSetGlyphAdvance                    gdc.MethodBindPtr
	fnGetGlyphAdvance                    gdc.MethodBindPtr
	fnSetGlyphOffset                     gdc.MethodBindPtr
	fnGetGlyphOffset                     gdc.MethodBindPtr
	fnSetGlyphSize                       gdc.MethodBindPtr
	fnGetGlyphSize                       gdc.MethodBindPtr
	fnSetGlyphUvRect                     gdc.MethodBindPtr
	fnGetGlyphUvRect                     gdc.MethodBindPtr
	fnSetGlyphTextureIdx                 gdc.MethodBindPtr
	fnGetGlyphTextureIdx                 gdc.MethodBindPtr
	fnGetKerningList                     gdc.MethodBindPtr
	fnClearKerningMap                    gdc.MethodBindPtr
	fnRemoveKerning                      gdc.MethodBindPtr
	fnSetKerning                         gdc.MethodBindPtr
	fnGetKerning                         gdc.MethodBindPtr
	fnRenderRange                        gdc.MethodBindPtr
	fnRenderGlyph                        gdc.MethodBindPtr
	fnSetLanguageSupportOverride         gdc.MethodBindPtr
	fnGetLanguageSupportOverride         gdc.MethodBindPtr
	fnRemoveLanguageSupportOverride      gdc.MethodBindPtr
	fnGetLanguageSupportOverrides        gdc.MethodBindPtr
	fnSetScriptSupportOverride           gdc.MethodBindPtr
	fnGetScriptSupportOverride           gdc.MethodBindPtr
	fnRemoveScriptSupportOverride        gdc.MethodBindPtr
	fnGetScriptSupportOverrides          gdc.MethodBindPtr
	fnSetOpentypeFeatureOverrides        gdc.MethodBindPtr
	fnGetOpentypeFeatureOverrides        gdc.MethodBindPtr
	fnGetGlyphIndex                      gdc.MethodBindPtr
	fnGetCharFromGlyphIndex              gdc.MethodBindPtr
}

var ptrsForFontFile ptrsForFontFileList

func initFontFilePtrs(iface gdc.Interface) {

	className := StringNameFromStr("FontFile")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("load_bitmap_font")
		defer methodName.Destroy()
		ptrsForFontFile.fnLoadBitmapFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("load_dynamic_font")
		defer methodName.Destroy()
		ptrsForFontFile.fnLoadDynamicFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("set_data")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2971499966))
	}
	{
		methodName := StringNameFromStr("get_data")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
	}
	{
		methodName := StringNameFromStr("set_font_name")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFontName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_font_style_name")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFontStyleName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_font_style")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFontStyle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 918070724))
	}
	{
		methodName := StringNameFromStr("set_font_weight")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFontWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_font_stretch")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFontStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_antialiasing")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1669900))
	}
	{
		methodName := StringNameFromStr("get_antialiasing")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4262718649))
	}
	{
		methodName := StringNameFromStr("set_generate_mipmaps")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_generate_mipmaps")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGenerateMipmaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_multichannel_signed_distance_field")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetMultichannelSignedDistanceField = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_multichannel_signed_distance_field")
		defer methodName.Destroy()
		ptrsForFontFile.fnIsMultichannelSignedDistanceField = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_msdf_pixel_range")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetMsdfPixelRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_msdf_size")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetMsdfSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_msdf_size")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetMsdfSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_fixed_size")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFixedSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_fixed_size")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetFixedSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_fixed_size_scale_mode")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFixedSizeScaleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1660989956))
	}
	{
		methodName := StringNameFromStr("get_fixed_size_scale_mode")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetFixedSizeScaleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 753873478))
	}
	{
		methodName := StringNameFromStr("set_allow_system_fallback")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetAllowSystemFallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_allow_system_fallback")
		defer methodName.Destroy()
		ptrsForFontFile.fnIsAllowSystemFallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_force_autohinter")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetForceAutohinter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_force_autohinter")
		defer methodName.Destroy()
		ptrsForFontFile.fnIsForceAutohinter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hinting")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetHinting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1827459492))
	}
	{
		methodName := StringNameFromStr("get_hinting")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetHinting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3683214614))
	}
	{
		methodName := StringNameFromStr("set_subpixel_positioning")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetSubpixelPositioning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4225742182))
	}
	{
		methodName := StringNameFromStr("get_subpixel_positioning")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetSubpixelPositioning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1069238588))
	}
	{
		methodName := StringNameFromStr("set_oversampling")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_oversampling")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_cache_count")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetCacheCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("clear_cache")
		defer methodName.Destroy()
		ptrsForFontFile.fnClearCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("remove_cache")
		defer methodName.Destroy()
		ptrsForFontFile.fnRemoveCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_size_cache_list")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetSizeCacheList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("clear_size_cache")
		defer methodName.Destroy()
		ptrsForFontFile.fnClearSizeCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("remove_size_cache")
		defer methodName.Destroy()
		ptrsForFontFile.fnRemoveSizeCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311374912))
	}
	{
		methodName := StringNameFromStr("set_variation_coordinates")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetVariationCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 64545446))
	}
	{
		methodName := StringNameFromStr("get_variation_coordinates")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetVariationCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3485342025))
	}
	{
		methodName := StringNameFromStr("set_embolden")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetEmbolden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_embolden")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetEmbolden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_transform")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 30160968))
	}
	{
		methodName := StringNameFromStr("get_transform")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3836996910))
	}
	{
		methodName := StringNameFromStr("set_extra_spacing")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetExtraSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 62942285))
	}
	{
		methodName := StringNameFromStr("get_extra_spacing")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetExtraSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1924257185))
	}
	{
		methodName := StringNameFromStr("set_face_index")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetFaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_face_index")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetFaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_cache_ascent")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetCacheAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
	}
	{
		methodName := StringNameFromStr("get_cache_ascent")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetCacheAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("set_cache_descent")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetCacheDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
	}
	{
		methodName := StringNameFromStr("get_cache_descent")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetCacheDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("set_cache_underline_position")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetCacheUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
	}
	{
		methodName := StringNameFromStr("get_cache_underline_position")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetCacheUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("set_cache_underline_thickness")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetCacheUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
	}
	{
		methodName := StringNameFromStr("get_cache_underline_thickness")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetCacheUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("set_cache_scale")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetCacheScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
	}
	{
		methodName := StringNameFromStr("get_cache_scale")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetCacheScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("get_texture_count")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetTextureCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1987661582))
	}
	{
		methodName := StringNameFromStr("clear_textures")
		defer methodName.Destroy()
		ptrsForFontFile.fnClearTextures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311374912))
	}
	{
		methodName := StringNameFromStr("remove_texture")
		defer methodName.Destroy()
		ptrsForFontFile.fnRemoveTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2328951467))
	}
	{
		methodName := StringNameFromStr("set_texture_image")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetTextureImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4157974066))
	}
	{
		methodName := StringNameFromStr("get_texture_image")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetTextureImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3878418953))
	}
	{
		methodName := StringNameFromStr("set_texture_offsets")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetTextureOffsets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2849993437))
	}
	{
		methodName := StringNameFromStr("get_texture_offsets")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetTextureOffsets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3703444828))
	}
	{
		methodName := StringNameFromStr("get_glyph_list")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGlyphList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 681709689))
	}
	{
		methodName := StringNameFromStr("clear_glyphs")
		defer methodName.Destroy()
		ptrsForFontFile.fnClearGlyphs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311374912))
	}
	{
		methodName := StringNameFromStr("remove_glyph")
		defer methodName.Destroy()
		ptrsForFontFile.fnRemoveGlyph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2328951467))
	}
	{
		methodName := StringNameFromStr("set_glyph_advance")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetGlyphAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 947991729))
	}
	{
		methodName := StringNameFromStr("get_glyph_advance")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGlyphAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1601573536))
	}
	{
		methodName := StringNameFromStr("set_glyph_offset")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetGlyphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 921719850))
	}
	{
		methodName := StringNameFromStr("get_glyph_offset")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGlyphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3205412300))
	}
	{
		methodName := StringNameFromStr("set_glyph_size")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetGlyphSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 921719850))
	}
	{
		methodName := StringNameFromStr("get_glyph_size")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGlyphSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3205412300))
	}
	{
		methodName := StringNameFromStr("set_glyph_uv_rect")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetGlyphUvRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3821620992))
	}
	{
		methodName := StringNameFromStr("get_glyph_uv_rect")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGlyphUvRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927917900))
	}
	{
		methodName := StringNameFromStr("set_glyph_texture_idx")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetGlyphTextureIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 355564111))
	}
	{
		methodName := StringNameFromStr("get_glyph_texture_idx")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGlyphTextureIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1629411054))
	}
	{
		methodName := StringNameFromStr("get_kerning_list")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetKerningList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2345056839))
	}
	{
		methodName := StringNameFromStr("clear_kerning_map")
		defer methodName.Destroy()
		ptrsForFontFile.fnClearKerningMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_kerning")
		defer methodName.Destroy()
		ptrsForFontFile.fnRemoveKerning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3930204747))
	}
	{
		methodName := StringNameFromStr("set_kerning")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetKerning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3182200918))
	}
	{
		methodName := StringNameFromStr("get_kerning")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetKerning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611912865))
	}
	{
		methodName := StringNameFromStr("render_range")
		defer methodName.Destroy()
		ptrsForFontFile.fnRenderRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 355564111))
	}
	{
		methodName := StringNameFromStr("render_glyph")
		defer methodName.Destroy()
		ptrsForFontFile.fnRenderGlyph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2328951467))
	}
	{
		methodName := StringNameFromStr("set_language_support_override")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetLanguageSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
	}
	{
		methodName := StringNameFromStr("get_language_support_override")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetLanguageSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("remove_language_support_override")
		defer methodName.Destroy()
		ptrsForFontFile.fnRemoveLanguageSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language_support_overrides")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetLanguageSupportOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_script_support_override")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetScriptSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
	}
	{
		methodName := StringNameFromStr("get_script_support_override")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetScriptSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("remove_script_support_override")
		defer methodName.Destroy()
		ptrsForFontFile.fnRemoveScriptSupportOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_script_support_overrides")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetScriptSupportOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_opentype_feature_overrides")
		defer methodName.Destroy()
		ptrsForFontFile.fnSetOpentypeFeatureOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
	}
	{
		methodName := StringNameFromStr("get_opentype_feature_overrides")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetOpentypeFeatureOverrides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_glyph_index")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetGlyphIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 864943070))
	}
	{
		methodName := StringNameFromStr("get_char_from_glyph_index")
		defer methodName.Destroy()
		ptrsForFontFile.fnGetCharFromGlyphIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
}

type FontFile struct {
	Font
}

func (me *FontFile) BaseClass() string {
	return "FontFile"
}

func NewFontFile() *FontFile {
	str := StringNameFromStr("FontFile") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &FontFile{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *FontFile) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *FontFile) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *FontFile) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *FontFile) LoadBitmapFont(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnLoadBitmapFont), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FontFile) LoadDynamicFont(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnLoadDynamicFont), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FontFile) SetData(data PackedByteArray) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetData() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetFontName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFontName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetFontStyleName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFontStyleName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetFontStyle(style TextServerFontStyle) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&style)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFontStyle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetFontWeight(weight int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&weight)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFontWeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetFontStretch(stretch int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFontStretch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetAntialiasing(antialiasing TextServerFontAntialiasing) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiasing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetAntialiasing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetAntialiasing() TextServerFontAntialiasing {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerFontAntialiasing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetAntialiasing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FontFile) SetGenerateMipmaps(generate_mipmaps bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&generate_mipmaps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetGenerateMipmaps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetGenerateMipmaps() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGenerateMipmaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetMultichannelSignedDistanceField(msdf bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetMultichannelSignedDistanceField), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) IsMultichannelSignedDistanceField() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnIsMultichannelSignedDistanceField), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetMsdfPixelRange(msdf_pixel_range int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_pixel_range)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetMsdfPixelRange() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetMsdfPixelRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetMsdfSize(msdf_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msdf_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetMsdfSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetMsdfSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetMsdfSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetFixedSize(fixed_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fixed_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFixedSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetFixedSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetFixedSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetFixedSizeScaleMode(fixed_size_scale_mode TextServerFixedSizeScaleMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fixed_size_scale_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFixedSizeScaleMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetFixedSizeScaleMode() TextServerFixedSizeScaleMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerFixedSizeScaleMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetFixedSizeScaleMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FontFile) SetAllowSystemFallback(allow_system_fallback bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_system_fallback)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetAllowSystemFallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) IsAllowSystemFallback() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnIsAllowSystemFallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetForceAutohinter(force_autohinter bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_autohinter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetForceAutohinter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) IsForceAutohinter() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnIsForceAutohinter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetHinting(hinting TextServerHinting) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hinting)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetHinting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetHinting() TextServerHinting {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerHinting

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetHinting), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FontFile) SetSubpixelPositioning(subpixel_positioning TextServerSubpixelPositioning) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subpixel_positioning)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetSubpixelPositioning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetSubpixelPositioning() TextServerSubpixelPositioning {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerSubpixelPositioning

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetSubpixelPositioning), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FontFile) SetOversampling(oversampling float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversampling)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetOversampling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetOversampling() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetOversampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) GetCacheCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetCacheCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) ClearCache() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnClearCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) RemoveCache(cache_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRemoveCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetSizeCacheList(cache_index int64) []Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&cache_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetSizeCacheList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *FontFile) ClearSizeCache(cache_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnClearSizeCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) RemoveSizeCache(cache_index int64, size Vector2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRemoveSizeCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetVariationCoordinates(cache_index int64, variation_coordinates Dictionary) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), variation_coordinates.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetVariationCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetVariationCoordinates(cache_index int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&cache_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetVariationCoordinates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetEmbolden(cache_index int64, strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetEmbolden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetEmbolden(cache_index int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&cache_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetEmbolden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetTransform(cache_index int64, transform Transform2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetTransform(cache_index int64) Transform2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()
	pinner.Pin(&cache_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetExtraSpacing(cache_index int64, spacing TextServerSpacingType, value int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&spacing), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetExtraSpacing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetExtraSpacing(cache_index int64, spacing TextServerSpacingType) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&spacing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&cache_index)
	pinner.Pin(&spacing)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetExtraSpacing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetFaceIndex(cache_index int64, face_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&face_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetFaceIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetFaceIndex(cache_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&cache_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetFaceIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetCacheAscent(cache_index int64, size int64, ascent float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&ascent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetCacheAscent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetCacheAscent(cache_index int64, size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetCacheAscent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetCacheDescent(cache_index int64, size int64, descent float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&descent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetCacheDescent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetCacheDescent(cache_index int64, size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetCacheDescent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetCacheUnderlinePosition(cache_index int64, size int64, underline_position float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetCacheUnderlinePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetCacheUnderlinePosition(cache_index int64, size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetCacheUnderlinePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetCacheUnderlineThickness(cache_index int64, size int64, underline_thickness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_thickness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetCacheUnderlineThickness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetCacheUnderlineThickness(cache_index int64, size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetCacheUnderlineThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) SetCacheScale(cache_index int64, size int64, scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetCacheScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetCacheScale(cache_index int64, size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetCacheScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) GetTextureCount(cache_index int64, size Vector2i) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&cache_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetTextureCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) ClearTextures(cache_index int64, size Vector2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnClearTextures), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) RemoveTexture(cache_index int64, size Vector2i, texture_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRemoveTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetTextureImage(cache_index int64, size Vector2i, texture_index int64, image Image) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index), image.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetTextureImage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetTextureImage(cache_index int64, size Vector2i, texture_index int64) Image {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()
	pinner.Pin(&cache_index)
	pinner.Pin(&texture_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetTextureImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetTextureOffsets(cache_index int64, size Vector2i, texture_index int64, offset PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index), offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetTextureOffsets), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetTextureOffsets(cache_index int64, size Vector2i, texture_index int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&texture_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&cache_index)
	pinner.Pin(&texture_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetTextureOffsets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) GetGlyphList(cache_index int64, size Vector2i) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&cache_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGlyphList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) ClearGlyphs(cache_index int64, size Vector2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnClearGlyphs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) RemoveGlyph(cache_index int64, size Vector2i, glyph int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRemoveGlyph), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetGlyphAdvance(cache_index int64, size int64, glyph int64, advance Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph), advance.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetGlyphAdvance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetGlyphAdvance(cache_index int64, size int64, glyph int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGlyphAdvance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetGlyphOffset(cache_index int64, size Vector2i, glyph int64, offset Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetGlyphOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetGlyphOffset(cache_index int64, size Vector2i, glyph int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&cache_index)
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGlyphOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetGlyphSize(cache_index int64, size Vector2i, glyph int64, gl_size Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), gl_size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetGlyphSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetGlyphSize(cache_index int64, size Vector2i, glyph int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&cache_index)
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGlyphSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetGlyphUvRect(cache_index int64, size Vector2i, glyph int64, uv_rect Rect2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), uv_rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetGlyphUvRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetGlyphUvRect(cache_index int64, size Vector2i, glyph int64) Rect2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&cache_index)
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGlyphUvRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetGlyphTextureIdx(cache_index int64, size Vector2i, glyph int64, texture_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(&texture_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetGlyphTextureIdx), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetGlyphTextureIdx(cache_index int64, size Vector2i, glyph int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&glyph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&cache_index)
	pinner.Pin(&glyph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGlyphTextureIdx), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) GetKerningList(cache_index int64, size int64) []Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetKerningList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2i](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *FontFile) ClearKerningMap(cache_index int64, size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnClearKerningMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) RemoveKerning(cache_index int64, size int64, glyph_pair Vector2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), glyph_pair.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRemoveKerning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetKerning(cache_index int64, size int64, glyph_pair Vector2i, kerning Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), glyph_pair.AsCTypePtr(), kerning.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetKerning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetKerning(cache_index int64, size int64, glyph_pair Vector2i) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), gdc.ConstTypePtr(&size), glyph_pair.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&cache_index)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetKerning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) RenderRange(cache_index int64, size Vector2i, start int64, end int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRenderRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) RenderGlyph(cache_index int64, size Vector2i, index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_index), size.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRenderGlyph), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) SetLanguageSupportOverride(language String, supported bool) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), gdc.ConstTypePtr(&supported)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetLanguageSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetLanguageSupportOverride(language String) bool {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetLanguageSupportOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) RemoveLanguageSupportOverride(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRemoveLanguageSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetLanguageSupportOverrides() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetLanguageSupportOverrides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetScriptSupportOverride(script String, supported bool) {
	cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), gdc.ConstTypePtr(&supported)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetScriptSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetScriptSupportOverride(script String) bool {
	cargs := []gdc.ConstTypePtr{script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetScriptSupportOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) RemoveScriptSupportOverride(script String) {
	cargs := []gdc.ConstTypePtr{script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnRemoveScriptSupportOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetScriptSupportOverrides() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetScriptSupportOverrides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) SetOpentypeFeatureOverrides(overrides Dictionary) {
	cargs := []gdc.ConstTypePtr{overrides.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnSetOpentypeFeatureOverrides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FontFile) GetOpentypeFeatureOverrides() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetOpentypeFeatureOverrides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *FontFile) GetGlyphIndex(size int64, char int64, variation_selector int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&variation_selector)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&size)
	pinner.Pin(&char)
	pinner.Pin(&variation_selector)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetGlyphIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FontFile) GetCharFromGlyphIndex(size int64, glyph_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&size)
	pinner.Pin(&glyph_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFontFile.fnGetCharFromGlyphIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
