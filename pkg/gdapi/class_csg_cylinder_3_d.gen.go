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

type ptrsForCSGCylinder3DList struct {
	fnSetRadius      gdc.MethodBindPtr
	fnGetRadius      gdc.MethodBindPtr
	fnSetHeight      gdc.MethodBindPtr
	fnGetHeight      gdc.MethodBindPtr
	fnSetSides       gdc.MethodBindPtr
	fnGetSides       gdc.MethodBindPtr
	fnSetCone        gdc.MethodBindPtr
	fnIsCone         gdc.MethodBindPtr
	fnSetMaterial    gdc.MethodBindPtr
	fnGetMaterial    gdc.MethodBindPtr
	fnSetSmoothFaces gdc.MethodBindPtr
	fnGetSmoothFaces gdc.MethodBindPtr
}

var ptrsForCSGCylinder3D ptrsForCSGCylinder3DList

func initCSGCylinder3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CSGCylinder3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sides")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnSetSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sides")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnGetSides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_cone")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnSetCone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_cone")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnIsCone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("set_smooth_faces")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnSetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_smooth_faces")
		defer methodName.Destroy()
		ptrsForCSGCylinder3D.fnGetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type CSGCylinder3D struct {
	CSGPrimitive3D
}

func (me *CSGCylinder3D) BaseClass() string {
	return "CSGCylinder3D"
}

func NewCSGCylinder3D() *CSGCylinder3D {
	str := StringNameFromStr("CSGCylinder3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CSGCylinder3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CSGCylinder3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CSGCylinder3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CSGCylinder3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CSGCylinder3D) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGCylinder3D) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGCylinder3D) SetHeight(height float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGCylinder3D) GetHeight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGCylinder3D) SetSides(sides int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sides)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnSetSides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGCylinder3D) GetSides() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnGetSides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGCylinder3D) SetCone(cone bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cone)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnSetCone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGCylinder3D) IsCone() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnIsCone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CSGCylinder3D) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGCylinder3D) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CSGCylinder3D) SetSmoothFaces(smooth_faces bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth_faces)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnSetSmoothFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGCylinder3D) GetSmoothFaces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGCylinder3D.fnGetSmoothFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
