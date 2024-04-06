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

type ptrsForPhysicsPointQueryParameters3DList struct {
	fnSetPosition                gdc.MethodBindPtr
	fnGetPosition                gdc.MethodBindPtr
	fnSetCollisionMask           gdc.MethodBindPtr
	fnGetCollisionMask           gdc.MethodBindPtr
	fnSetExclude                 gdc.MethodBindPtr
	fnGetExclude                 gdc.MethodBindPtr
	fnSetCollideWithBodies       gdc.MethodBindPtr
	fnIsCollideWithBodiesEnabled gdc.MethodBindPtr
	fnSetCollideWithAreas        gdc.MethodBindPtr
	fnIsCollideWithAreasEnabled  gdc.MethodBindPtr
}

var ptrsForPhysicsPointQueryParameters3D ptrsForPhysicsPointQueryParameters3DList

func initPhysicsPointQueryParameters3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsPointQueryParameters3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnSetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnGetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_collide_with_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_bodies_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_with_areas")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_areas_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsPointQueryParameters3D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type PhysicsPointQueryParameters3D struct {
	RefCounted
}

func (me *PhysicsPointQueryParameters3D) BaseClass() string {
	return "PhysicsPointQueryParameters3D"
}

func NewPhysicsPointQueryParameters3D() *PhysicsPointQueryParameters3D {
	str := StringNameFromStr("PhysicsPointQueryParameters3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsPointQueryParameters3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsPointQueryParameters3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsPointQueryParameters3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsPointQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsPointQueryParameters3D) SetPosition(position Vector3) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters3D) GetPosition() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsPointQueryParameters3D) SetCollisionMask(collision_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters3D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsPointQueryParameters3D) SetExclude(exclude []RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnSetExclude), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters3D) GetExclude() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnGetExclude), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsPointQueryParameters3D) SetCollideWithBodies(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters3D) IsCollideWithBodiesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsPointQueryParameters3D) SetCollideWithAreas(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsPointQueryParameters3D) IsCollideWithAreasEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsPointQueryParameters3D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
