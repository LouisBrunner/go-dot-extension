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

type ptrsForCSGTorus3DList struct {
	fnSetInnerRadius gdc.MethodBindPtr
	fnGetInnerRadius gdc.MethodBindPtr
	fnSetOuterRadius gdc.MethodBindPtr
	fnGetOuterRadius gdc.MethodBindPtr
	fnSetSides       gdc.MethodBindPtr
	fnGetSides       gdc.MethodBindPtr
	fnSetRingSides   gdc.MethodBindPtr
	fnGetRingSides   gdc.MethodBindPtr
	fnSetMaterial    gdc.MethodBindPtr
	fnGetMaterial    gdc.MethodBindPtr
	fnSetSmoothFaces gdc.MethodBindPtr
	fnGetSmoothFaces gdc.MethodBindPtr
}

var ptrsForCSGTorus3D ptrsForCSGTorus3DList

func initCSGTorus3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CSGTorus3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_inner_radius")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnSetInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_inner_radius")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnGetInnerRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_outer_radius")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnSetOuterRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_outer_radius")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnGetOuterRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sides")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnSetSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sides")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnGetSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_ring_sides")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnSetRingSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_ring_sides")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnGetRingSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("set_smooth_faces")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnSetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_smooth_faces")
		defer methodName.Destroy()
		ptrsForCSGTorus3D.fnGetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type CSGTorus3D struct {
	CSGPrimitive3D
}

func (me *CSGTorus3D) BaseClass() string {
	return "CSGTorus3D"
}

func NewCSGTorus3D() *CSGTorus3D {
	str := StringNameFromStr("CSGTorus3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CSGTorus3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CSGTorus3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CSGTorus3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CSGTorus3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CSGTorus3D) SetInnerRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnSetInnerRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGTorus3D) GetInnerRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnGetInnerRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGTorus3D) SetOuterRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnSetOuterRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGTorus3D) GetOuterRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnGetOuterRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGTorus3D) SetSides(sides int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sides)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnSetSides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGTorus3D) GetSides() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnGetSides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGTorus3D) SetRingSides(sides int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sides)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnSetRingSides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGTorus3D) GetRingSides() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnGetRingSides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGTorus3D) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGTorus3D) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CSGTorus3D) SetSmoothFaces(smooth_faces bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth_faces)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnSetSmoothFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGTorus3D) GetSmoothFaces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGTorus3D.fnGetSmoothFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
