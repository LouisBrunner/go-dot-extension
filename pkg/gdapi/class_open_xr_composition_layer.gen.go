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

type ptrsForOpenXRCompositionLayerList struct {
	fnSetLayerViewport    gdc.MethodBindPtr
	fnGetLayerViewport    gdc.MethodBindPtr
	fnSetSortOrder        gdc.MethodBindPtr
	fnGetSortOrder        gdc.MethodBindPtr
	fnSetAlphaBlend       gdc.MethodBindPtr
	fnGetAlphaBlend       gdc.MethodBindPtr
	fnIsNativelySupported gdc.MethodBindPtr
	fnIntersectsRay       gdc.MethodBindPtr
}

var ptrsForOpenXRCompositionLayer ptrsForOpenXRCompositionLayerList

func initOpenXRCompositionLayerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRCompositionLayer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_layer_viewport")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnSetLayerViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3888077664))
	}
	{
		methodName := StringNameFromStr("get_layer_viewport")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnGetLayerViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3750751911))
	}
	{
		methodName := StringNameFromStr("set_sort_order")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnSetSortOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sort_order")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnGetSortOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_alpha_blend")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnSetAlphaBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_alpha_blend")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnGetAlphaBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_natively_supported")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnIsNativelySupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("intersects_ray")
		defer methodName.Destroy()
		ptrsForOpenXRCompositionLayer.fnIntersectsRay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1091262597))
	}

}

type OpenXRCompositionLayer struct {
	Node3D
}

func (me *OpenXRCompositionLayer) BaseClass() string {
	return "OpenXRCompositionLayer"
}

func NewOpenXRCompositionLayer() *OpenXRCompositionLayer {
	str := StringNameFromStr("OpenXRCompositionLayer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRCompositionLayer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRCompositionLayer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRCompositionLayer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRCompositionLayer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRCompositionLayer) SetLayerViewport(viewport SubViewport) {
	cargs := []gdc.ConstTypePtr{viewport.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnSetLayerViewport), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayer) GetLayerViewport() SubViewport {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSubViewport()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnGetLayerViewport), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRCompositionLayer) SetSortOrder(order int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnSetSortOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayer) GetSortOrder() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnGetSortOrder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayer) SetAlphaBlend(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnSetAlphaBlend), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRCompositionLayer) GetAlphaBlend() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnGetAlphaBlend), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayer) IsNativelySupported() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnIsNativelySupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OpenXRCompositionLayer) IntersectsRay(origin Vector3, direction Vector3) Vector2 {
	cargs := []gdc.ConstTypePtr{origin.AsCTypePtr(), direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRCompositionLayer.fnIntersectsRay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
