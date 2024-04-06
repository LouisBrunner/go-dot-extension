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

type ptrsForPathFollow2DList struct {
	fnSetProgress           gdc.MethodBindPtr
	fnGetProgress           gdc.MethodBindPtr
	fnSetHOffset            gdc.MethodBindPtr
	fnGetHOffset            gdc.MethodBindPtr
	fnSetVOffset            gdc.MethodBindPtr
	fnGetVOffset            gdc.MethodBindPtr
	fnSetProgressRatio      gdc.MethodBindPtr
	fnGetProgressRatio      gdc.MethodBindPtr
	fnSetRotates            gdc.MethodBindPtr
	fnIsRotating            gdc.MethodBindPtr
	fnSetCubicInterpolation gdc.MethodBindPtr
	fnGetCubicInterpolation gdc.MethodBindPtr
	fnSetLoop               gdc.MethodBindPtr
	fnHasLoop               gdc.MethodBindPtr
}

var ptrsForPathFollow2D ptrsForPathFollow2DList

func initPathFollow2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PathFollow2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_progress")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnSetProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_progress")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnGetProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_h_offset")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnSetHOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_h_offset")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnGetHOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_v_offset")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnSetVOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_v_offset")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnGetVOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_progress_ratio")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnSetProgressRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_progress_ratio")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnGetProgressRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_rotates")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnSetRotates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_rotating")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnIsRotating = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_cubic_interpolation")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnSetCubicInterpolation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_cubic_interpolation")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnGetCubicInterpolation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_loop")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_loop")
		defer methodName.Destroy()
		ptrsForPathFollow2D.fnHasLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type PathFollow2D struct {
	Node2D
}

func (me *PathFollow2D) BaseClass() string {
	return "PathFollow2D"
}

func NewPathFollow2D() *PathFollow2D {
	str := StringNameFromStr("PathFollow2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PathFollow2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PathFollow2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PathFollow2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PathFollow2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PathFollow2D) SetProgress(progress float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnSetProgress), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PathFollow2D) GetProgress() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnGetProgress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PathFollow2D) SetHOffset(h_offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&h_offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnSetHOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PathFollow2D) GetHOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnGetHOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PathFollow2D) SetVOffset(v_offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&v_offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnSetVOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PathFollow2D) GetVOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnGetVOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PathFollow2D) SetProgressRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnSetProgressRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PathFollow2D) GetProgressRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnGetProgressRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PathFollow2D) SetRotates(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnSetRotates), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PathFollow2D) IsRotating() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnIsRotating), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PathFollow2D) SetCubicInterpolation(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnSetCubicInterpolation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PathFollow2D) GetCubicInterpolation() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnGetCubicInterpolation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PathFollow2D) SetLoop(loop bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PathFollow2D) HasLoop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPathFollow2D.fnHasLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
