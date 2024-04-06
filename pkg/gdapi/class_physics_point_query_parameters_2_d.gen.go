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

type ptrsForPhysicsPointQueryParameters2DList struct {
	fnSetPosition                gdc.MethodBindPtr
	fnGetPosition                gdc.MethodBindPtr
	fnSetCanvasInstanceId        gdc.MethodBindPtr
	fnGetCanvasInstanceId        gdc.MethodBindPtr
	fnSetCollisionMask           gdc.MethodBindPtr
	fnGetCollisionMask           gdc.MethodBindPtr
	fnSetExclude                 gdc.MethodBindPtr
	fnGetExclude                 gdc.MethodBindPtr
	fnSetCollideWithBodies       gdc.MethodBindPtr
	fnIsCollideWithBodiesEnabled gdc.MethodBindPtr
	fnSetCollideWithAreas        gdc.MethodBindPtr
	fnIsCollideWithAreasEnabled  gdc.MethodBindPtr
}

var ptrsForPhysicsPointQueryParameters2D ptrsForPhysicsPointQueryParameters2DList

func initPhysicsPointQueryParameters2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsPointQueryParameters2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_canvas_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnSetCanvasInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_canvas_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnGetCanvasInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnSetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnGetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_collide_with_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_bodies_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_with_areas")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_areas_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters2D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type PhysicsPointQueryParameters2D struct {
	RefCounted
}

func (me *PhysicsPointQueryParameters2D) BaseClass() string {
	return "PhysicsPointQueryParameters2D"
}

func NewPhysicsPointQueryParameters2D() *PhysicsPointQueryParameters2D {
	str := StringNameFromStr("PhysicsPointQueryParameters2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsPointQueryParameters2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsPointQueryParameters2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsPointQueryParameters2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsPointQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsPointQueryParameters2D) SetPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters2D) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsPointQueryParameters2D) SetCanvasInstanceId(canvas_instance_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&canvas_instance_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnSetCanvasInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters2D) GetCanvasInstanceId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnGetCanvasInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsPointQueryParameters2D) SetCollisionMask(collision_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters2D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsPointQueryParameters2D) SetExclude(exclude []RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnSetExclude), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters2D) GetExclude() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnGetExclude), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsPointQueryParameters2D) SetCollideWithBodies(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters2D) IsCollideWithBodiesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsPointQueryParameters2D) SetCollideWithAreas(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters2D) IsCollideWithAreasEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters2D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
