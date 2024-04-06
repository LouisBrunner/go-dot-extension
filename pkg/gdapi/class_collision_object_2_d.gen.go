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

type ptrsForCollisionObject2DList struct {
	fnXInputEvent                        gdc.MethodBindPtr
	fnXMouseEnter                        gdc.MethodBindPtr
	fnXMouseExit                         gdc.MethodBindPtr
	fnXMouseShapeEnter                   gdc.MethodBindPtr
	fnXMouseShapeExit                    gdc.MethodBindPtr
	fnGetRid                             gdc.MethodBindPtr
	fnSetCollisionLayer                  gdc.MethodBindPtr
	fnGetCollisionLayer                  gdc.MethodBindPtr
	fnSetCollisionMask                   gdc.MethodBindPtr
	fnGetCollisionMask                   gdc.MethodBindPtr
	fnSetCollisionLayerValue             gdc.MethodBindPtr
	fnGetCollisionLayerValue             gdc.MethodBindPtr
	fnSetCollisionMaskValue              gdc.MethodBindPtr
	fnGetCollisionMaskValue              gdc.MethodBindPtr
	fnSetCollisionPriority               gdc.MethodBindPtr
	fnGetCollisionPriority               gdc.MethodBindPtr
	fnSetDisableMode                     gdc.MethodBindPtr
	fnGetDisableMode                     gdc.MethodBindPtr
	fnSetPickable                        gdc.MethodBindPtr
	fnIsPickable                         gdc.MethodBindPtr
	fnCreateShapeOwner                   gdc.MethodBindPtr
	fnRemoveShapeOwner                   gdc.MethodBindPtr
	fnGetShapeOwners                     gdc.MethodBindPtr
	fnShapeOwnerSetTransform             gdc.MethodBindPtr
	fnShapeOwnerGetTransform             gdc.MethodBindPtr
	fnShapeOwnerGetOwner                 gdc.MethodBindPtr
	fnShapeOwnerSetDisabled              gdc.MethodBindPtr
	fnIsShapeOwnerDisabled               gdc.MethodBindPtr
	fnShapeOwnerSetOneWayCollision       gdc.MethodBindPtr
	fnIsShapeOwnerOneWayCollisionEnabled gdc.MethodBindPtr
	fnShapeOwnerSetOneWayCollisionMargin gdc.MethodBindPtr
	fnGetShapeOwnerOneWayCollisionMargin gdc.MethodBindPtr
	fnShapeOwnerAddShape                 gdc.MethodBindPtr
	fnShapeOwnerGetShapeCount            gdc.MethodBindPtr
	fnShapeOwnerGetShape                 gdc.MethodBindPtr
	fnShapeOwnerGetShapeIndex            gdc.MethodBindPtr
	fnShapeOwnerRemoveShape              gdc.MethodBindPtr
	fnShapeOwnerClearShapes              gdc.MethodBindPtr
	fnShapeFindOwner                     gdc.MethodBindPtr
}

var ptrsForCollisionObject2D ptrsForCollisionObject2DList

func initCollisionObject2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CollisionObject2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_collision_layer")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnSetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_layer")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_collision_layer_value")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnSetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_layer_value")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetCollisionLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_collision_mask_value")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_collision_mask_value")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_collision_priority")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnSetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_collision_priority")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_disable_mode")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnSetDisableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1919204045))
	}
	{
		methodName := StringNameFromStr("get_disable_mode")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetDisableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3172846349))
	}
	{
		methodName := StringNameFromStr("set_pickable")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnSetPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_pickable")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnIsPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("create_shape_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnCreateShapeOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3429307534))
	}
	{
		methodName := StringNameFromStr("remove_shape_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnRemoveShapeOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_shape_owners")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetShapeOwners = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("shape_owner_set_transform")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 30160968))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_transform")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3836996910))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerGetOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3332903315))
	}
	{
		methodName := StringNameFromStr("shape_owner_set_disabled")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerSetDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_shape_owner_disabled")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnIsShapeOwnerDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("shape_owner_set_one_way_collision")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerSetOneWayCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_shape_owner_one_way_collision_enabled")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnIsShapeOwnerOneWayCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("shape_owner_set_one_way_collision_margin")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerSetOneWayCollisionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_shape_owner_one_way_collision_margin")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnGetShapeOwnerOneWayCollisionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("shape_owner_add_shape")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerAddShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2077425081))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_shape_count")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerGetShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_shape")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3106725749))
	}
	{
		methodName := StringNameFromStr("shape_owner_get_shape_index")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerGetShapeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("shape_owner_remove_shape")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerRemoveShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("shape_owner_clear_shapes")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeOwnerClearShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("shape_find_owner")
		defer methodName.Destroy()
		ptrsForCollisionObject2D.fnShapeFindOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}

}

type CollisionObject2D struct {
	Node2D
}

func (me *CollisionObject2D) BaseClass() string {
	return "CollisionObject2D"
}

