// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FontFile struct {
  obj gdc.ObjectPtr
}

func (me *FontFile) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FontFile) BaseClass() string {
  return "FontFile"
}

func  (me *FontFile) LoadBitmapFont(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) LoadDynamicFont(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetData(data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetData() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetFontName(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetFontStyleName(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetFontStyle(style TextServerFontStyle, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetFontWeight(weight int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetFontStretch(stretch int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetAntialiasing(antialiasing TextServerFontAntialiasing, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetAntialiasing() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetGenerateMipmaps(generate_mipmaps bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGenerateMipmaps() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetMultichannelSignedDistanceField(msdf bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) IsMultichannelSignedDistanceField() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetMsdfPixelRange(msdf_pixel_range int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetMsdfPixelRange() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetMsdfSize(msdf_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetMsdfSize() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetFixedSize(fixed_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetFixedSize() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetAllowSystemFallback(allow_system_fallback bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) IsAllowSystemFallback() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetForceAutohinter(force_autohinter bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) IsForceAutohinter() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetHinting(hinting TextServerHinting, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetHinting() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetSubpixelPositioning(subpixel_positioning TextServerSubpixelPositioning, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetSubpixelPositioning() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetOversampling(oversampling float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetOversampling() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetCacheCount() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) ClearCache() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RemoveCache(cache_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetSizeCacheList(cache_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) ClearSizeCache(cache_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RemoveSizeCache(cache_index int, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetVariationCoordinates(cache_index int, variation_coordinates Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetVariationCoordinates(cache_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetEmbolden(cache_index int, strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetEmbolden(cache_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetTransform(cache_index int, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetTransform(cache_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetFaceIndex(cache_index int, face_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetFaceIndex(cache_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetCacheAscent(cache_index int, size int, ascent float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetCacheAscent(cache_index int, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetCacheDescent(cache_index int, size int, descent float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetCacheDescent(cache_index int, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetCacheUnderlinePosition(cache_index int, size int, underline_position float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetCacheUnderlinePosition(cache_index int, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetCacheUnderlineThickness(cache_index int, size int, underline_thickness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetCacheUnderlineThickness(cache_index int, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetCacheScale(cache_index int, size int, scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetCacheScale(cache_index int, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetTextureCount(cache_index int, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) ClearTextures(cache_index int, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RemoveTexture(cache_index int, size Vector2i, texture_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetTextureImage(cache_index int, size Vector2i, texture_index int, image Image, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetTextureImage(cache_index int, size Vector2i, texture_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetTextureOffsets(cache_index int, size Vector2i, texture_index int, offset PackedInt32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetTextureOffsets(cache_index int, size Vector2i, texture_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGlyphList(cache_index int, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) ClearGlyphs(cache_index int, size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RemoveGlyph(cache_index int, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetGlyphAdvance(cache_index int, size int, glyph int, advance Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGlyphAdvance(cache_index int, size int, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetGlyphOffset(cache_index int, size Vector2i, glyph int, offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGlyphOffset(cache_index int, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetGlyphSize(cache_index int, size Vector2i, glyph int, gl_size Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGlyphSize(cache_index int, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetGlyphUvRect(cache_index int, size Vector2i, glyph int, uv_rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGlyphUvRect(cache_index int, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetGlyphTextureIdx(cache_index int, size Vector2i, glyph int, texture_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGlyphTextureIdx(cache_index int, size Vector2i, glyph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetKerningList(cache_index int, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) ClearKerningMap(cache_index int, size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RemoveKerning(cache_index int, size int, glyph_pair Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetKerning(cache_index int, size int, glyph_pair Vector2i, kerning Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetKerning(cache_index int, size int, glyph_pair Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RenderRange(cache_index int, size Vector2i, start int, end int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RenderGlyph(cache_index int, size Vector2i, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetLanguageSupportOverride(language String, supported bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetLanguageSupportOverride(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RemoveLanguageSupportOverride(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetLanguageSupportOverrides() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetScriptSupportOverride(script String, supported bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetScriptSupportOverride(script String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) RemoveScriptSupportOverride(script String, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetScriptSupportOverrides() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) SetOpentypeFeatureOverrides(overrides Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetOpentypeFeatureOverrides() { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetGlyphIndex(size int, char int, variation_selector int, ) { // TODO: return value
  // TODO: implement
}

func  (me *FontFile) GetCharFromGlyphIndex(size int, glyph_index int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
