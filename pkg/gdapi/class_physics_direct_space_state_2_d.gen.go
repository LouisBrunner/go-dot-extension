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

type ptrsForPhysicsDirectSpaceState2DList struct {
	fnIntersectPoint gdc.MethodBindPtr
	fnIntersectRay   gdc.MethodBindPtr
	fnIntersectShape gdc.MethodBindPtr
	fnCastMotion     gdc.MethodBindPtr
	fnCollideShape   gdc.MethodBindPtr
	fnGetRestInfo    gdc.MethodBindPtr
}

var ptrsForPhysicsDirectSpaceState2D ptrsForPhysicsDirectSpaceState2DList

func initPhysicsDirectSpaceState2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsDirectSpaceState2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("intersect_point")
		defer methodName.Destroy()
		ptrsForPhysicsDirectSpaceState2D.fnIntersectPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2118456068))
	}
	{
		methodName := StringNameFromStr("intersect_ray")
		defer methodName.Destroy()
		ptrsForPhysicsDirectSpaceState2D.fnIntersectRay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1590275562))
	}
	{
		methodName := StringNameFromStr("intersect_shape")
		defer methodName.Destroy()
		ptrsForPhysicsDirectSpaceState2D.fnIntersectShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2488867228))
	}
	{
		methodName := StringNameFromStr("cast_motion")
		defer methodName.Destroy()
		ptrsForPhysicsDirectSpaceState2D.fnCastMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711275086))
	}
	{
		methodName := StringNameFromStr("collide_shape")
		defer methodName.Destroy()
		ptrsForPhysicsDirectSpaceState2D.fnCollideShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2488867228))
	}
	{
		methodName := StringNameFromStr("get_rest_info")
		defer methodName.Destroy()
		ptrsForPhysicsDirectSpaceState2D.fnGetRestInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2803666496))
	}
}

type PhysicsDirectSpaceState2D struct {
	Object
}

func (me *PhysicsDirectSpaceState2D) BaseClass() string {
	return "PhysicsDirectSpaceState2D"
}

func NewPhysicsDirectSpaceState2D() *PhysicsDirectSpaceState2D {
	str := StringNameFromStr("PhysicsDirectSpaceState2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsDirectSpaceState2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsDirectSpaceState2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsDirectSpaceState2D) IntersectPoint(parameters PhysicsPointQueryParameters2D, max_results int64) []Dictionary {
	cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&max_results)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState2D.fnIntersectPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsDirectSpaceState2D) IntersectRay(parameters PhysicsRayQueryParameters2D) Dictionary {
	cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState2D.fnIntersectRay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectSpaceState2D) IntersectShape(parameters PhysicsShapeQueryParameters2D, max_results int64) []Dictionary {
	cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&max_results)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState2D.fnIntersectShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsDirectSpaceState2D) CastMotion(parameters PhysicsShapeQueryParameters2D) PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState2D.fnCastMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectSpaceState2D) CollideShape(parameters PhysicsShapeQueryParameters2D, max_results int64) []Vector2 {
	cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&max_results)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState2D.fnCollideShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Vector2](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsDirectSpaceState2D) GetRestInfo(parameters PhysicsShapeQueryParameters2D) Dictionary {
	cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState2D.fnGetRestInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
