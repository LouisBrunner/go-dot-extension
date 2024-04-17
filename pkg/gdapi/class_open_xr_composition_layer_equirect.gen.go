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

type ptrsForOpenXRCompositionLayerEquirectList struct {
	fnSetRadius                 gdc.MethodBindPtr
	fnGetRadius                 gdc.MethodBindPtr
	fnSetCentralHorizontalAngle gdc.MethodBindPtr
	fnGetCentralHorizontalAngle gdc.MethodBindPtr
	fnSetUpperVerticalAngle     gdc.MethodBindPtr
	fnGetUpperVerticalAngle     gdc.MethodBindPtr
	fnSetLowerVerticalAngle     gdc.MethodBindPtr
	fnGetLowerVerticalAngle     gdc.MethodBindPtr
	fnSetFallbackSegments       gdc.MethodBindPtr
	fnGetFallbackSegments       gdc.MethodBindPtr
}

var ptrsForOpenXRCompositionLayerEquirect ptrsForOpenXRCompositionLayerEquirectList

func initOpenXRCompositionLayerEquirectPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRCompositionLayerEquirect")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_central_horizontal_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnSetCentralHorizontalAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_central_horizontal_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnGetCentralHorizontalAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_upper_vertical_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnSetUpperVerticalAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_upper_vertical_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnGetUpperVerticalAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_lower_vertical_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnSetLowerVerticalAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_lower_vertical_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnGetLowerVerticalAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fallback_segments")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnSetFallbackSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_fallback_segments")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerEquirect.fnGetFallbackSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type OpenXRCompositionLayerEquirect struct {
	OpenXRCompositionLayer
}

func (me *OpenXRCompositionLayerEquirect) BaseClass() string {
	return "OpenXRCompositionLayerEquirect"
}

func NewOpenXRCompositionLayerEquirect() *OpenXRCompositionLayerEquirect {
	str := StringNameFromStr("OpenXRCompositionLayerEquirect") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRCompositionLayerEquirect{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRCompositionLayerEquirect) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRCompositionLayerEquirect) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRCompositionLayerEquirect) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRCompositionLayerEquirect) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerEquirect) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayerEquirect) SetCentralHorizontalAngle(angle float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnSetCentralHorizontalAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerEquirect) GetCentralHorizontalAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnGetCentralHorizontalAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayerEquirect) SetUpperVerticalAngle(angle float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnSetUpperVerticalAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerEquirect) GetUpperVerticalAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnGetUpperVerticalAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayerEquirect) SetLowerVerticalAngle(angle float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnSetLowerVerticalAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerEquirect) GetLowerVerticalAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnGetLowerVerticalAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayerEquirect) SetFallbackSegments(segments int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnSetFallbackSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerEquirect) GetFallbackSegments() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerEquirect.fnGetFallbackSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
