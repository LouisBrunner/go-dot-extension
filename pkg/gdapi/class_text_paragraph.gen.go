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

type ptrsForTextParagraphList struct {
	fnClear                     gdc.MethodBindPtr
	fnSetDirection              gdc.MethodBindPtr
	fnGetDirection              gdc.MethodBindPtr
	fnSetCustomPunctuation      gdc.MethodBindPtr
	fnGetCustomPunctuation      gdc.MethodBindPtr
	fnSetOrientation            gdc.MethodBindPtr
	fnGetOrientation            gdc.MethodBindPtr
	fnSetPreserveInvalid        gdc.MethodBindPtr
	fnGetPreserveInvalid        gdc.MethodBindPtr
	fnSetPreserveControl        gdc.MethodBindPtr
	fnGetPreserveControl        gdc.MethodBindPtr
	fnSetBidiOverride           gdc.MethodBindPtr
	fnSetDropcap                gdc.MethodBindPtr
	fnClearDropcap              gdc.MethodBindPtr
	fnAddString                 gdc.MethodBindPtr
	fnAddObject                 gdc.MethodBindPtr
	fnResizeObject              gdc.MethodBindPtr
	fnSetAlignment              gdc.MethodBindPtr
	fnGetAlignment              gdc.MethodBindPtr
	fnTabAlign                  gdc.MethodBindPtr
	fnSetBreakFlags             gdc.MethodBindPtr
	fnGetBreakFlags             gdc.MethodBindPtr
	fnSetJustificationFlags     gdc.MethodBindPtr
	fnGetJustificationFlags     gdc.MethodBindPtr
	fnSetTextOverrunBehavior    gdc.MethodBindPtr
	fnGetTextOverrunBehavior    gdc.MethodBindPtr
	fnSetEllipsisChar           gdc.MethodBindPtr
	fnGetEllipsisChar           gdc.MethodBindPtr
	fnSetWidth                  gdc.MethodBindPtr
	fnGetWidth                  gdc.MethodBindPtr
	fnGetNonWrappedSize         gdc.MethodBindPtr
	fnGetSize                   gdc.MethodBindPtr
	fnGetRid                    gdc.MethodBindPtr
	fnGetLineRid                gdc.MethodBindPtr
	fnGetDropcapRid             gdc.MethodBindPtr
	fnGetLineCount              gdc.MethodBindPtr
	fnSetMaxLinesVisible        gdc.MethodBindPtr
	fnGetMaxLinesVisible        gdc.MethodBindPtr
	fnGetLineObjects            gdc.MethodBindPtr
	fnGetLineObjectRect         gdc.MethodBindPtr
	fnGetLineSize               gdc.MethodBindPtr
	fnGetLineRange              gdc.MethodBindPtr
	fnGetLineAscent             gdc.MethodBindPtr
	fnGetLineDescent            gdc.MethodBindPtr
	fnGetLineWidth              gdc.MethodBindPtr
	fnGetLineUnderlinePosition  gdc.MethodBindPtr
	fnGetLineUnderlineThickness gdc.MethodBindPtr
	fnGetDropcapSize            gdc.MethodBindPtr
	fnGetDropcapLines           gdc.MethodBindPtr
	fnDraw                      gdc.MethodBindPtr
	fnDrawOutline               gdc.MethodBindPtr
	fnDrawLine                  gdc.MethodBindPtr
	fnDrawLineOutline           gdc.MethodBindPtr
	fnDrawDropcap               gdc.MethodBindPtr
	fnDrawDropcapOutline        gdc.MethodBindPtr
	fnHitTest                   gdc.MethodBindPtr
}

var ptrsForTextParagraph ptrsForTextParagraphList

func initTextParagraphPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextParagraph")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_direction")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1418190634))
	}
	{
		methodName := StringNameFromStr("get_direction")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2516697328))
	}
	{
		methodName := StringNameFromStr("set_custom_punctuation")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetCustomPunctuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_custom_punctuation")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetCustomPunctuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_orientation")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 42823726))
	}
	{
		methodName := StringNameFromStr("get_orientation")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 175768116))
	}
	{
		methodName := StringNameFromStr("set_preserve_invalid")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetPreserveInvalid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_preserve_invalid")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetPreserveInvalid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_preserve_control")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetPreserveControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_preserve_control")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetPreserveControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_bidi_override")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("set_dropcap")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetDropcap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2498990330))
	}
	{
		methodName := StringNameFromStr("clear_dropcap")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnClearDropcap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_string")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnAddString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 621426851))
	}
	{
		methodName := StringNameFromStr("add_object")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnAddObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1316529304))
	}
	{
		methodName := StringNameFromStr("resize_object")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnResizeObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2095776372))
	}
	{
		methodName := StringNameFromStr("set_alignment")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_alignment")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("tab_align")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnTabAlign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}
	{
		methodName := StringNameFromStr("set_break_flags")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetBreakFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2809697122))
	}
	{
		methodName := StringNameFromStr("get_break_flags")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetBreakFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2340632602))
	}
	{
		methodName := StringNameFromStr("set_justification_flags")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2877345813))
	}
	{
		methodName := StringNameFromStr("get_justification_flags")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1583363614))
	}
	{
		methodName := StringNameFromStr("set_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1008890932))
	}
	{
		methodName := StringNameFromStr("get_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3779142101))
	}
	{
		methodName := StringNameFromStr("set_ellipsis_char")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetEllipsisChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_ellipsis_char")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetEllipsisChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_width")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_non_wrapped_size")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetNonWrappedSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_line_rid")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 495598643))
	}
	{
		methodName := StringNameFromStr("get_dropcap_rid")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetDropcapRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_line_count")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_lines_visible")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnSetMaxLinesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_lines_visible")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetMaxLinesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_line_objects")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("get_line_object_rect")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineObjectRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 204315017))
	}
	{
		methodName := StringNameFromStr("get_line_size")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("get_line_range")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 880721226))
	}
	{
		methodName := StringNameFromStr("get_line_ascent")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("get_line_descent")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("get_line_width")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("get_line_underline_position")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("get_line_underline_thickness")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetLineUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("get_dropcap_size")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetDropcapSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_dropcap_lines")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnGetDropcapLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("draw")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1567802413))
	}
	{
		methodName := StringNameFromStr("draw_outline")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnDrawOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1893131224))
	}
	{
		methodName := StringNameFromStr("draw_line")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnDrawLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1242169894))
	}
	{
		methodName := StringNameFromStr("draw_line_outline")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnDrawLineOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2664926980))
	}
	{
		methodName := StringNameFromStr("draw_dropcap")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnDrawDropcap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 856975658))
	}
	{
		methodName := StringNameFromStr("draw_dropcap_outline")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnDrawDropcapOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1343401456))
	}
	{
		methodName := StringNameFromStr("hit_test")
		defer methodName.Destroy()
		ptrsForTextParagraph.fnHitTest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3820158470))
	}

}

