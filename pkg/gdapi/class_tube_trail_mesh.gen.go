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

type ptrsForTubeTrailMeshList struct {
	fnSetRadius        gdc.MethodBindPtr
	fnGetRadius        gdc.MethodBindPtr
	fnSetRadialSteps   gdc.MethodBindPtr
	fnGetRadialSteps   gdc.MethodBindPtr
	fnSetSections      gdc.MethodBindPtr
	fnGetSections      gdc.MethodBindPtr
	fnSetSectionLength gdc.MethodBindPtr
	fnGetSectionLength gdc.MethodBindPtr
	fnSetSectionRings  gdc.MethodBindPtr
	fnGetSectionRings  gdc.MethodBindPtr
	fnSetCapTop        gdc.MethodBindPtr
	fnIsCapTop         gdc.MethodBindPtr
	fnSetCapBottom     gdc.MethodBindPtr
	fnIsCapBottom      gdc.MethodBindPtr
	fnSetCurve         gdc.MethodBindPtr
	fnGetCurve         gdc.MethodBindPtr
}

var ptrsForTubeTrailMesh ptrsForTubeTrailMeshList

func initTubeTrailMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TubeTrailMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_radial_steps")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetRadialSteps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_radial_steps")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnGetRadialSteps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_sections")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetSections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sections")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnGetSections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_section_length")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetSectionLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_section_length")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnGetSectionLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_section_rings")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetSectionRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_section_rings")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnGetSectionRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_cap_top")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetCapTop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_cap_top")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnIsCapTop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_cap_bottom")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetCapBottom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_cap_bottom")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnIsCapBottom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_curve")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnSetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_curve")
		defer methodName.Destroy()
		ptrsForTubeTrailMesh.fnGetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
}

type TubeTrailMesh struct {
	PrimitiveMesh
}

func (me *TubeTrailMesh) BaseClass() string {
	return "TubeTrailMesh"
}

func NewTubeTrailMesh() *TubeTrailMesh {
	str := StringNameFromStr("TubeTrailMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TubeTrailMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *TubeTrailMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TubeTrailMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TubeTrailMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TubeTrailMesh) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TubeTrailMesh) SetRadialSteps(radial_steps int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radial_steps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetRadialSteps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) GetRadialSteps() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnGetRadialSteps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TubeTrailMesh) SetSections(sections int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sections)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetSections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) GetSections() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnGetSections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TubeTrailMesh) SetSectionLength(section_length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetSectionLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) GetSectionLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnGetSectionLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TubeTrailMesh) SetSectionRings(section_rings int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_rings)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetSectionRings), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) GetSectionRings() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnGetSectionRings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TubeTrailMesh) SetCapTop(cap_top bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_top)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetCapTop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) IsCapTop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnIsCapTop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TubeTrailMesh) SetCapBottom(cap_bottom bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_bottom)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetCapBottom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) IsCapBottom() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnIsCapBottom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TubeTrailMesh) SetCurve(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnSetCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TubeTrailMesh) GetCurve() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTubeTrailMesh.fnGetCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
