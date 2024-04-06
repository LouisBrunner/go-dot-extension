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

type ptrsForConvexPolygonShape3DList struct {
	fnSetPoints gdc.MethodBindPtr
	fnGetPoints gdc.MethodBindPtr
}

var ptrsForConvexPolygonShape3D ptrsForConvexPolygonShape3DList

func initConvexPolygonShape3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ConvexPolygonShape3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_points")
		defer methodName.Destroy()
		ptrsForConvexPolygonShape3D.fnSetPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
	}
	{
		methodName := StringNameFromStr("get_points")
		defer methodName.Destroy()
		ptrsForConvexPolygonShape3D.fnGetPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
	}
}

type ConvexPolygonShape3D struct {
	Shape3D
}

func (me *ConvexPolygonShape3D) BaseClass() string {
	return "ConvexPolygonShape3D"
}

func NewConvexPolygonShape3D() *ConvexPolygonShape3D {
	str := StringNameFromStr("ConvexPolygonShape3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ConvexPolygonShape3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ConvexPolygonShape3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ConvexPolygonShape3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ConvexPolygonShape3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ConvexPolygonShape3D) SetPoints(points PackedVector3Array) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConvexPolygonShape3D.fnSetPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConvexPolygonShape3D) GetPoints() PackedVector3Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConvexPolygonShape3D.fnGetPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
