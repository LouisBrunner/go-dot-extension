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

type ptrsForMarker2DList struct {
	fnSetGizmoExtents gdc.MethodBindPtr
	fnGetGizmoExtents gdc.MethodBindPtr
}

var ptrsForMarker2D ptrsForMarker2DList

func initMarker2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Marker2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_gizmo_extents")
		defer methodName.Destroy()
		ptrsForMarker2D.fnSetGizmoExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_gizmo_extents")
		defer methodName.Destroy()
		ptrsForMarker2D.fnGetGizmoExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type Marker2D struct {
	Node2D
}

func (me *Marker2D) BaseClass() string {
	return "Marker2D"
}

func NewMarker2D() *Marker2D {
	str := StringNameFromStr("Marker2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Marker2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Marker2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Marker2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Marker2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Marker2D) SetGizmoExtents(extents float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&extents)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarker2D.fnSetGizmoExtents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Marker2D) GetGizmoExtents() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarker2D.fnGetGizmoExtents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
