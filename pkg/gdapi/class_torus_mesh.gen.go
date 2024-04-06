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

type ptrsForTorusMeshList struct {
	fnSetInnerRadius  gdc.MethodBindPtr
	fnGetInnerRadius  gdc.MethodBindPtr
	fnSetOuterRadius  gdc.MethodBindPtr
	fnGetOuterRadius  gdc.MethodBindPtr
	fnSetRings        gdc.MethodBindPtr
	fnGetRings        gdc.MethodBindPtr
	fnSetRingSegments gdc.MethodBindPtr
	fnGetRingSegments gdc.MethodBindPtr
}

var ptrsForTorusMesh ptrsForTorusMeshList

func initTorusMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TorusMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_inner_radius")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnSetInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_inner_radius")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnGetInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_outer_radius")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnSetOuterRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_outer_radius")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnGetOuterRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_rings")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnSetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_rings")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnGetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_ring_segments")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnSetRingSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_ring_segments")
		defer methodName.Destroy()
		ptrsForTorusMesh.fnGetRingSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type TorusMesh struct {
	PrimitiveMesh
}

func (me *TorusMesh) BaseClass() string {
	return "TorusMesh"
}

func NewTorusMesh() *TorusMesh {
	str := StringNameFromStr("TorusMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TorusMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TorusMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TorusMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TorusMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TorusMesh) SetInnerRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnSetInnerRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TorusMesh) GetInnerRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnGetInnerRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TorusMesh) SetOuterRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnSetOuterRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TorusMesh) GetOuterRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnGetOuterRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TorusMesh) SetRings(rings int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnSetRings), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TorusMesh) GetRings() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnGetRings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TorusMesh) SetRingSegments(rings int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnSetRingSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TorusMesh) GetRingSegments() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTorusMesh.fnGetRingSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
