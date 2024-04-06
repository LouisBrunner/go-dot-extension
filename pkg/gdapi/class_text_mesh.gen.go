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

type ptrsForTextMeshList struct {
	fnSetHorizontalAlignment               gdc.MethodBindPtr
	fnGetHorizontalAlignment               gdc.MethodBindPtr
	fnSetVerticalAlignment                 gdc.MethodBindPtr
	fnGetVerticalAlignment                 gdc.MethodBindPtr
	fnSetText                              gdc.MethodBindPtr
	fnGetText                              gdc.MethodBindPtr
	fnSetFont                              gdc.MethodBindPtr
	fnGetFont                              gdc.MethodBindPtr
	fnSetFontSize                          gdc.MethodBindPtr
	fnGetFontSize                          gdc.MethodBindPtr
	fnSetLineSpacing                       gdc.MethodBindPtr
	fnGetLineSpacing                       gdc.MethodBindPtr
	fnSetAutowrapMode                      gdc.MethodBindPtr
	fnGetAutowrapMode                      gdc.MethodBindPtr
	fnSetJustificationFlags                gdc.MethodBindPtr
	fnGetJustificationFlags                gdc.MethodBindPtr
	fnSetDepth                             gdc.MethodBindPtr
	fnGetDepth                             gdc.MethodBindPtr
	fnSetWidth                             gdc.MethodBindPtr
	fnGetWidth                             gdc.MethodBindPtr
	fnSetPixelSize                         gdc.MethodBindPtr
	fnGetPixelSize                         gdc.MethodBindPtr
	fnSetOffset                            gdc.MethodBindPtr
	fnGetOffset                            gdc.MethodBindPtr
	fnSetCurveStep                         gdc.MethodBindPtr
	fnGetCurveStep                         gdc.MethodBindPtr
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
}

var ptrsForTextMesh ptrsForTextMeshList

func initTextMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("set_vertical_alignment")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetVerticalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1796458609))
	}
	{
		methodName := StringNameFromStr("get_vertical_alignment")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetVerticalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3274884059))
	}
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_font")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262170328))
	}
	{
		methodName := StringNameFromStr("get_font")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229501585))
	}
	{
		methodName := StringNameFromStr("set_font_size")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_font_size")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_line_spacing")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetLineSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_line_spacing")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetLineSpacing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_autowrap_mode")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289138044))
	}
	{
		methodName := StringNameFromStr("get_autowrap_mode")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549071663))
	}
	{
		methodName := StringNameFromStr("set_justification_flags")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2877345813))
	}
	{
		methodName := StringNameFromStr("get_justification_flags")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1583363614))
	}
	{
		methodName := StringNameFromStr("set_depth")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_width")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_pixel_size")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetPixelSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pixel_size")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetPixelSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_curve_step")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetCurveStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_curve_step")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetCurveStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1418190634))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2516697328))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 55961453))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385126229))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForTextMesh.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_uppercase")
		defer methodName.Destroy()
		ptrsForTextMesh.fnSetUppercase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_uppercase")
		defer methodName.Destroy()
		ptrsForTextMesh.fnIsUppercase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type TextMesh struct {
	PrimitiveMesh
}

func (me *TextMesh) BaseClass() string {
	return "TextMesh"
}

func NewTextMesh() *TextMesh {
	str := StringNameFromStr("TextMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TextMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextMesh) SetHorizontalAlignment(alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetHorizontalAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextMesh) SetVerticalAlignment(alignment VerticalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetVerticalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetVerticalAlignment() VerticalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VerticalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetVerticalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextMesh) SetText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextMesh) SetFont(font Font) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetFont() Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextMesh) SetFontSize(font_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetFontSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextMesh) SetLineSpacing(line_spacing float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_spacing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetLineSpacing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetLineSpacing() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetLineSpacing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextMesh) SetAutowrapMode(autowrap_mode TextServerAutowrapMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetAutowrapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetAutowrapMode() TextServerAutowrapMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerAutowrapMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetAutowrapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextMesh) SetJustificationFlags(justification_flags TextServerJustificationFlag) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&justification_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetJustificationFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetJustificationFlags() TextServerJustificationFlag {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerJustificationFlag

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetJustificationFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextMesh) SetDepth(depth float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextMesh) SetWidth(width float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextMesh) SetPixelSize(pixel_size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixel_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetPixelSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetPixelSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetPixelSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextMesh) SetOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextMesh) SetCurveStep(curve_step float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve_step)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetCurveStep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetCurveStep() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetCurveStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextMesh) SetTextDirection(direction TextServerDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetTextDirection() TextServerDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextMesh) SetLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextMesh) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerStructuredTextParser

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextMesh) SetStructuredTextBidiOverrideOptions(args Array) {
	cargs := []gdc.ConstTypePtr{args.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) GetStructuredTextBidiOverrideOptions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextMesh) SetUppercase(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnSetUppercase), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextMesh) IsUppercase() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextMesh.fnIsUppercase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
