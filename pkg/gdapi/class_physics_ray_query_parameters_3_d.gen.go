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

type ptrsForPhysicsRayQueryParameters3DList struct {
	fnCreate                     gdc.MethodBindPtr
	fnSetFrom                    gdc.MethodBindPtr
	fnGetFrom                    gdc.MethodBindPtr
	fnSetTo                      gdc.MethodBindPtr
	fnGetTo                      gdc.MethodBindPtr
	fnSetCollisionMask           gdc.MethodBindPtr
	fnGetCollisionMask           gdc.MethodBindPtr
	fnSetExclude                 gdc.MethodBindPtr
	fnGetExclude                 gdc.MethodBindPtr
	fnSetCollideWithBodies       gdc.MethodBindPtr
	fnIsCollideWithBodiesEnabled gdc.MethodBindPtr
	fnSetCollideWithAreas        gdc.MethodBindPtr
	fnIsCollideWithAreasEnabled  gdc.MethodBindPtr
	fnSetHitFromInside           gdc.MethodBindPtr
	fnIsHitFromInsideEnabled     gdc.MethodBindPtr
	fnSetHitBackFaces            gdc.MethodBindPtr
	fnIsHitBackFacesEnabled      gdc.MethodBindPtr
}

var ptrsForPhysicsRayQueryParameters3D ptrsForPhysicsRayQueryParameters3DList

func initPhysicsRayQueryParameters3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsRayQueryParameters3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3110599579))
	}
	{
		methodName := StringNameFromStr("set_from")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_from")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnGetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_to")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_to")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnGetTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnGetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_collide_with_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_bodies_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_with_areas")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_areas_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hit_from_inside")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetHitFromInside = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hit_from_inside_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnIsHitFromInsideEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hit_back_faces")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnSetHitBackFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hit_back_faces_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters3D.fnIsHitBackFacesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type PhysicsRayQueryParameters3D struct {
	RefCounted
}

func (me *PhysicsRayQueryParameters3D) BaseClass() string {
	return "PhysicsRayQueryParameters3D"
}

func NewPhysicsRayQueryParameters3D() *PhysicsRayQueryParameters3D {
	str := StringNameFromStr("PhysicsRayQueryParameters3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsRayQueryParameters3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsRayQueryParameters3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsRayQueryParameters3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsRayQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func PhysicsRayQueryParameters3DCreate(from Vector3, to Vector3, collision_mask int64, exclude []RID) PhysicsRayQueryParameters3D {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&collision_mask), gdc.ConstTypePtr(&exclude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsRayQueryParameters3D()
	pinner.Pin(&collision_mask)
	pinner.Pin(&exclude)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnCreate), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsRayQueryParameters3D) SetFrom(from Vector3) {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) GetFrom() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnGetFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsRayQueryParameters3D) SetTo(to Vector3) {
	cargs := []gdc.ConstTypePtr{to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetTo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) GetTo() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnGetTo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsRayQueryParameters3D) SetCollisionMask(collision_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsRayQueryParameters3D) SetExclude(exclude []RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetExclude), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) GetExclude() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnGetExclude), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsRayQueryParameters3D) SetCollideWithBodies(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) IsCollideWithBodiesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsRayQueryParameters3D) SetCollideWithAreas(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) IsCollideWithAreasEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsRayQueryParameters3D) SetHitFromInside(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetHitFromInside), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) IsHitFromInsideEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnIsHitFromInsideEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsRayQueryParameters3D) SetHitBackFaces(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnSetHitBackFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters3D) IsHitBackFacesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters3D.fnIsHitBackFacesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
