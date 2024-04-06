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

type ptrsForRibbonTrailMeshList struct {
	fnSetSize            gdc.MethodBindPtr
	fnGetSize            gdc.MethodBindPtr
	fnSetSections        gdc.MethodBindPtr
	fnGetSections        gdc.MethodBindPtr
	fnSetSectionLength   gdc.MethodBindPtr
	fnGetSectionLength   gdc.MethodBindPtr
	fnSetSectionSegments gdc.MethodBindPtr
	fnGetSectionSegments gdc.MethodBindPtr
	fnSetCurve           gdc.MethodBindPtr
	fnGetCurve           gdc.MethodBindPtr
	fnSetShape           gdc.MethodBindPtr
	fnGetShape           gdc.MethodBindPtr
}

var ptrsForRibbonTrailMesh ptrsForRibbonTrailMeshList

func initRibbonTrailMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RibbonTrailMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sections")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnSetSections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_sections")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnGetSections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_section_length")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnSetSectionLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_section_length")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnGetSectionLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_section_segments")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnSetSectionSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_section_segments")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnGetSectionSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_curve")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnSetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_curve")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnGetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_shape")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1684440262))
	}
	{
		methodName := StringNameFromStr("get_shape")
		defer methodName.Destroy()
		ptrsForRibbonTrailMesh.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1317484155))
	}

}

type RibbonTrailMesh struct {
	PrimitiveMesh
}

func (me *RibbonTrailMesh) BaseClass() string {
	return "RibbonTrailMesh"
}

func NewRibbonTrailMesh() *RibbonTrailMesh {
	str := StringNameFromStr("RibbonTrailMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RibbonTrailMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type RibbonTrailMeshShape int

const (
	RibbonTrailMeshShapeShapeFlat  RibbonTrailMeshShape = 0
	RibbonTrailMeshShapeShapeCross RibbonTrailMeshShape = 1
)

func (me *RibbonTrailMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RibbonTrailMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RibbonTrailMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RibbonTrailMesh) SetSize(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RibbonTrailMesh) GetSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RibbonTrailMesh) SetSections(sections int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sections)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnSetSections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RibbonTrailMesh) GetSections() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnGetSections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RibbonTrailMesh) SetSectionLength(section_length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnSetSectionLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RibbonTrailMesh) GetSectionLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnGetSectionLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RibbonTrailMesh) SetSectionSegments(section_segments int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&section_segments)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnSetSectionSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RibbonTrailMesh) GetSectionSegments() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnGetSectionSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RibbonTrailMesh) SetCurve(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnSetCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RibbonTrailMesh) GetCurve() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnGetCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RibbonTrailMesh) SetShape(shape RibbonTrailMeshShape) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RibbonTrailMesh) GetShape() RibbonTrailMeshShape {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RibbonTrailMeshShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRibbonTrailMesh.fnGetShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
