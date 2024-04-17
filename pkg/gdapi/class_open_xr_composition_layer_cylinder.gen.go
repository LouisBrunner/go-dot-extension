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

type ptrsForOpenXRCompositionLayerCylinderList struct {
	fnSetRadius           gdc.MethodBindPtr
	fnGetRadius           gdc.MethodBindPtr
	fnSetAspectRatio      gdc.MethodBindPtr
	fnGetAspectRatio      gdc.MethodBindPtr
	fnSetCentralAngle     gdc.MethodBindPtr
	fnGetCentralAngle     gdc.MethodBindPtr
	fnSetFallbackSegments gdc.MethodBindPtr
	fnGetFallbackSegments gdc.MethodBindPtr
}

var ptrsForOpenXRCompositionLayerCylinder ptrsForOpenXRCompositionLayerCylinderList

func initOpenXRCompositionLayerCylinderPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRCompositionLayerCylinder")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_aspect_ratio")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnSetAspectRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_aspect_ratio")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnGetAspectRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_central_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnSetCentralAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_central_angle")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnGetCentralAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fallback_segments")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnSetFallbackSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_fallback_segments")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayerCylinder.fnGetFallbackSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type OpenXRCompositionLayerCylinder struct {
	OpenXRCompositionLayer
}

func (me *OpenXRCompositionLayerCylinder) BaseClass() string {
	return "OpenXRCompositionLayerCylinder"
}

func NewOpenXRCompositionLayerCylinder() *OpenXRCompositionLayerCylinder {
	str := StringNameFromStr("OpenXRCompositionLayerCylinder") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRCompositionLayerCylinder{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRCompositionLayerCylinder) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRCompositionLayerCylinder) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRCompositionLayerCylinder) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRCompositionLayerCylinder) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerCylinder) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayerCylinder) SetAspectRatio(aspect_ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aspect_ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnSetAspectRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerCylinder) GetAspectRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnGetAspectRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayerCylinder) SetCentralAngle(angle float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnSetCentralAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerCylinder) GetCentralAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnGetCentralAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayerCylinder) SetFallbackSegments(segments int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnSetFallbackSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayerCylinder) GetFallbackSegments() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayerCylinder.fnGetFallbackSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
