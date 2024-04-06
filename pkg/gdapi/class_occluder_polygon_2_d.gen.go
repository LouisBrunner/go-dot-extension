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

type ptrsForOccluderPolygon2DList struct {
	fnSetClosed   gdc.MethodBindPtr
	fnIsClosed    gdc.MethodBindPtr
	fnSetCullMode gdc.MethodBindPtr
	fnGetCullMode gdc.MethodBindPtr
	fnSetPolygon  gdc.MethodBindPtr
	fnGetPolygon  gdc.MethodBindPtr
}

var ptrsForOccluderPolygon2D ptrsForOccluderPolygon2DList

func initOccluderPolygon2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OccluderPolygon2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_closed")
		defer methodName.Destroy()
		ptrsForOccluderPolygon2D.fnSetClosed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_closed")
		defer methodName.Destroy()
		ptrsForOccluderPolygon2D.fnIsClosed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_cull_mode")
		defer methodName.Destroy()
		ptrsForOccluderPolygon2D.fnSetCullMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3500863002))
	}
	{
		methodName := StringNameFromStr("get_cull_mode")
		defer methodName.Destroy()
		ptrsForOccluderPolygon2D.fnGetCullMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 33931036))
	}
	{
		methodName := StringNameFromStr("set_polygon")
		defer methodName.Destroy()
		ptrsForOccluderPolygon2D.fnSetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_polygon")
		defer methodName.Destroy()
		ptrsForOccluderPolygon2D.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
}

type OccluderPolygon2D struct {
	Resource
}

func (me *OccluderPolygon2D) BaseClass() string {
	return "OccluderPolygon2D"
}

func NewOccluderPolygon2D() *OccluderPolygon2D {
	str := StringNameFromStr("OccluderPolygon2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OccluderPolygon2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type OccluderPolygon2DCullMode int

const (
	OccluderPolygon2DCullModeCullDisabled         OccluderPolygon2DCullMode = 0
	OccluderPolygon2DCullModeCullClockwise        OccluderPolygon2DCullMode = 1
	OccluderPolygon2DCullModeCullCounterClockwise OccluderPolygon2DCullMode = 2
)

func (me *OccluderPolygon2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OccluderPolygon2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OccluderPolygon2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OccluderPolygon2D) SetClosed(closed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&closed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderPolygon2D.fnSetClosed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OccluderPolygon2D) IsClosed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderPolygon2D.fnIsClosed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OccluderPolygon2D) SetCullMode(cull_mode OccluderPolygon2DCullMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderPolygon2D.fnSetCullMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OccluderPolygon2D) GetCullMode() OccluderPolygon2DCullMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret OccluderPolygon2DCullMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderPolygon2D.fnGetCullMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *OccluderPolygon2D) SetPolygon(polygon PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderPolygon2D.fnSetPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OccluderPolygon2D) GetPolygon() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderPolygon2D.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
