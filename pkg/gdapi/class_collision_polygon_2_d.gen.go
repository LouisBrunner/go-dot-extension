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

type ptrsForCollisionPolygon2DList struct {
	fnSetPolygon               gdc.MethodBindPtr
	fnGetPolygon               gdc.MethodBindPtr
	fnSetBuildMode             gdc.MethodBindPtr
	fnGetBuildMode             gdc.MethodBindPtr
	fnSetDisabled              gdc.MethodBindPtr
	fnIsDisabled               gdc.MethodBindPtr
	fnSetOneWayCollision       gdc.MethodBindPtr
	fnIsOneWayCollisionEnabled gdc.MethodBindPtr
	fnSetOneWayCollisionMargin gdc.MethodBindPtr
	fnGetOneWayCollisionMargin gdc.MethodBindPtr
}

var ptrsForCollisionPolygon2D ptrsForCollisionPolygon2DList

func initCollisionPolygon2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CollisionPolygon2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_polygon")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnSetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_polygon")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_build_mode")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnSetBuildMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2780803135))
	}
	{
		methodName := StringNameFromStr("get_build_mode")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnGetBuildMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3044948800))
	}
	{
		methodName := StringNameFromStr("set_disabled")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnSetDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_disabled")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnIsDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_one_way_collision")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnSetOneWayCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_one_way_collision_enabled")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnIsOneWayCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_one_way_collision_margin")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnSetOneWayCollisionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_one_way_collision_margin")
		defer methodName.Destroy()
		ptrsForCollisionPolygon2D.fnGetOneWayCollisionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type CollisionPolygon2D struct {
	Node2D
}

func (me *CollisionPolygon2D) BaseClass() string {
	return "CollisionPolygon2D"
}

func NewCollisionPolygon2D() *CollisionPolygon2D {
	str := StringNameFromStr("CollisionPolygon2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CollisionPolygon2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CollisionPolygon2DBuildMode int

const (
	CollisionPolygon2DBuildModeBuildSolids   CollisionPolygon2DBuildMode = 0
	CollisionPolygon2DBuildModeBuildSegments CollisionPolygon2DBuildMode = 1
)

func (me *CollisionPolygon2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CollisionPolygon2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CollisionPolygon2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CollisionPolygon2D) SetPolygon(polygon PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnSetPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon2D) GetPolygon() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionPolygon2D) SetBuildMode(build_mode CollisionPolygon2DBuildMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&build_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnSetBuildMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon2D) GetBuildMode() CollisionPolygon2DBuildMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CollisionPolygon2DBuildMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnGetBuildMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CollisionPolygon2D) SetDisabled(disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnSetDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon2D) IsDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnIsDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionPolygon2D) SetOneWayCollision(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnSetOneWayCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon2D) IsOneWayCollisionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnIsOneWayCollisionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionPolygon2D) SetOneWayCollisionMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnSetOneWayCollisionMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionPolygon2D) GetOneWayCollisionMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionPolygon2D.fnGetOneWayCollisionMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
