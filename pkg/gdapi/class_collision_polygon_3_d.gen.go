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

type ptrsForCollisionPolygon3DList struct {
	fnSetDepth    gdc.MethodBindPtr
	fnGetDepth    gdc.MethodBindPtr
	fnSetPolygon  gdc.MethodBindPtr
	fnGetPolygon  gdc.MethodBindPtr
	fnSetDisabled gdc.MethodBindPtr
	fnIsDisabled  gdc.MethodBindPtr
	fnSetMargin   gdc.MethodBindPtr
	fnGetMargin   gdc.MethodBindPtr
}

var ptrsForCollisionPolygon3D ptrsForCollisionPolygon3DList

func initCollisionPolygon3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CollisionPolygon3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_depth")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnSetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_polygon")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnSetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_polygon")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_disabled")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnSetDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_disabled")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnIsDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_margin")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_margin")
		defer methodName.Destroy()
		ptrsForCollisionPolygon3D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type CollisionPolygon3D struct {
	Node3D
}

func (me *CollisionPolygon3D) BaseClass() string {
	return "CollisionPolygon3D"
}

func NewCollisionPolygon3D() *CollisionPolygon3D {
	str := StringNameFromStr("CollisionPolygon3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CollisionPolygon3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CollisionPolygon3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CollisionPolygon3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CollisionPolygon3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CollisionPolygon3D) SetDepth(depth float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnSetDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon3D) GetDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionPolygon3D) SetPolygon(polygon PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnSetPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon3D) GetPolygon() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionPolygon3D) SetDisabled(disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnSetDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon3D) IsDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnIsDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionPolygon3D) SetMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon3D) GetMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon3D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
