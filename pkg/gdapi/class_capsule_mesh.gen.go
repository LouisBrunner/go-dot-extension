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

type ptrsForCapsuleMeshList struct {
	fnSetRadius         gdc.MethodBindPtr
	fnGetRadius         gdc.MethodBindPtr
	fnSetHeight         gdc.MethodBindPtr
	fnGetHeight         gdc.MethodBindPtr
	fnSetRadialSegments gdc.MethodBindPtr
	fnGetRadialSegments gdc.MethodBindPtr
	fnSetRings          gdc.MethodBindPtr
	fnGetRings          gdc.MethodBindPtr
}

var ptrsForCapsuleMesh ptrsForCapsuleMeshList

func initCapsuleMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CapsuleMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_radial_segments")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnSetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_radial_segments")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnGetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_rings")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnSetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_rings")
		defer methodName.Destroy()
		ptrsForCapsuleMesh.fnGetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type CapsuleMesh struct {
	PrimitiveMesh
}

func (me *CapsuleMesh) BaseClass() string {
	return "CapsuleMesh"
}

func NewCapsuleMesh() *CapsuleMesh {
	str := StringNameFromStr("CapsuleMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CapsuleMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CapsuleMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CapsuleMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CapsuleMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CapsuleMesh) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CapsuleMesh) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CapsuleMesh) SetHeight(height float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CapsuleMesh) GetHeight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CapsuleMesh) SetRadialSegments(segments int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnSetRadialSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CapsuleMesh) GetRadialSegments() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnGetRadialSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CapsuleMesh) SetRings(rings int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnSetRings), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CapsuleMesh) GetRings() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleMesh.fnGetRings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
