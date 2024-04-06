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

type ptrsForCharacterBody2DList struct {
	fnMoveAndSlide                 gdc.MethodBindPtr
	fnApplyFloorSnap               gdc.MethodBindPtr
	fnSetVelocity                  gdc.MethodBindPtr
	fnGetVelocity                  gdc.MethodBindPtr
	fnSetSafeMargin                gdc.MethodBindPtr
	fnGetSafeMargin                gdc.MethodBindPtr
	fnIsFloorStopOnSlopeEnabled    gdc.MethodBindPtr
	fnSetFloorStopOnSlopeEnabled   gdc.MethodBindPtr
	fnSetFloorConstantSpeedEnabled gdc.MethodBindPtr
	fnIsFloorConstantSpeedEnabled  gdc.MethodBindPtr
	fnSetFloorBlockOnWallEnabled   gdc.MethodBindPtr
	fnIsFloorBlockOnWallEnabled    gdc.MethodBindPtr
	fnSetSlideOnCeilingEnabled     gdc.MethodBindPtr
	fnIsSlideOnCeilingEnabled      gdc.MethodBindPtr
	fnSetPlatformFloorLayers       gdc.MethodBindPtr
	fnGetPlatformFloorLayers       gdc.MethodBindPtr
	fnSetPlatformWallLayers        gdc.MethodBindPtr
	fnGetPlatformWallLayers        gdc.MethodBindPtr
	fnGetMaxSlides                 gdc.MethodBindPtr
	fnSetMaxSlides                 gdc.MethodBindPtr
	fnGetFloorMaxAngle             gdc.MethodBindPtr
	fnSetFloorMaxAngle             gdc.MethodBindPtr
	fnGetFloorSnapLength           gdc.MethodBindPtr
	fnSetFloorSnapLength           gdc.MethodBindPtr
	fnGetWallMinSlideAngle         gdc.MethodBindPtr
	fnSetWallMinSlideAngle         gdc.MethodBindPtr
	fnGetUpDirection               gdc.MethodBindPtr
	fnSetUpDirection               gdc.MethodBindPtr
	fnSetMotionMode                gdc.MethodBindPtr
	fnGetMotionMode                gdc.MethodBindPtr
	fnSetPlatformOnLeave           gdc.MethodBindPtr
	fnGetPlatformOnLeave           gdc.MethodBindPtr
	fnIsOnFloor                    gdc.MethodBindPtr
	fnIsOnFloorOnly                gdc.MethodBindPtr
	fnIsOnCeiling                  gdc.MethodBindPtr
	fnIsOnCeilingOnly              gdc.MethodBindPtr
	fnIsOnWall                     gdc.MethodBindPtr
	fnIsOnWallOnly                 gdc.MethodBindPtr
	fnGetFloorNormal               gdc.MethodBindPtr
	fnGetWallNormal                gdc.MethodBindPtr
	fnGetLastMotion                gdc.MethodBindPtr
	fnGetPositionDelta             gdc.MethodBindPtr
	fnGetRealVelocity              gdc.MethodBindPtr
	fnGetFloorAngle                gdc.MethodBindPtr
	fnGetPlatformVelocity          gdc.MethodBindPtr
	fnGetSlideCollisionCount       gdc.MethodBindPtr
	fnGetSlideCollision            gdc.MethodBindPtr
	fnGetLastSlideCollision        gdc.MethodBindPtr
}

var ptrsForCharacterBody2D ptrsForCharacterBody2DList

func initCharacterBody2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CharacterBody2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("move_and_slide")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnMoveAndSlide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("apply_floor_snap")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnApplyFloorSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_safe_margin")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetSafeMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_safe_margin")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetSafeMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("is_floor_stop_on_slope_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsFloorStopOnSlopeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_floor_stop_on_slope_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetFloorStopOnSlopeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_floor_constant_speed_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetFloorConstantSpeedEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_floor_constant_speed_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsFloorConstantSpeedEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_floor_block_on_wall_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetFloorBlockOnWallEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_floor_block_on_wall_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsFloorBlockOnWallEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_slide_on_ceiling_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetSlideOnCeilingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_slide_on_ceiling_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsSlideOnCeilingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_platform_floor_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetPlatformFloorLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_platform_floor_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetPlatformFloorLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_platform_wall_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetPlatformWallLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_platform_wall_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetPlatformWallLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_max_slides")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetMaxSlides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_slides")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetMaxSlides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_floor_max_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetFloorMaxAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_floor_max_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetFloorMaxAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_floor_snap_length")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetFloorSnapLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_floor_snap_length")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetFloorSnapLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_wall_min_slide_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetWallMinSlideAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_wall_min_slide_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetWallMinSlideAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_up_direction")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetUpDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_up_direction")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetUpDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_motion_mode")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetMotionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1224392233))
	}
	{
		methodName := StringNameFromStr("get_motion_mode")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetMotionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1160151236))
	}
	{
		methodName := StringNameFromStr("set_platform_on_leave")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnSetPlatformOnLeave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2423324375))
	}
	{
		methodName := StringNameFromStr("get_platform_on_leave")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetPlatformOnLeave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4054324341))
	}
	{
		methodName := StringNameFromStr("is_on_floor")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsOnFloor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_floor_only")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsOnFloorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_ceiling")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsOnCeiling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_ceiling_only")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsOnCeilingOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_wall")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsOnWall = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_wall_only")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnIsOnWallOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_floor_normal")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetFloorNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_wall_normal")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetWallNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_last_motion")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetLastMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_position_delta")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetPositionDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_real_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetRealVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_floor_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetFloorAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841063350))
	}
	{
		methodName := StringNameFromStr("get_platform_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetPlatformVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_slide_collision_count")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetSlideCollisionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_slide_collision")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetSlideCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 860659811))
	}
	{
		methodName := StringNameFromStr("get_last_slide_collision")
		defer methodName.Destroy()
		ptrsForCharacterBody2D.fnGetLastSlideCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2161834755))
	}

}

