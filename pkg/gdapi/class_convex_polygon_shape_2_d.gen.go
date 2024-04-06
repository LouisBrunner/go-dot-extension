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

type ptrsForConvexPolygonShape2DList struct {
	fnSetPointCloud gdc.MethodBindPtr
	fnSetPoints     gdc.MethodBindPtr
	fnGetPoints     gdc.MethodBindPtr
}

var ptrsForConvexPolygonShape2D ptrsForConvexPolygonShape2DList

func initConvexPolygonShape2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ConvexPolygonShape2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_point_cloud")
		defer methodName.Destroy()
		ptrsForConvexPolygonShape2D.fnSetPointCloud = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("set_points")
		defer methodName.Destroy()
		ptrsForConvexPolygonShape2D.fnSetPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_points")
		defer methodName.Destroy()
		ptrsForConvexPolygonShape2D.fnGetPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}

}

type ConvexPolygonShape2D struct {
	Shape2D
}

func (me *ConvexPolygonShape2D) BaseClass() string {
	return "ConvexPolygonShape2D"
}

func NewConvexPolygonShape2D() *ConvexPolygonShape2D {
	str := StringNameFromStr("ConvexPolygonShape2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ConvexPolygonShape2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ConvexPolygonShape2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ConvexPolygonShape2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ConvexPolygonShape2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ConvexPolygonShape2D) SetPointCloud(point_cloud PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{point_cloud.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConvexPolygonShape2D.fnSetPointCloud), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConvexPolygonShape2D) SetPoints(points PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConvexPolygonShape2D.fnSetPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConvexPolygonShape2D) GetPoints() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConvexPolygonShape2D.fnGetPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
