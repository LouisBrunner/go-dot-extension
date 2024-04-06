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

type ptrsForPlaneMeshList struct {
	fnSetSize           gdc.MethodBindPtr
	fnGetSize           gdc.MethodBindPtr
	fnSetSubdivideWidth gdc.MethodBindPtr
	fnGetSubdivideWidth gdc.MethodBindPtr
	fnSetSubdivideDepth gdc.MethodBindPtr
	fnGetSubdivideDepth gdc.MethodBindPtr
	fnSetCenterOffset   gdc.MethodBindPtr
	fnGetCenterOffset   gdc.MethodBindPtr
	fnSetOrientation    gdc.MethodBindPtr
	fnGetOrientation    gdc.MethodBindPtr
}

var ptrsForPlaneMesh ptrsForPlaneMeshList

func initPlaneMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PlaneMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_subdivide_width")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnSetSubdivideWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_subdivide_width")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnGetSubdivideWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_subdivide_depth")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnSetSubdivideDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_subdivide_depth")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnGetSubdivideDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_center_offset")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnSetCenterOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_center_offset")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnGetCenterOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_orientation")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnSetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2751399687))
	}
	{
		methodName := StringNameFromStr("get_orientation")
		defer methodName.Destroy()
		ptrsForPlaneMesh.fnGetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227599250))
	}
}

type PlaneMesh struct {
	PrimitiveMesh
}

func (me *PlaneMesh) BaseClass() string {
	return "PlaneMesh"
}

func NewPlaneMesh() *PlaneMesh {
	str := StringNameFromStr("PlaneMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PlaneMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type PlaneMeshOrientation int

const (
	PlaneMeshOrientationFaceX PlaneMeshOrientation = 0
	PlaneMeshOrientationFaceY PlaneMeshOrientation = 1
	PlaneMeshOrientationFaceZ PlaneMeshOrientation = 2
)

func (me *PlaneMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PlaneMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PlaneMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PlaneMesh) SetSize(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PlaneMesh) GetSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PlaneMesh) SetSubdivideWidth(subdivide int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivide)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnSetSubdivideWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PlaneMesh) GetSubdivideWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnGetSubdivideWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PlaneMesh) SetSubdivideDepth(subdivide int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivide)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnSetSubdivideDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PlaneMesh) GetSubdivideDepth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnGetSubdivideDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PlaneMesh) SetCenterOffset(offset Vector3) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnSetCenterOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PlaneMesh) GetCenterOffset() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnGetCenterOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PlaneMesh) SetOrientation(orientation PlaneMeshOrientation) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnSetOrientation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PlaneMesh) GetOrientation() PlaneMeshOrientation {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PlaneMeshOrientation

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaneMesh.fnGetOrientation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
