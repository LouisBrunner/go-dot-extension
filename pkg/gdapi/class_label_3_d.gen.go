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

type ptrsForLabel3DList struct {
	fnSetHorizontalAlignment               gdc.MethodBindPtr
	fnGetHorizontalAlignment               gdc.MethodBindPtr
	fnSetVerticalAlignment                 gdc.MethodBindPtr
	fnGetVerticalAlignment                 gdc.MethodBindPtr
	fnSetModulate                          gdc.MethodBindPtr
	fnGetModulate                          gdc.MethodBindPtr
	fnSetOutlineModulate                   gdc.MethodBindPtr
	fnGetOutlineModulate                   gdc.MethodBindPtr
	fnSetText                              gdc.MethodBindPtr
	fnGetText                              gdc.MethodBindPtr
	fnSetTextDirection                     gdc.MethodBindPtr
	fnGetTextDirection                     gdc.MethodBindPtr
	fnSetLanguage                          gdc.MethodBindPtr
	fnGetLanguage                          gdc.MethodBindPtr
	fnSetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnGetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnSetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnGetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnSetUppercase                         gdc.MethodBindPtr
	fnIsUppercase                          gdc.MethodBindPtr
	fnSetRenderPriority                    gdc.MethodBindPtr
	fnGetRenderPriority                    gdc.MethodBindPtr
	fnSetOutlineRenderPriority             gdc.MethodBindPtr
	fnGetOutlineRenderPriority             gdc.MethodBindPtr
	fnSetFont                              gdc.MethodBindPtr
	fnGetFont                              gdc.MethodBindPtr
	fnSetFontSize                          gdc.MethodBindPtr
	fnGetFontSize                          gdc.MethodBindPtr
	fnSetOutlineSize                       gdc.MethodBindPtr
	fnGetOutlineSize                       gdc.MethodBindPtr
	fnSetLineSpacing                       gdc.MethodBindPtr
	fnGetLineSpacing                       gdc.MethodBindPtr
	fnSetAutowrapMode                      gdc.MethodBindPtr
	fnGetAutowrapMode                      gdc.MethodBindPtr
	fnSetJustificationFlags                gdc.MethodBindPtr
	fnGetJustificationFlags                gdc.MethodBindPtr
	fnSetWidth                             gdc.MethodBindPtr
	fnGetWidth                             gdc.MethodBindPtr
	fnSetPixelSize                         gdc.MethodBindPtr
	fnGetPixelSize                         gdc.MethodBindPtr
	fnSetOffset                            gdc.MethodBindPtr
	fnGetOffset                            gdc.MethodBindPtr
	fnSetDrawFlag                          gdc.MethodBindPtr
	fnGetDrawFlag                          gdc.MethodBindPtr
	fnSetBillboardMode                     gdc.MethodBindPtr
	fnGetBillboardMode                     gdc.MethodBindPtr
	fnSetAlphaCutMode                      gdc.MethodBindPtr
	fnGetAlphaCutMode                      gdc.MethodBindPtr
	fnSetAlphaScissorThreshold             gdc.MethodBindPtr
	fnGetAlphaScissorThreshold             gdc.MethodBindPtr
	fnSetAlphaHashScale                    gdc.MethodBindPtr
	fnGetAlphaHashScale                    gdc.MethodBindPtr
	fnSetAlphaAntialiasing                 gdc.MethodBindPtr
	fnGetAlphaAntialiasing                 gdc.MethodBindPtr
	fnSetAlphaAntialiasingEdge             gdc.MethodBindPtr
	fnGetAlphaAntialiasingEdge             gdc.MethodBindPtr
	fnSetTextureFilter                     gdc.MethodBindPtr
	fnGetTextureFilter                     gdc.MethodBindPtr
	fnGenerateTriangleMesh                 gdc.MethodBindPtr
}

var ptrsForLabel3D ptrsForLabel3DList

func initLabel3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Label3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("set_vertical_alignment")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetVerticalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1796458609))
	}
	{
		methodName := StringNameFromStr("get_vertical_alignment")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetVerticalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3274884059))
	}
	{
		methodName := StringNameFromStr("set_modulate")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_modulate")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_outline_modulate")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetOutlineModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_outline_modulate")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetOutlineModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1418190634))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2516697328))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 55961453))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385126229))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_uppercase")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetUppercase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_uppercase")
		defer methodName.Destroy()
		ptrsForLabel3D.fnIsUppercase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_render_priority")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_render_priority")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_outline_render_priority")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetOutlineRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_outline_render_priority")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetOutlineRenderPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_font")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262170328))
	}
	{
		methodName := StringNameFromStr("get_font")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229501585))
	}
	{
		methodName := StringNameFromStr("set_font_size")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_font_size")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_outline_size")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetOutlineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_outline_size")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetOutlineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_line_spacing")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetLineSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_line_spacing")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetLineSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_autowrap_mode")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289138044))
	}
	{
		methodName := StringNameFromStr("get_autowrap_mode")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549071663))
	}
	{
		methodName := StringNameFromStr("set_justification_flags")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2877345813))
	}
	{
		methodName := StringNameFromStr("get_justification_flags")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1583363614))
	}
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_width")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_pixel_size")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetPixelSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pixel_size")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetPixelSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_draw_flag")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetDrawFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1285833066))
	}
	{
		methodName := StringNameFromStr("get_draw_flag")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetDrawFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259226453))
	}
	{
		methodName := StringNameFromStr("set_billboard_mode")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetBillboardMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4202036497))
	}
	{
		methodName := StringNameFromStr("get_billboard_mode")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetBillboardMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1283840139))
	}
	{
		methodName := StringNameFromStr("set_alpha_cut_mode")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetAlphaCutMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2549142916))
	}
	{
		methodName := StringNameFromStr("get_alpha_cut_mode")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetAlphaCutMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 219468601))
	}
	{
		methodName := StringNameFromStr("set_alpha_scissor_threshold")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetAlphaScissorThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_alpha_scissor_threshold")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetAlphaScissorThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_alpha_hash_scale")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetAlphaHashScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_alpha_hash_scale")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetAlphaHashScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_alpha_antialiasing")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetAlphaAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3212649852))
	}
	{
		methodName := StringNameFromStr("get_alpha_antialiasing")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetAlphaAntialiasing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2889939400))
	}
	{
		methodName := StringNameFromStr("set_alpha_antialiasing_edge")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetAlphaAntialiasingEdge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_alpha_antialiasing_edge")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetAlphaAntialiasingEdge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_texture_filter")
		defer methodName.Destroy()
		ptrsForLabel3D.fnSetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 22904437))
	}
	{
		methodName := StringNameFromStr("get_texture_filter")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289213076))
	}
	{
		methodName := StringNameFromStr("generate_triangle_mesh")
		defer methodName.Destroy()
		ptrsForLabel3D.fnGenerateTriangleMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3476533166))
	}
}

