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

type ptrsForWorldBoundaryShape3DList struct {
	fnSetPlane gdc.MethodBindPtr
	fnGetPlane gdc.MethodBindPtr
}

var ptrsForWorldBoundaryShape3D ptrsForWorldBoundaryShape3DList

func initWorldBoundaryShape3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WorldBoundaryShape3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_plane")
		defer methodName.Destroy()
		ptrsForWorldBoundaryShape3D.fnSetPlane = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3505987427))
	}
	{
		methodName := StringNameFromStr("get_plane")
		defer methodName.Destroy()
		ptrsForWorldBoundaryShape3D.fnGetPlane = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2753500971))
	}
}

type WorldBoundaryShape3D struct {
	Shape3D
}

func (me *WorldBoundaryShape3D) BaseClass() string {
	return "WorldBoundaryShape3D"
}

func NewWorldBoundaryShape3D() *WorldBoundaryShape3D {
	str := StringNameFromStr("WorldBoundaryShape3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &WorldBoundaryShape3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *WorldBoundaryShape3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *WorldBoundaryShape3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *WorldBoundaryShape3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *WorldBoundaryShape3D) SetPlane(plane Plane) {
	cargs := []gdc.ConstTypePtr{plane.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldBoundaryShape3D.fnSetPlane), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WorldBoundaryShape3D) GetPlane() Plane {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPlane()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldBoundaryShape3D.fnGetPlane), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