type TextParagraph struct {
	RefCounted
}

func (me *TextParagraph) BaseClass() string {
	return "TextParagraph"
}

func NewTextParagraph() *TextParagraph {
	str := StringNameFromStr("TextParagraph") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextParagraph{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TextParagraph) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextParagraph) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextParagraph) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextParagraph) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) SetDirection(direction TextServerDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetDirection() TextServerDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextParagraph) SetCustomPunctuation(custom_punctuation String) {
	cargs := []gdc.ConstTypePtr{custom_punctuation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetCustomPunctuation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetCustomPunctuation() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetCustomPunctuation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) SetOrientation(orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetOrientation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetOrientation() TextServerOrientation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOrientation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetOrientation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextParagraph) SetPreserveInvalid(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetPreserveInvalid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetPreserveInvalid() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetPreserveInvalid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) SetPreserveControl(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetPreserveControl), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetPreserveControl() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetPreserveControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) SetBidiOverride(override Array) {
	cargs := []gdc.ConstTypePtr{override.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) SetDropcap(text String, font Font, font_size int64, dropcap_margins Rect2, language String) bool {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), font.AsCTypePtr(), gdc.ConstTypePtr(&font_size), dropcap_margins.AsCTypePtr(), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetDropcap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) ClearDropcap() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnClearDropcap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) AddString(text String, font Font, font_size int64, language String, meta Variant) bool {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), font.AsCTypePtr(), gdc.ConstTypePtr(&font_size), language.AsCTypePtr(), meta.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnAddString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) AddObject(key Variant, size Vector2, inline_align InlineAlignment, length int64, baseline float64) bool {
	cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&length), gdc.ConstTypePtr(&baseline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&inline_align)
	pinner.Pin(&length)
	pinner.Pin(&baseline)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnAddObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) ResizeObject(key Variant, size Vector2, inline_align InlineAlignment, baseline float64) bool {
	cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&baseline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&inline_align)
	pinner.Pin(&baseline)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnResizeObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) SetAlignment(alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextParagraph) TabAlign(tab_stops PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{tab_stops.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnTabAlign), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) SetBreakFlags(flags TextServerLineBreakFlag) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetBreakFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetBreakFlags() TextServerLineBreakFlag {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerLineBreakFlag

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetBreakFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextParagraph) SetJustificationFlags(flags TextServerJustificationFlag) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetJustificationFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetJustificationFlags() TextServerJustificationFlag {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerJustificationFlag

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetJustificationFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextParagraph) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetTextOverrunBehavior() TextServerOverrunBehavior {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOverrunBehavior

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextParagraph) SetEllipsisChar(char String) {
	cargs := []gdc.ConstTypePtr{char.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetEllipsisChar), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetEllipsisChar() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetEllipsisChar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) SetWidth(width float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) GetNonWrappedSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetNonWrappedSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetLineRid(line int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetDropcapRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetDropcapRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) SetMaxLinesVisible(max_lines_visible int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_lines_visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnSetMaxLinesVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) GetMaxLinesVisible() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetMaxLinesVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) GetLineObjects(line int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineObjects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetLineObjectRect(line int64, key Variant) Rect2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), key.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineObjectRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetLineSize(line int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetLineRange(line int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetLineAscent(line int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineAscent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) GetLineDescent(line int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineDescent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) GetLineWidth(line int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) GetLineUnderlinePosition(line int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineUnderlinePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) GetLineUnderlineThickness(line int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetLineUnderlineThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) GetDropcapSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetDropcapSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextParagraph) GetDropcapLines() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnGetDropcapLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextParagraph) Draw(canvas RID, pos Vector2, color Color, dc_color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), color.AsCTypePtr(), dc_color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) DrawOutline(canvas RID, pos Vector2, outline_size int64, color Color, dc_color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&outline_size), color.AsCTypePtr(), dc_color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnDrawOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) DrawLine(canvas RID, pos Vector2, line int64, color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&line), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnDrawLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) DrawLineOutline(canvas RID, pos Vector2, line int64, outline_size int64, color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&outline_size), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnDrawLineOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) DrawDropcap(canvas RID, pos Vector2, color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnDrawDropcap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) DrawDropcapOutline(canvas RID, pos Vector2, outline_size int64, color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&outline_size), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnDrawDropcapOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextParagraph) HitTest(coords Vector2) int64 {
	cargs := []gdc.ConstTypePtr{coords.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextParagraph.fnHitTest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
