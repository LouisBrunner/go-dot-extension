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

type ptrsForFontList struct {
	fnSetFallbacks               gdc.MethodBindPtr
	fnGetFallbacks               gdc.MethodBindPtr
	fnFindVariation              gdc.MethodBindPtr
	fnGetRids                    gdc.MethodBindPtr
	fnGetHeight                  gdc.MethodBindPtr
	fnGetAscent                  gdc.MethodBindPtr
	fnGetDescent                 gdc.MethodBindPtr
	fnGetUnderlinePosition       gdc.MethodBindPtr
	fnGetUnderlineThickness      gdc.MethodBindPtr
	fnGetFontName                gdc.MethodBindPtr
	fnGetFontStyleName           gdc.MethodBindPtr
	fnGetOtNameStrings           gdc.MethodBindPtr
	fnGetFontStyle               gdc.MethodBindPtr
	fnGetFontWeight              gdc.MethodBindPtr
	fnGetFontStretch             gdc.MethodBindPtr
	fnGetSpacing                 gdc.MethodBindPtr
	fnGetOpentypeFeatures        gdc.MethodBindPtr
	fnSetCacheCapacity           gdc.MethodBindPtr
	fnGetStringSize              gdc.MethodBindPtr
	fnGetMultilineStringSize     gdc.MethodBindPtr
	fnDrawString                 gdc.MethodBindPtr
	fnDrawMultilineString        gdc.MethodBindPtr
	fnDrawStringOutline          gdc.MethodBindPtr
	fnDrawMultilineStringOutline gdc.MethodBindPtr
	fnGetCharSize                gdc.MethodBindPtr
	fnDrawChar                   gdc.MethodBindPtr
	fnDrawCharOutline            gdc.MethodBindPtr
	fnHasChar                    gdc.MethodBindPtr
	fnGetSupportedChars          gdc.MethodBindPtr
	fnIsLanguageSupported        gdc.MethodBindPtr
	fnIsScriptSupported          gdc.MethodBindPtr
	fnGetSupportedFeatureList    gdc.MethodBindPtr
	fnGetSupportedVariationList  gdc.MethodBindPtr
	fnGetFaceCount               gdc.MethodBindPtr
}

var ptrsForFont ptrsForFontList

func initFontPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Font")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_fallbacks")
		defer methodName.Destroy()
		ptrsForFont.fnSetFallbacks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_fallbacks")
		defer methodName.Destroy()
		ptrsForFont.fnGetFallbacks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("find_variation")
		defer methodName.Destroy()
		ptrsForFont.fnFindVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2553855095))
	}
	{
		methodName := StringNameFromStr("get_rids")
		defer methodName.Destroy()
		ptrsForFont.fnGetRids = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForFont.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 378113874))
	}
	{
		methodName := StringNameFromStr("get_ascent")
		defer methodName.Destroy()
		ptrsForFont.fnGetAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 378113874))
	}
	{
		methodName := StringNameFromStr("get_descent")
		defer methodName.Destroy()
		ptrsForFont.fnGetDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 378113874))
	}
	{
		methodName := StringNameFromStr("get_underline_position")
		defer methodName.Destroy()
		ptrsForFont.fnGetUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 378113874))
	}
	{
		methodName := StringNameFromStr("get_underline_thickness")
		defer methodName.Destroy()
		ptrsForFont.fnGetUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 378113874))
	}
	{
		methodName := StringNameFromStr("get_font_name")
		defer methodName.Destroy()
		ptrsForFont.fnGetFontName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_font_style_name")
		defer methodName.Destroy()
		ptrsForFont.fnGetFontStyleName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_ot_name_strings")
		defer methodName.Destroy()
		ptrsForFont.fnGetOtNameStrings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_font_style")
		defer methodName.Destroy()
		ptrsForFont.fnGetFontStyle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2520224254))
	}
	{
		methodName := StringNameFromStr("get_font_weight")
		defer methodName.Destroy()
		ptrsForFont.fnGetFontWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_font_stretch")
		defer methodName.Destroy()
		ptrsForFont.fnGetFontStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_spacing")
		defer methodName.Destroy()
		ptrsForFont.fnGetSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1310880908))
	}
	{
		methodName := StringNameFromStr("get_opentype_features")
		defer methodName.Destroy()
		ptrsForFont.fnGetOpentypeFeatures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("set_cache_capacity")
		defer methodName.Destroy()
		ptrsForFont.fnSetCacheCapacity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_string_size")
		defer methodName.Destroy()
		ptrsForFont.fnGetStringSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1868866121))
	}
	{
		methodName := StringNameFromStr("get_multiline_string_size")
		defer methodName.Destroy()
		ptrsForFont.fnGetMultilineStringSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 519636710))
	}
	{
		methodName := StringNameFromStr("draw_string")
		defer methodName.Destroy()
		ptrsForFont.fnDrawString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1983721962))
	}
	{
		methodName := StringNameFromStr("draw_multiline_string")
		defer methodName.Destroy()
		ptrsForFont.fnDrawMultilineString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1171506176))
	}
	{
		methodName := StringNameFromStr("draw_string_outline")
		defer methodName.Destroy()
		ptrsForFont.fnDrawStringOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 623754045))
	}
	{
		methodName := StringNameFromStr("draw_multiline_string_outline")
		defer methodName.Destroy()
		ptrsForFont.fnDrawMultilineStringOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3206388178))
	}
	{
		methodName := StringNameFromStr("get_char_size")
		defer methodName.Destroy()
		ptrsForFont.fnGetCharSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3016396712))
	}
	{
		methodName := StringNameFromStr("draw_char")
		defer methodName.Destroy()
		ptrsForFont.fnDrawChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3815617597))
	}
	{
		methodName := StringNameFromStr("draw_char_outline")
		defer methodName.Destroy()
		ptrsForFont.fnDrawCharOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 209525354))
	}
	{
		methodName := StringNameFromStr("has_char")
		defer methodName.Destroy()
		ptrsForFont.fnHasChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_supported_chars")
		defer methodName.Destroy()
		ptrsForFont.fnGetSupportedChars = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_language_supported")
		defer methodName.Destroy()
		ptrsForFont.fnIsLanguageSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("is_script_supported")
		defer methodName.Destroy()
		ptrsForFont.fnIsScriptSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("get_supported_feature_list")
		defer methodName.Destroy()
		ptrsForFont.fnGetSupportedFeatureList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_supported_variation_list")
		defer methodName.Destroy()
		ptrsForFont.fnGetSupportedVariationList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_face_count")
		defer methodName.Destroy()
		ptrsForFont.fnGetFaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type Font struct {
	Resource
}

func (me *Font) BaseClass() string {
	return "Font"
}

func NewFont() *Font {
	str := StringNameFromStr("Font") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Font{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *Font) SetFallbacks(fallbacks []Font) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fallbacks)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnSetFallbacks), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Font) GetFallbacks() []Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetFallbacks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Font](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Font) FindVariation(variation_coordinates Dictionary, face_index int64, strength float64, transform Transform2D, spacing_top int64, spacing_bottom int64, spacing_space int64, spacing_glyph int64, baseline_offset float64) RID {
	cargs := []gdc.ConstTypePtr{variation_coordinates.AsCTypePtr(), gdc.ConstTypePtr(&face_index), gdc.ConstTypePtr(&strength), transform.AsCTypePtr(), gdc.ConstTypePtr(&spacing_top), gdc.ConstTypePtr(&spacing_bottom), gdc.ConstTypePtr(&spacing_space), gdc.ConstTypePtr(&spacing_glyph), gdc.ConstTypePtr(&baseline_offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&face_index)
	pinner.Pin(&strength)
	pinner.Pin(&spacing_top)
	pinner.Pin(&spacing_bottom)
	pinner.Pin(&spacing_space)
	pinner.Pin(&spacing_glyph)
	pinner.Pin(&baseline_offset)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnFindVariation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) GetRids() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetRids), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Font) GetHeight(font_size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetAscent(font_size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetAscent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetDescent(font_size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetDescent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetUnderlinePosition(font_size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetUnderlinePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetUnderlineThickness(font_size int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetUnderlineThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetFontName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetFontName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) GetFontStyleName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetFontStyleName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) GetOtNameStrings() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetOtNameStrings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) GetFontStyle() TextServerFontStyle {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerFontStyle

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetFontStyle), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Font) GetFontWeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetFontWeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetFontStretch() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetFontStretch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetSpacing(spacing TextServerSpacingType) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spacing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&spacing)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetSpacing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetOpentypeFeatures() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetOpentypeFeatures), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) SetCacheCapacity(single_line int64, multi_line int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&single_line), gdc.ConstTypePtr(&multi_line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnSetCacheCapacity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Font) GetStringSize(text String, alignment HorizontalAlignment, width float64, font_size int64, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) Vector2 {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&alignment)
	pinner.Pin(&width)
	pinner.Pin(&font_size)
	pinner.Pin(&justification_flags)
	pinner.Pin(&direction)
	pinner.Pin(&orientation)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetStringSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) GetMultilineStringSize(text String, alignment HorizontalAlignment, width float64, font_size int64, max_lines int64, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) Vector2 {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&alignment)
	pinner.Pin(&width)
	pinner.Pin(&font_size)
	pinner.Pin(&max_lines)
	pinner.Pin(&brk_flags)
	pinner.Pin(&justification_flags)
	pinner.Pin(&direction)
	pinner.Pin(&orientation)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetMultilineStringSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) DrawString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), modulate.AsCTypePtr(), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnDrawString), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Font) DrawMultilineString(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, max_lines int64, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), modulate.AsCTypePtr(), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnDrawMultilineString), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Font) DrawStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, size int64, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), modulate.AsCTypePtr(), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnDrawStringOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Font) DrawMultilineStringOutline(canvas_item RID, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, max_lines int64, size int64, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(&size), modulate.AsCTypePtr(), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnDrawMultilineStringOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Font) GetCharSize(char int64, font_size int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&char)
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetCharSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) DrawChar(canvas_item RID, pos Vector2, char int64, font_size int64, modulate Color) float64 {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&font_size), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&char)
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnDrawChar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) DrawCharOutline(canvas_item RID, pos Vector2, char int64, font_size int64, size int64, modulate Color) float64 {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&char)
	pinner.Pin(&font_size)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnDrawCharOutline), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) HasChar(char int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&char)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&char)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnHasChar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetSupportedChars() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetSupportedChars), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) IsLanguageSupported(language String) bool {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnIsLanguageSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) IsScriptSupported(script String) bool {
	cargs := []gdc.ConstTypePtr{script.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnIsScriptSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Font) GetSupportedFeatureList() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetSupportedFeatureList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) GetSupportedVariationList() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetSupportedVariationList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Font) GetFaceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFont.fnGetFaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