type CharacterBody2D struct {
	PhysicsBody2D
}

func (me *CharacterBody2D) BaseClass() string {
	return "CharacterBody2D"
}

func NewCharacterBody2D() *CharacterBody2D {
	str := StringNameFromStr("CharacterBody2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CharacterBody2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CharacterBody2DMotionMode int

const (
	CharacterBody2DMotionModeMotionModeGrounded CharacterBody2DMotionMode = 0
	CharacterBody2DMotionModeMotionModeFloating CharacterBody2DMotionMode = 1
)

type CharacterBody2DPlatformOnLeave int

const (
	CharacterBody2DPlatformOnLeavePlatformOnLeaveAddVelocity       CharacterBody2DPlatformOnLeave = 0
	CharacterBody2DPlatformOnLeavePlatformOnLeaveAddUpwardVelocity CharacterBody2DPlatformOnLeave = 1
	CharacterBody2DPlatformOnLeavePlatformOnLeaveDoNothing         CharacterBody2DPlatformOnLeave = 2
)

func (me *CharacterBody2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CharacterBody2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CharacterBody2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CharacterBody2D) MoveAndSlide() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnMoveAndSlide), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) ApplyFloorSnap() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnApplyFloorSnap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) SetVelocity(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) SetSafeMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetSafeMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetSafeMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetSafeMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) IsFloorStopOnSlopeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsFloorStopOnSlopeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetFloorStopOnSlopeEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetFloorStopOnSlopeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) SetFloorConstantSpeedEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetFloorConstantSpeedEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) IsFloorConstantSpeedEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsFloorConstantSpeedEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetFloorBlockOnWallEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetFloorBlockOnWallEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) IsFloorBlockOnWallEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsFloorBlockOnWallEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetSlideOnCeilingEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetSlideOnCeilingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) IsSlideOnCeilingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsSlideOnCeilingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetPlatformFloorLayers(exclude_layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetPlatformFloorLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetPlatformFloorLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetPlatformFloorLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetPlatformWallLayers(exclude_layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetPlatformWallLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetPlatformWallLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetPlatformWallLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) GetMaxSlides() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetMaxSlides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetMaxSlides(max_slides int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_slides)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetMaxSlides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetFloorMaxAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetFloorMaxAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetFloorMaxAngle(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetFloorMaxAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetFloorSnapLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetFloorSnapLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetFloorSnapLength(floor_snap_length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&floor_snap_length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetFloorSnapLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetWallMinSlideAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetWallMinSlideAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) SetWallMinSlideAngle(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetWallMinSlideAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetUpDirection() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetUpDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) SetUpDirection(up_direction Vector2) {
	cargs := []gdc.ConstTypePtr{up_direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetUpDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) SetMotionMode(mode CharacterBody2DMotionMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetMotionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetMotionMode() CharacterBody2DMotionMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CharacterBody2DMotionMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetMotionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CharacterBody2D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody2DPlatformOnLeave) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_leave_apply_velocity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnSetPlatformOnLeave), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody2D) GetPlatformOnLeave() CharacterBody2DPlatformOnLeave {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CharacterBody2DPlatformOnLeave

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetPlatformOnLeave), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CharacterBody2D) IsOnFloor() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsOnFloor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) IsOnFloorOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsOnFloorOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) IsOnCeiling() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsOnCeiling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) IsOnCeilingOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsOnCeilingOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) IsOnWall() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsOnWall), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) IsOnWallOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnIsOnWallOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) GetFloorNormal() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetFloorNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) GetWallNormal() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetWallNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) GetLastMotion() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetLastMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) GetPositionDelta() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetPositionDelta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) GetRealVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetRealVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) GetFloorAngle(up_direction Vector2) float64 {
	cargs := []gdc.ConstTypePtr{up_direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetFloorAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) GetPlatformVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetPlatformVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) GetSlideCollisionCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetSlideCollisionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody2D) GetSlideCollision(slide_idx int64) KinematicCollision2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slide_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewKinematicCollision2D()
	pinner.Pin(&slide_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetSlideCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody2D) GetLastSlideCollision() KinematicCollision2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewKinematicCollision2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody2D.fnGetLastSlideCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
