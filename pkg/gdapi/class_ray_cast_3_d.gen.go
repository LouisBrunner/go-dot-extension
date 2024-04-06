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

type ptrsForRayCast3DList struct {
	fnSetEnabled                 gdc.MethodBindPtr
	fnIsEnabled                  gdc.MethodBindPtr
	fnSetTargetPosition          gdc.MethodBindPtr
	fnGetTargetPosition          gdc.MethodBindPtr
	fnIsColliding                gdc.MethodBindPtr
	fnForceRaycastUpdate         gdc.MethodBindPtr
	fnGetCollider                gdc.MethodBindPtr
	fnGetColliderRid             gdc.MethodBindPtr
	fnGetColliderShape           gdc.MethodBindPtr
	fnGetCollisionPoint          gdc.MethodBindPtr
	fnGetCollisionNormal         gdc.MethodBindPtr
	fnGetCollisionFaceIndex      gdc.MethodBindPtr
	fnAddExceptionRid            gdc.MethodBindPtr
	fnAddException               gdc.MethodBindPtr
	fnRemoveExceptionRid         gdc.MethodBindPtr
	fnRemoveException            gdc.MethodBindPtr
	fnClearExceptions            gdc.MethodBindPtr
	fnSetCollisionMask           gdc.MethodBindPtr
	fnGetCollisionMask           gdc.MethodBindPtr
	fnSetCollisionMaskValue      gdc.MethodBindPtr
	fnGetCollisionMaskValue      gdc.MethodBindPtr
	fnSetExcludeParentBody       gdc.MethodBindPtr
	fnGetExcludeParentBody       gdc.MethodBindPtr
	fnSetCollideWithAreas        gdc.MethodBindPtr
	fnIsCollideWithAreasEnabled  gdc.MethodBindPtr
	fnSetCollideWithBodies       gdc.MethodBindPtr
	fnIsCollideWithBodiesEnabled gdc.MethodBindPtr
	fnSetHitFromInside           gdc.MethodBindPtr
	fnIsHitFromInsideEnabled     gdc.MethodBindPtr
	fnSetHitBackFaces            gdc.MethodBindPtr
	fnIsHitBackFacesEnabled      gdc.MethodBindPtr
	fnSetDebugShapeCustomColor   gdc.MethodBindPtr
	fnGetDebugShapeCustomColor   gdc.MethodBindPtr
	fnSetDebugShapeThickness     gdc.MethodBindPtr
	fnGetDebugShapeThickness     gdc.MethodBindPtr
}

var ptrsForRayCast3D ptrsForRayCast3DList

func initRayCast3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RayCast3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_enabled")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_target_position")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_target_position")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("is_colliding")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnIsColliding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("force_raycast_update")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnForceRaycastUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_collider")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1981248198))
	}
	{
		methodName := StringNameFromStr("get_collider_rid")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_collider_shape")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_collision_point")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetCollisionPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_collision_normal")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetCollisionNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_collision_face_index")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetCollisionFaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_exception_rid")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnAddExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("add_exception")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnAddException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1976431078))
	}
	{
		methodName := StringNameFromStr("remove_exception_rid")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnRemoveExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("remove_exception")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnRemoveException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1976431078))
	}
	{
		methodName := StringNameFromStr("clear_exceptions")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnClearExceptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_mask_value")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_mask_value")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_exclude_parent_body")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_exclude_parent_body")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_with_areas")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_areas_enabled")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_with_bodies")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_bodies_enabled")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hit_from_inside")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetHitFromInside = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hit_from_inside_enabled")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnIsHitFromInsideEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hit_back_faces")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetHitBackFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hit_back_faces_enabled")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnIsHitBackFacesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_debug_shape_custom_color")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetDebugShapeCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_debug_shape_custom_color")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetDebugShapeCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_debug_shape_thickness")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnSetDebugShapeThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_debug_shape_thickness")
		defer methodName.Destroy()
		ptrsForRayCast3D.fnGetDebugShapeThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type RayCast3D struct {
	Node3D
}

func (me *RayCast3D) BaseClass() string {
	return "RayCast3D"
}

func NewRayCast3D() *RayCast3D {
	str := StringNameFromStr("RayCast3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RayCast3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RayCast3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RayCast3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RayCast3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RayCast3D) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) IsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetTargetPosition(local_point Vector3) {
	cargs := []gdc.ConstTypePtr{local_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) GetTargetPosition() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RayCast3D) IsColliding() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnIsColliding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) ForceRaycastUpdate() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnForceRaycastUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) GetCollider() Object {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RayCast3D) GetColliderRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RayCast3D) GetColliderShape() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) GetCollisionPoint() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetCollisionPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RayCast3D) GetCollisionNormal() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetCollisionNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RayCast3D) GetCollisionFaceIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetCollisionFaceIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) AddExceptionRid(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnAddExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) AddException(node CollisionObject3D) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnAddException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) RemoveExceptionRid(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnRemoveExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) RemoveException(node CollisionObject3D) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnRemoveException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) ClearExceptions() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnClearExceptions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) SetCollisionMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetCollisionMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) GetCollisionMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetExcludeParentBody(mask bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetExcludeParentBody), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) GetExcludeParentBody() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetExcludeParentBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetCollideWithAreas(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) IsCollideWithAreasEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetCollideWithBodies(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) IsCollideWithBodiesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetHitFromInside(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetHitFromInside), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) IsHitFromInsideEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnIsHitFromInsideEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetHitBackFaces(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetHitBackFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) IsHitBackFacesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnIsHitBackFacesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RayCast3D) SetDebugShapeCustomColor(debug_shape_custom_color Color) {
	cargs := []gdc.ConstTypePtr{debug_shape_custom_color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetDebugShapeCustomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) GetDebugShapeCustomColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetDebugShapeCustomColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RayCast3D) SetDebugShapeThickness(debug_shape_thickness int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&debug_shape_thickness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnSetDebugShapeThickness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RayCast3D) GetDebugShapeThickness() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast3D.fnGetDebugShapeThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