type Label3D struct {
	GeometryInstance3D
}

func (me *Label3D) BaseClass() string {
	return "Label3D"
}

func NewLabel3D() *Label3D {
	str := StringNameFromStr("Label3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Label3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type Label3DDrawFlags int

const (
	Label3DDrawFlagsFlagShaded           Label3DDrawFlags = 0
	Label3DDrawFlagsFlagDoubleSided      Label3DDrawFlags = 1
	Label3DDrawFlagsFlagDisableDepthTest Label3DDrawFlags = 2
	Label3DDrawFlagsFlagFixedSize        Label3DDrawFlags = 3
	Label3DDrawFlagsFlagMax              Label3DDrawFlags = 4
)

type Label3DAlphaCutMode int

const (
	Label3DAlphaCutModeAlphaCutDisabled      Label3DAlphaCutMode = 0
	Label3DAlphaCutModeAlphaCutDiscard       Label3DAlphaCutMode = 1
	Label3DAlphaCutModeAlphaCutOpaquePrepass Label3DAlphaCutMode = 2
	Label3DAlphaCutModeAlphaCutHash          Label3DAlphaCutMode = 3
)

func (me *Label3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Label3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Label3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Label3D) SetHorizontalAlignment(alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetHorizontalAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetVerticalAlignment(alignment VerticalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetVerticalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetVerticalAlignment() VerticalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VerticalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetVerticalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetModulate(modulate Color) {
	cargs := []gdc.ConstTypePtr{modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetModulate() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label3D) SetOutlineModulate(modulate Color) {
	cargs := []gdc.ConstTypePtr{modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetOutlineModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetOutlineModulate() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetOutlineModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label3D) SetText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label3D) SetTextDirection(direction TextServerDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetTextDirection() TextServerDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label3D) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerStructuredTextParser

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetStructuredTextBidiOverrideOptions(args Array) {
	cargs := []gdc.ConstTypePtr{args.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetStructuredTextBidiOverrideOptions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label3D) SetUppercase(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetUppercase), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) IsUppercase() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnIsUppercase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetRenderPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetRenderPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetRenderPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetRenderPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetOutlineRenderPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetOutlineRenderPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetOutlineRenderPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetOutlineRenderPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetFont(font Font) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetFont() Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label3D) SetFontSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetFontSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetOutlineSize(outline_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outline_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetOutlineSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetOutlineSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetOutlineSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetLineSpacing(line_spacing float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_spacing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetLineSpacing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetLineSpacing() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetLineSpacing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetAutowrapMode(autowrap_mode TextServerAutowrapMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetAutowrapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetAutowrapMode() TextServerAutowrapMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerAutowrapMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetAutowrapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetJustificationFlags(justification_flags TextServerJustificationFlag) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&justification_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetJustificationFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetJustificationFlags() TextServerJustificationFlag {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerJustificationFlag

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetJustificationFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetWidth(width float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetPixelSize(pixel_size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixel_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetPixelSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetPixelSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetPixelSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label3D) SetDrawFlag(flag Label3DDrawFlags, enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetDrawFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetDrawFlag(flag Label3DDrawFlags) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetDrawFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetBillboardMode(mode BaseMaterial3DBillboardMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetBillboardMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetBillboardMode() BaseMaterial3DBillboardMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DBillboardMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetBillboardMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetAlphaCutMode(mode Label3DAlphaCutMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetAlphaCutMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetAlphaCutMode() Label3DAlphaCutMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Label3DAlphaCutMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetAlphaCutMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetAlphaScissorThreshold(threshold float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetAlphaScissorThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetAlphaScissorThreshold() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetAlphaScissorThreshold), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetAlphaHashScale(threshold float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetAlphaHashScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetAlphaHashScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetAlphaHashScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alpha_aa)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetAlphaAntialiasing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetAlphaAntialiasing() BaseMaterial3DAlphaAntiAliasing {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DAlphaAntiAliasing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetAlphaAntialiasing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) SetAlphaAntialiasingEdge(edge float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetAlphaAntialiasingEdge), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetAlphaAntialiasingEdge() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetAlphaAntialiasingEdge), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label3D) SetTextureFilter(mode BaseMaterial3DTextureFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnSetTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label3D) GetTextureFilter() BaseMaterial3DTextureFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BaseMaterial3DTextureFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGetTextureFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label3D) GenerateTriangleMesh() TriangleMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTriangleMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel3D.fnGenerateTriangleMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
