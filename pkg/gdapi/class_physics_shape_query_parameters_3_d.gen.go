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

type ptrsForPhysicsShapeQueryParameters3DList struct {
	fnSetShape                   gdc.MethodBindPtr
	fnGetShape                   gdc.MethodBindPtr
	fnSetShapeRid                gdc.MethodBindPtr
	fnGetShapeRid                gdc.MethodBindPtr
	fnSetTransform               gdc.MethodBindPtr
	fnGetTransform               gdc.MethodBindPtr
	fnSetMotion                  gdc.MethodBindPtr
	fnGetMotion                  gdc.MethodBindPtr
	fnSetMargin                  gdc.MethodBindPtr
	fnGetMargin                  gdc.MethodBindPtr
	fnSetCollisionMask           gdc.MethodBindPtr
	fnGetCollisionMask           gdc.MethodBindPtr
	fnSetExclude                 gdc.MethodBindPtr
	fnGetExclude                 gdc.MethodBindPtr
	fnSetCollideWithBodies       gdc.MethodBindPtr
	fnIsCollideWithBodiesEnabled gdc.MethodBindPtr
	fnSetCollideWithAreas        gdc.MethodBindPtr
	fnIsCollideWithAreasEnabled  gdc.MethodBindPtr
}

var ptrsForPhysicsShapeQueryParameters3D ptrsForPhysicsShapeQueryParameters3DList

func initPhysicsShapeQueryParameters3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsShapeQueryParameters3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_shape")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 968641751))
	}
	{
		methodName := StringNameFromStr("get_shape")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
	}
	{
		methodName := StringNameFromStr("set_shape_rid")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetShapeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_shape_rid")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnGetShapeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_transform")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_transform")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("set_motion")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_motion")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnGetMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_margin")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_margin")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnGetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_collide_with_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_bodies_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_with_areas")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_areas_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsShapeQueryParameters3D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type PhysicsShapeQueryParameters3D struct {
	RefCounted
}

func (me *PhysicsShapeQueryParameters3D) BaseClass() string {
	return "PhysicsShapeQueryParameters3D"
}

func NewPhysicsShapeQueryParameters3D() *PhysicsShapeQueryParameters3D {
	str := StringNameFromStr("PhysicsShapeQueryParameters3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsShapeQueryParameters3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsShapeQueryParameters3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsShapeQueryParameters3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsShapeQueryParameters3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsShapeQueryParameters3D) SetShape(shape Resource) {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) GetShape() Resource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsShapeQueryParameters3D) SetShapeRid(shape RID) {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetShapeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) GetShapeRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnGetShapeRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsShapeQueryParameters3D) SetTransform(transform Transform3D) {
	cargs := []gdc.ConstTypePtr{transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) GetTransform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsShapeQueryParameters3D) SetMotion(motion Vector3) {
	cargs := []gdc.ConstTypePtr{motion.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetMotion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) GetMotion() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnGetMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsShapeQueryParameters3D) SetMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) GetMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsShapeQueryParameters3D) SetCollisionMask(collision_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsShapeQueryParameters3D) SetExclude(exclude []RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetExclude), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) GetExclude() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnGetExclude), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsShapeQueryParameters3D) SetCollideWithBodies(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) IsCollideWithBodiesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsShapeQueryParameters3D) SetCollideWithAreas(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsShapeQueryParameters3D) IsCollideWithAreasEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsShapeQueryParameters3D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
