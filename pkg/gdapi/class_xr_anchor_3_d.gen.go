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

type ptrsForXRAnchor3DList struct {
	fnGetSize  gdc.MethodBindPtr
	fnGetPlane gdc.MethodBindPtr
}

var ptrsForXRAnchor3D ptrsForXRAnchor3DList

func initXRAnchor3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRAnchor3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForXRAnchor3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_plane")
		defer methodName.Destroy()
		ptrsForXRAnchor3D.fnGetPlane = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2753500971))
	}
}

type XRAnchor3D struct {
	XRNode3D
}

func (me *XRAnchor3D) BaseClass() string {
	return "XRAnchor3D"
}

func NewXRAnchor3D() *XRAnchor3D {
	str := StringNameFromStr("XRAnchor3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRAnchor3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *XRAnchor3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRAnchor3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRAnchor3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRAnchor3D) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRAnchor3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRAnchor3D) GetPlane() Plane {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPlane()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRAnchor3D.fnGetPlane), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
