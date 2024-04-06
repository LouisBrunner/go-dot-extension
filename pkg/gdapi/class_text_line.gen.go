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

type ptrsForTextLineList struct {
	fnClear                     gdc.MethodBindPtr
	fnSetDirection              gdc.MethodBindPtr
	fnGetDirection              gdc.MethodBindPtr
	fnSetOrientation            gdc.MethodBindPtr
	fnGetOrientation            gdc.MethodBindPtr
	fnSetPreserveInvalid        gdc.MethodBindPtr
	fnGetPreserveInvalid        gdc.MethodBindPtr
	fnSetPreserveControl        gdc.MethodBindPtr
	fnGetPreserveControl        gdc.MethodBindPtr
	fnSetBidiOverride           gdc.MethodBindPtr
	fnAddString                 gdc.MethodBindPtr
	fnAddObject                 gdc.MethodBindPtr
	fnResizeObject              gdc.MethodBindPtr
	fnSetWidth                  gdc.MethodBindPtr
	fnGetWidth                  gdc.MethodBindPtr
	fnSetHorizontalAlignment    gdc.MethodBindPtr
	fnGetHorizontalAlignment    gdc.MethodBindPtr
	fnTabAlign                  gdc.MethodBindPtr
	fnSetFlags                  gdc.MethodBindPtr
	fnGetFlags                  gdc.MethodBindPtr
	fnSetTextOverrunBehavior    gdc.MethodBindPtr
	fnGetTextOverrunBehavior    gdc.MethodBindPtr
	fnGetObjects                gdc.MethodBindPtr
	fnGetObjectRect             gdc.MethodBindPtr
	fnGetSize                   gdc.MethodBindPtr
	fnGetRid                    gdc.MethodBindPtr
	fnGetLineAscent             gdc.MethodBindPtr
	fnGetLineDescent            gdc.MethodBindPtr
	fnGetLineWidth              gdc.MethodBindPtr
	fnGetLineUnderlinePosition  gdc.MethodBindPtr
	fnGetLineUnderlineThickness gdc.MethodBindPtr
	fnDraw                      gdc.MethodBindPtr
	fnDrawOutline               gdc.MethodBindPtr
	fnHitTest                   gdc.MethodBindPtr
}

var ptrsForTextLine ptrsForTextLineList

func initTextLinePtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextLine")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForTextLine.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_direction")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1418190634))
	}
	{
		methodName := StringNameFromStr("get_direction")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2516697328))
	}
	{
		methodName := StringNameFromStr("set_orientation")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 42823726))
	}
	{
		methodName := StringNameFromStr("get_orientation")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 175768116))
	}
	{
		methodName := StringNameFromStr("set_preserve_invalid")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetPreserveInvalid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_preserve_invalid")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetPreserveInvalid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_preserve_control")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetPreserveControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_preserve_control")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetPreserveControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_bidi_override")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("add_string")
		defer methodName.Destroy()
		ptrsForTextLine.fnAddString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 621426851))
	}
	{
		methodName := StringNameFromStr("add_object")
		defer methodName.Destroy()
		ptrsForTextLine.fnAddObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1316529304))
	}
	{
		methodName := StringNameFromStr("resize_object")
		defer methodName.Destroy()
		ptrsForTextLine.fnResizeObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2095776372))
	}
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_width")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("tab_align")
		defer methodName.Destroy()
		ptrsForTextLine.fnTabAlign = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}
	{
		methodName := StringNameFromStr("set_flags")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2877345813))
	}
	{
		methodName := StringNameFromStr("get_flags")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1583363614))
	}
	{
		methodName := StringNameFromStr("set_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForTextLine.fnSetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1008890932))
	}
	{
		methodName := StringNameFromStr("get_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3779142101))
	}
	{
		methodName := StringNameFromStr("get_objects")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_object_rect")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetObjectRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1742700391))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_line_ascent")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetLineAscent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_line_descent")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetLineDescent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_line_width")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetLineWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_line_underline_position")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetLineUnderlinePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_line_underline_thickness")
		defer methodName.Destroy()
		ptrsForTextLine.fnGetLineUnderlineThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("draw")
		defer methodName.Destroy()
		ptrsForTextLine.fnDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 856975658))
	}
	{
		methodName := StringNameFromStr("draw_outline")
		defer methodName.Destroy()
		ptrsForTextLine.fnDrawOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1343401456))
	}
	{
		methodName := StringNameFromStr("hit_test")
		defer methodName.Destroy()
		ptrsForTextLine.fnHitTest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2401831903))
	}
}

type TextLine struct {
	RefCounted
}

func (me *TextLine) BaseClass() string {
	return "TextLine"
}

func NewTextLine() *TextLine {
	str := StringNameFromStr("TextLine") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextLine{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TextLine) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextLine) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextLine) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextLine) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) SetDirection(direction TextServerDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetDirection() TextServerDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextLine) SetOrientation(orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetOrientation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetOrientation() TextServerOrientation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOrientation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetOrientation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextLine) SetPreserveInvalid(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetPreserveInvalid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetPreserveInvalid() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetPreserveInvalid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) SetPreserveControl(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetPreserveControl), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetPreserveControl() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetPreserveControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) SetBidiOverride(override Array) {
	cargs := []gdc.ConstTypePtr{override.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) AddString(text String, font Font, font_size int64, language String, meta Variant) bool {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), font.AsCTypePtr(), gdc.ConstTypePtr(&font_size), language.AsCTypePtr(), meta.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&font_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnAddString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) AddObject(key Variant, size Vector2, inline_align InlineAlignment, length int64, baseline float64) bool {
	cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&length), gdc.ConstTypePtr(&baseline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&inline_align)
	pinner.Pin(&length)
	pinner.Pin(&baseline)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnAddObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) ResizeObject(key Variant, size Vector2, inline_align InlineAlignment, baseline float64) bool {
	cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&baseline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&inline_align)
	pinner.Pin(&baseline)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnResizeObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) SetWidth(width float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) SetHorizontalAlignment(alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetHorizontalAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextLine) TabAlign(tab_stops PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{tab_stops.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnTabAlign), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) SetFlags(flags TextServerJustificationFlag) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetFlags() TextServerJustificationFlag {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerJustificationFlag

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextLine) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnSetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) GetTextOverrunBehavior() TextServerOverrunBehavior {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOverrunBehavior

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextLine) GetObjects() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetObjects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextLine) GetObjectRect(key Variant) Rect2 {
	cargs := []gdc.ConstTypePtr{key.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetObjectRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextLine) GetSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextLine) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextLine) GetLineAscent() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetLineAscent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) GetLineDescent() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetLineDescent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) GetLineWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetLineWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) GetLineUnderlinePosition() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetLineUnderlinePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) GetLineUnderlineThickness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnGetLineUnderlineThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextLine) Draw(canvas RID, pos Vector2, color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) DrawOutline(canvas RID, pos Vector2, outline_size int64, color Color) {
	cargs := []gdc.ConstTypePtr{canvas.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&outline_size), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnDrawOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextLine) HitTest(coords float64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&coords)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&coords)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextLine.fnHitTest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
