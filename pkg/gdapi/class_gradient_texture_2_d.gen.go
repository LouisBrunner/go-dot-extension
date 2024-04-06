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

type ptrsForGradientTexture2DList struct {
	fnSetGradient gdc.MethodBindPtr
	fnGetGradient gdc.MethodBindPtr
	fnSetWidth    gdc.MethodBindPtr
	fnSetHeight   gdc.MethodBindPtr
	fnSetUseHdr   gdc.MethodBindPtr
	fnIsUsingHdr  gdc.MethodBindPtr
	fnSetFill     gdc.MethodBindPtr
	fnGetFill     gdc.MethodBindPtr
	fnSetFillFrom gdc.MethodBindPtr
	fnGetFillFrom gdc.MethodBindPtr
	fnSetFillTo   gdc.MethodBindPtr
	fnGetFillTo   gdc.MethodBindPtr
	fnSetRepeat   gdc.MethodBindPtr
	fnGetRepeat   gdc.MethodBindPtr
}

var ptrsForGradientTexture2D ptrsForGradientTexture2DList

func initGradientTexture2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GradientTexture2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_gradient")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetGradient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
	}
	{
		methodName := StringNameFromStr("get_gradient")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnGetGradient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
	}
	{
		methodName := StringNameFromStr("set_width")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_use_hdr")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetUseHdr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_hdr")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnIsUsingHdr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_fill")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetFill = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3623927636))
	}
	{
		methodName := StringNameFromStr("get_fill")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnGetFill = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1876227217))
	}
	{
		methodName := StringNameFromStr("set_fill_from")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetFillFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_fill_from")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnGetFillFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_fill_to")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetFillTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_fill_to")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnGetFillTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_repeat")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnSetRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1357597002))
	}
	{
		methodName := StringNameFromStr("get_repeat")
		defer methodName.Destroy()
		ptrsForGradientTexture2D.fnGetRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3351758665))
	}

}

type GradientTexture2D struct {
	Texture2D
}

func (me *GradientTexture2D) BaseClass() string {
	return "GradientTexture2D"
}

func NewGradientTexture2D() *GradientTexture2D {
	str := StringNameFromStr("GradientTexture2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GradientTexture2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GradientTexture2DFill int

const (
	GradientTexture2DFillFillLinear GradientTexture2DFill = 0
	GradientTexture2DFillFillRadial GradientTexture2DFill = 1
	GradientTexture2DFillFillSquare GradientTexture2DFill = 2
)

type GradientTexture2DRepeat int

const (
	GradientTexture2DRepeatRepeatNone   GradientTexture2DRepeat = 0
	GradientTexture2DRepeatRepeat       GradientTexture2DRepeat = 1
	GradientTexture2DRepeatRepeatMirror GradientTexture2DRepeat = 2
)

func (me *GradientTexture2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GradientTexture2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GradientTexture2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GradientTexture2D) SetGradient(gradient Gradient) {
	cargs := []gdc.ConstTypePtr{gradient.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetGradient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) GetGradient() Gradient {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGradient()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnGetGradient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GradientTexture2D) SetWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) SetHeight(height int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) SetUseHdr(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetUseHdr), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) IsUsingHdr() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnIsUsingHdr), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GradientTexture2D) SetFill(fill GradientTexture2DFill) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fill)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetFill), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) GetFill() GradientTexture2DFill {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GradientTexture2DFill

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnGetFill), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GradientTexture2D) SetFillFrom(fill_from Vector2) {
	cargs := []gdc.ConstTypePtr{fill_from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetFillFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) GetFillFrom() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnGetFillFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GradientTexture2D) SetFillTo(fill_to Vector2) {
	cargs := []gdc.ConstTypePtr{fill_to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetFillTo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) GetFillTo() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnGetFillTo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GradientTexture2D) SetRepeat(repeat GradientTexture2DRepeat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnSetRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GradientTexture2D) GetRepeat() GradientTexture2DRepeat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GradientTexture2DRepeat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture2D.fnGetRepeat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
