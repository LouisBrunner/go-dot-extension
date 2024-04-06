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

type ptrsForCharacterBody3DList struct {
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
	fnGetPlatformAngularVelocity   gdc.MethodBindPtr
	fnGetSlideCollisionCount       gdc.MethodBindPtr
	fnGetSlideCollision            gdc.MethodBindPtr
	fnGetLastSlideCollision        gdc.MethodBindPtr
}

var ptrsForCharacterBody3D ptrsForCharacterBody3DList

func initCharacterBody3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CharacterBody3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("move_and_slide")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnMoveAndSlide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("apply_floor_snap")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnApplyFloorSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_safe_margin")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetSafeMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_safe_margin")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetSafeMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("is_floor_stop_on_slope_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsFloorStopOnSlopeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_floor_stop_on_slope_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetFloorStopOnSlopeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_floor_constant_speed_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetFloorConstantSpeedEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_floor_constant_speed_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsFloorConstantSpeedEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_floor_block_on_wall_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetFloorBlockOnWallEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_floor_block_on_wall_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsFloorBlockOnWallEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_slide_on_ceiling_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetSlideOnCeilingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_slide_on_ceiling_enabled")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsSlideOnCeilingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_platform_floor_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetPlatformFloorLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_platform_floor_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetPlatformFloorLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_platform_wall_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetPlatformWallLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_platform_wall_layers")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetPlatformWallLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_max_slides")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetMaxSlides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_slides")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetMaxSlides = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_floor_max_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetFloorMaxAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_floor_max_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetFloorMaxAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_floor_snap_length")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetFloorSnapLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_floor_snap_length")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetFloorSnapLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_wall_min_slide_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetWallMinSlideAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_wall_min_slide_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetWallMinSlideAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_up_direction")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetUpDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_up_direction")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetUpDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_motion_mode")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetMotionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2690739026))
	}
	{
		methodName := StringNameFromStr("get_motion_mode")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetMotionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3529553604))
	}
	{
		methodName := StringNameFromStr("set_platform_on_leave")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnSetPlatformOnLeave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1459986142))
	}
	{
		methodName := StringNameFromStr("get_platform_on_leave")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetPlatformOnLeave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 996491171))
	}
	{
		methodName := StringNameFromStr("is_on_floor")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsOnFloor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_floor_only")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsOnFloorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_ceiling")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsOnCeiling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_ceiling_only")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsOnCeilingOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_wall")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsOnWall = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_on_wall_only")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnIsOnWallOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_floor_normal")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetFloorNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_wall_normal")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetWallNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_last_motion")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetLastMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_position_delta")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetPositionDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_real_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetRealVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_floor_angle")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetFloorAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2906300789))
	}
	{
		methodName := StringNameFromStr("get_platform_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetPlatformVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_platform_angular_velocity")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetPlatformAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_slide_collision_count")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetSlideCollisionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_slide_collision")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetSlideCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107003663))
	}
	{
		methodName := StringNameFromStr("get_last_slide_collision")
		defer methodName.Destroy()
		ptrsForCharacterBody3D.fnGetLastSlideCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 186875014))
	}
}

type CharacterBody3D struct {
	PhysicsBody3D
}

func (me *CharacterBody3D) BaseClass() string {
	return "CharacterBody3D"
}

func NewCharacterBody3D() *CharacterBody3D {
	str := StringNameFromStr("CharacterBody3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CharacterBody3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CharacterBody3DMotionMode int

const (
	CharacterBody3DMotionModeMotionModeGrounded CharacterBody3DMotionMode = 0
	CharacterBody3DMotionModeMotionModeFloating CharacterBody3DMotionMode = 1
)

type CharacterBody3DPlatformOnLeave int

const (
	CharacterBody3DPlatformOnLeavePlatformOnLeaveAddVelocity       CharacterBody3DPlatformOnLeave = 0
	CharacterBody3DPlatformOnLeavePlatformOnLeaveAddUpwardVelocity CharacterBody3DPlatformOnLeave = 1
	CharacterBody3DPlatformOnLeavePlatformOnLeaveDoNothing         CharacterBody3DPlatformOnLeave = 2
)

func (me *CharacterBody3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CharacterBody3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CharacterBody3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CharacterBody3D) MoveAndSlide() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnMoveAndSlide), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) ApplyFloorSnap() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnApplyFloorSnap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) SetVelocity(velocity Vector3) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) SetSafeMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetSafeMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetSafeMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetSafeMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) IsFloorStopOnSlopeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsFloorStopOnSlopeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetFloorStopOnSlopeEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetFloorStopOnSlopeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) SetFloorConstantSpeedEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetFloorConstantSpeedEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) IsFloorConstantSpeedEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsFloorConstantSpeedEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetFloorBlockOnWallEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetFloorBlockOnWallEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) IsFloorBlockOnWallEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsFloorBlockOnWallEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetSlideOnCeilingEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetSlideOnCeilingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) IsSlideOnCeilingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsSlideOnCeilingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetPlatformFloorLayers(exclude_layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetPlatformFloorLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetPlatformFloorLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetPlatformFloorLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetPlatformWallLayers(exclude_layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetPlatformWallLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetPlatformWallLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetPlatformWallLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) GetMaxSlides() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetMaxSlides), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetMaxSlides(max_slides int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_slides)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetMaxSlides), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetFloorMaxAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetFloorMaxAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetFloorMaxAngle(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetFloorMaxAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetFloorSnapLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetFloorSnapLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetFloorSnapLength(floor_snap_length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&floor_snap_length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetFloorSnapLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetWallMinSlideAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetWallMinSlideAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) SetWallMinSlideAngle(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetWallMinSlideAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetUpDirection() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetUpDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) SetUpDirection(up_direction Vector3) {
	cargs := []gdc.ConstTypePtr{up_direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetUpDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) SetMotionMode(mode CharacterBody3DMotionMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetMotionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetMotionMode() CharacterBody3DMotionMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CharacterBody3DMotionMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetMotionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CharacterBody3D) SetPlatformOnLeave(on_leave_apply_velocity CharacterBody3DPlatformOnLeave) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_leave_apply_velocity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnSetPlatformOnLeave), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CharacterBody3D) GetPlatformOnLeave() CharacterBody3DPlatformOnLeave {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CharacterBody3DPlatformOnLeave

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetPlatformOnLeave), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CharacterBody3D) IsOnFloor() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsOnFloor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) IsOnFloorOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsOnFloorOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) IsOnCeiling() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsOnCeiling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) IsOnCeilingOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsOnCeilingOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) IsOnWall() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsOnWall), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) IsOnWallOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnIsOnWallOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) GetFloorNormal() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetFloorNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetWallNormal() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetWallNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetLastMotion() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetLastMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetPositionDelta() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetPositionDelta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetRealVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetRealVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetFloorAngle(up_direction Vector3) float64 {
	cargs := []gdc.ConstTypePtr{up_direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetFloorAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) GetPlatformVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetPlatformVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetPlatformAngularVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetPlatformAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetSlideCollisionCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetSlideCollisionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CharacterBody3D) GetSlideCollision(slide_idx int64) KinematicCollision3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slide_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewKinematicCollision3D()
	pinner.Pin(&slide_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetSlideCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CharacterBody3D) GetLastSlideCollision() KinematicCollision3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewKinematicCollision3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCharacterBody3D.fnGetLastSlideCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
