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

type ptrsForCollisionObject3DList struct {
	fnXInputEvent             gdc.MethodBindPtr
	fnXMouseEnter             gdc.MethodBindPtr
	fnXMouseExit              gdc.MethodBindPtr
	fnSetCollisionLayer       gdc.MethodBindPtr
	fnGetCollisionLayer       gdc.MethodBindPtr
	fnSetCollisionMask        gdc.MethodBindPtr
	fnGetCollisionMask        gdc.MethodBindPtr
	fnSetCollisionLayerValue  gdc.MethodBindPtr
	fnGetCollisionLayerValue  gdc.MethodBindPtr
	fnSetCollisionMaskValue   gdc.MethodBindPtr
	fnGetCollisionMaskValue   gdc.MethodBindPtr
	fnSetCollisionPriority    gdc.MethodBindPtr
	fnGetCollisionPriority    gdc.MethodBindPtr
	fnSetDisableMode          gdc.MethodBindPtr
	fnGetDisableMode          gdc.MethodBindPtr
	fnSetRayPickable          gdc.MethodBindPtr
	fnIsRayPickable           gdc.MethodBindPtr
	fnSetCaptureInputOnDrag   gdc.MethodBindPtr
	fnGetCaptureInputOnDrag   gdc.MethodBindPtr
	fnGetRid                  gdc.MethodBindPtr
	fnCreateShapeOwner        gdc.MethodBindPtr
	fnRemoveShapeOwner        gdc.MethodBindPtr
	fnGetShapeOwners          gdc.MethodBindPtr
	fnShapeOwnerSetTransform  gdc.MethodBindPtr
	fnShapeOwnerGetTransform  gdc.MethodBindPtr
	fnShapeOwnerGetOwner      gdc.MethodBindPtr
	fnShapeOwnerSetDisabled   gdc.MethodBindPtr
	fnIsShapeOwnerDisabled    gdc.MethodBindPtr
	fnShapeOwnerAddShape      gdc.MethodBindPtr
	fnShapeOwnerGetShapeCount gdc.MethodBindPtr
	fnShapeOwnerGetShape      gdc.MethodBindPtr
	fnShapeOwnerGetShapeIndex gdc.MethodBindPtr
	fnShapeOwnerRemoveShape   gdc.MethodBindPtr
	fnShapeOwnerClearShapes   gdc.MethodBindPtr
	fnShapeFindOwner          gdc.MethodBindPtr
}

var ptrsForCollisionObject3D ptrsForCollisionObject3DList

func initCollisionObject3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CollisionObject3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_collision_layer")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_layer")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_layer_value")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_layer_value")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_collision_mask_value")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_mask_value")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_collision_priority")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_collision_priority")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_disable_mode")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetDisableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1623620376))
	}
	{
		methodName := StringNameFromStr("get_disable_mode")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetDisableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 410164780))
	}
	{
		methodName := StringNameFromStr("set_ray_pickable")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetRayPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ray_pickable")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnIsRayPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_capture_input_on_drag")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnSetCaptureInputOnDrag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_capture_input_on_drag")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetCaptureInputOnDrag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("create_shape_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnCreateShapeOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3429307534))
	}
	{
		methodName := StringNameFromStr("remove_shape_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnRemoveShapeOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_shape_owners")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnGetShapeOwners = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("shape_owner_set_transform")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_transform")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerGetOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3332903315))
	}
	{
		methodName := StringNameFromStr("shape_owner_set_disabled")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerSetDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_shape_owner_disabled")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnIsShapeOwnerDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("shape_owner_add_shape")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerAddShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2566676345))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_shape_count")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerGetShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_shape")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015519174))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_shape_index")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerGetShapeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("shape_owner_remove_shape")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerRemoveShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("shape_owner_clear_shapes")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeOwnerClearShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("shape_find_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject3D.fnShapeFindOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
}

type CollisionObject3D struct {
	Node3D
}

func (me *CollisionObject3D) BaseClass() string {
	return "CollisionObject3D"
}

func NewCollisionObject3D() *CollisionObject3D {
	str := StringNameFromStr("CollisionObject3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CollisionObject3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CollisionObject3DDisableMode int

const (
	CollisionObject3DDisableModeDisableModeRemove     CollisionObject3DDisableMode = 0
	CollisionObject3DDisableModeDisableModeMakeStatic CollisionObject3DDisableMode = 1
	CollisionObject3DDisableModeDisableModeKeepActive CollisionObject3DDisableMode = 2
)

func (me *CollisionObject3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CollisionObject3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CollisionObject3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CollisionObject3D) SetCollisionLayer(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetCollisionLayer() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) SetCollisionMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) SetCollisionLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetCollisionLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) SetCollisionMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetCollisionMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) SetCollisionPriority(priority float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetCollisionPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetCollisionPriority() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetCollisionPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) SetDisableMode(mode CollisionObject3DDisableMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetDisableMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetDisableMode() CollisionObject3DDisableMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CollisionObject3DDisableMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetDisableMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CollisionObject3D) SetRayPickable(ray_pickable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ray_pickable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetRayPickable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) IsRayPickable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnIsRayPickable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) SetCaptureInputOnDrag(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnSetCaptureInputOnDrag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetCaptureInputOnDrag() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetCaptureInputOnDrag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject3D) CreateShapeOwner(owner Object) int64 {
	cargs := []gdc.ConstTypePtr{owner.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnCreateShapeOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) RemoveShapeOwner(owner_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnRemoveShapeOwner), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) GetShapeOwners() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnGetShapeOwners), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject3D) ShapeOwnerSetTransform(owner_id int64, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) ShapeOwnerGetTransform(owner_id int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject3D) ShapeOwnerGetOwner(owner_id int64) Object {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerGetOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject3D) ShapeOwnerSetDisabled(owner_id int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerSetDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) IsShapeOwnerDisabled(owner_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnIsShapeOwnerDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) ShapeOwnerAddShape(owner_id int64, shape Shape3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerAddShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) ShapeOwnerGetShapeCount(owner_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerGetShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) ShapeOwnerGetShape(owner_id int64, shape_id int64) Shape3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShape3D()
	pinner.Pin(&owner_id)
	pinner.Pin(&shape_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject3D) ShapeOwnerGetShapeIndex(owner_id int64, shape_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&owner_id)
	pinner.Pin(&shape_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerGetShapeIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject3D) ShapeOwnerRemoveShape(owner_id int64, shape_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerRemoveShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) ShapeOwnerClearShapes(owner_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeOwnerClearShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject3D) ShapeFindOwner(shape_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&shape_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject3D.fnShapeFindOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CollisionObject3DInputEventSignalFn func(camera Node, event InputEvent, position Vector3, normal Vector3, shape_idx int)

func (me *CollisionObject3D) ConnectInputEvent(subs SignalSubscribers, fn CollisionObject3DInputEventSignalFn) {
	sig := StringNameFromStr("input_event")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject3D) DisconnectInputEvent(subs SignalSubscribers, fn CollisionObject3DInputEventSignalFn) {
	sig := StringNameFromStr("input_event")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject3DMouseEnteredSignalFn func()

func (me *CollisionObject3D) ConnectMouseEntered(subs SignalSubscribers, fn CollisionObject3DMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject3D) DisconnectMouseEntered(subs SignalSubscribers, fn CollisionObject3DMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject3DMouseExitedSignalFn func()

func (me *CollisionObject3D) ConnectMouseExited(subs SignalSubscribers, fn CollisionObject3DMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject3D) DisconnectMouseExited(subs SignalSubscribers, fn CollisionObject3DMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