func NewCollisionObject2D() *CollisionObject2D {
	str := StringNameFromStr("CollisionObject2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CollisionObject2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CollisionObject2DDisableMode int

const (
	CollisionObject2DDisableModeDisableModeRemove     CollisionObject2DDisableMode = 0
	CollisionObject2DDisableModeDisableModeMakeStatic CollisionObject2DDisableMode = 1
	CollisionObject2DDisableModeDisableModeKeepActive CollisionObject2DDisableMode = 2
)

func (me *CollisionObject2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CollisionObject2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CollisionObject2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CollisionObject2D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject2D) SetCollisionLayer(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnSetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetCollisionLayer() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) SetCollisionMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) SetCollisionLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnSetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetCollisionLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetCollisionLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) SetCollisionMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetCollisionMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) SetCollisionPriority(priority float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnSetCollisionPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetCollisionPriority() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetCollisionPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) SetDisableMode(mode CollisionObject2DDisableMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnSetDisableMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetDisableMode() CollisionObject2DDisableMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CollisionObject2DDisableMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetDisableMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CollisionObject2D) SetPickable(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnSetPickable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) IsPickable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnIsPickable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) CreateShapeOwner(owner Object) int64 {
	cargs := []gdc.ConstTypePtr{owner.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnCreateShapeOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) RemoveShapeOwner(owner_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnRemoveShapeOwner), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetShapeOwners() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetShapeOwners), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject2D) ShapeOwnerSetTransform(owner_id int64, transform Transform2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) ShapeOwnerGetTransform(owner_id int64) Transform2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject2D) ShapeOwnerGetOwner(owner_id int64) Object {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerGetOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject2D) ShapeOwnerSetDisabled(owner_id int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerSetDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) IsShapeOwnerDisabled(owner_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnIsShapeOwnerDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) ShapeOwnerSetOneWayCollision(owner_id int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerSetOneWayCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) IsShapeOwnerOneWayCollisionEnabled(owner_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnIsShapeOwnerOneWayCollisionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) ShapeOwnerSetOneWayCollisionMargin(owner_id int64, margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerSetOneWayCollisionMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) GetShapeOwnerOneWayCollisionMargin(owner_id int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnGetShapeOwnerOneWayCollisionMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) ShapeOwnerAddShape(owner_id int64, shape Shape2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerAddShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) ShapeOwnerGetShapeCount(owner_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&owner_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerGetShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) ShapeOwnerGetShape(owner_id int64, shape_id int64) Shape2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShape2D()
	pinner.Pin(&owner_id)
	pinner.Pin(&shape_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CollisionObject2D) ShapeOwnerGetShapeIndex(owner_id int64, shape_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&owner_id)
	pinner.Pin(&shape_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerGetShapeIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CollisionObject2D) ShapeOwnerRemoveShape(owner_id int64, shape_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id), gdc.ConstTypePtr(&shape_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerRemoveShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) ShapeOwnerClearShapes(owner_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeOwnerClearShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CollisionObject2D) ShapeFindOwner(shape_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&shape_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionObject2D.fnShapeFindOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CollisionObject2DInputEventSignalFn func(viewport Node, event InputEvent, shape_idx int)

func (me *CollisionObject2D) ConnectInputEvent(subs SignalSubscribers, fn CollisionObject2DInputEventSignalFn) {
	sig := StringNameFromStr("input_event")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectInputEvent(subs SignalSubscribers, fn CollisionObject2DInputEventSignalFn) {
	sig := StringNameFromStr("input_event")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseEnteredSignalFn func()

func (me *CollisionObject2D) ConnectMouseEntered(subs SignalSubscribers, fn CollisionObject2DMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseEntered(subs SignalSubscribers, fn CollisionObject2DMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseExitedSignalFn func()

func (me *CollisionObject2D) ConnectMouseExited(subs SignalSubscribers, fn CollisionObject2DMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseExited(subs SignalSubscribers, fn CollisionObject2DMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseShapeEnteredSignalFn func(shape_idx int)

func (me *CollisionObject2D) ConnectMouseShapeEntered(subs SignalSubscribers, fn CollisionObject2DMouseShapeEnteredSignalFn) {
	sig := StringNameFromStr("mouse_shape_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseShapeEntered(subs SignalSubscribers, fn CollisionObject2DMouseShapeEnteredSignalFn) {
	sig := StringNameFromStr("mouse_shape_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type CollisionObject2DMouseShapeExitedSignalFn func(shape_idx int)

func (me *CollisionObject2D) ConnectMouseShapeExited(subs SignalSubscribers, fn CollisionObject2DMouseShapeExitedSignalFn) {
	sig := StringNameFromStr("mouse_shape_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CollisionObject2D) DisconnectMouseShapeExited(subs SignalSubscribers, fn CollisionObject2DMouseShapeExitedSignalFn) {
	sig := StringNameFromStr("mouse_shape_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
