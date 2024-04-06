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

type ptrsForArea2DList struct {
	fnSetGravitySpaceOverrideMode     gdc.MethodBindPtr
	fnGetGravitySpaceOverrideMode     gdc.MethodBindPtr
	fnSetGravityIsPoint               gdc.MethodBindPtr
	fnIsGravityAPoint                 gdc.MethodBindPtr
	fnSetGravityPointUnitDistance     gdc.MethodBindPtr
	fnGetGravityPointUnitDistance     gdc.MethodBindPtr
	fnSetGravityPointCenter           gdc.MethodBindPtr
	fnGetGravityPointCenter           gdc.MethodBindPtr
	fnSetGravityDirection             gdc.MethodBindPtr
	fnGetGravityDirection             gdc.MethodBindPtr
	fnSetGravity                      gdc.MethodBindPtr
	fnGetGravity                      gdc.MethodBindPtr
	fnSetLinearDampSpaceOverrideMode  gdc.MethodBindPtr
	fnGetLinearDampSpaceOverrideMode  gdc.MethodBindPtr
	fnSetAngularDampSpaceOverrideMode gdc.MethodBindPtr
	fnGetAngularDampSpaceOverrideMode gdc.MethodBindPtr
	fnSetLinearDamp                   gdc.MethodBindPtr
	fnGetLinearDamp                   gdc.MethodBindPtr
	fnSetAngularDamp                  gdc.MethodBindPtr
	fnGetAngularDamp                  gdc.MethodBindPtr
	fnSetPriority                     gdc.MethodBindPtr
	fnGetPriority                     gdc.MethodBindPtr
	fnSetMonitoring                   gdc.MethodBindPtr
	fnIsMonitoring                    gdc.MethodBindPtr
	fnSetMonitorable                  gdc.MethodBindPtr
	fnIsMonitorable                   gdc.MethodBindPtr
	fnGetOverlappingBodies            gdc.MethodBindPtr
	fnGetOverlappingAreas             gdc.MethodBindPtr
	fnHasOverlappingBodies            gdc.MethodBindPtr
	fnHasOverlappingAreas             gdc.MethodBindPtr
	fnOverlapsBody                    gdc.MethodBindPtr
	fnOverlapsArea                    gdc.MethodBindPtr
	fnSetAudioBusName                 gdc.MethodBindPtr
	fnGetAudioBusName                 gdc.MethodBindPtr
	fnSetAudioBusOverride             gdc.MethodBindPtr
	fnIsOverridingAudioBus            gdc.MethodBindPtr
}

var ptrsForArea2D ptrsForArea2DList

func initArea2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Area2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_gravity_space_override_mode")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetGravitySpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2879900038))
	}
	{
		methodName := StringNameFromStr("get_gravity_space_override_mode")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetGravitySpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3990256304))
	}
	{
		methodName := StringNameFromStr("set_gravity_is_point")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetGravityIsPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_gravity_a_point")
		defer methodName.Destroy()
		ptrsForArea2D.fnIsGravityAPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_gravity_point_unit_distance")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetGravityPointUnitDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_gravity_point_unit_distance")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetGravityPointUnitDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_gravity_point_center")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetGravityPointCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_gravity_point_center")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetGravityPointCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_gravity_direction")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetGravityDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_gravity_direction")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetGravityDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_gravity")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_gravity")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_linear_damp_space_override_mode")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetLinearDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2879900038))
	}
	{
		methodName := StringNameFromStr("get_linear_damp_space_override_mode")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetLinearDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3990256304))
	}
	{
		methodName := StringNameFromStr("set_angular_damp_space_override_mode")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetAngularDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2879900038))
	}
	{
		methodName := StringNameFromStr("get_angular_damp_space_override_mode")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetAngularDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3990256304))
	}
	{
		methodName := StringNameFromStr("set_linear_damp")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_linear_damp")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_angular_damp")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_angular_damp")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_priority")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_priority")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_monitoring")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetMonitoring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_monitoring")
		defer methodName.Destroy()
		ptrsForArea2D.fnIsMonitoring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_monitorable")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetMonitorable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_monitorable")
		defer methodName.Destroy()
		ptrsForArea2D.fnIsMonitorable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_overlapping_bodies")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetOverlappingBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_overlapping_areas")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetOverlappingAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("has_overlapping_bodies")
		defer methodName.Destroy()
		ptrsForArea2D.fnHasOverlappingBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("has_overlapping_areas")
		defer methodName.Destroy()
		ptrsForArea2D.fnHasOverlappingAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("overlaps_body")
		defer methodName.Destroy()
		ptrsForArea2D.fnOverlapsBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093956946))
	}
	{
		methodName := StringNameFromStr("overlaps_area")
		defer methodName.Destroy()
		ptrsForArea2D.fnOverlapsArea = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093956946))
	}
	{
		methodName := StringNameFromStr("set_audio_bus_name")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetAudioBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_audio_bus_name")
		defer methodName.Destroy()
		ptrsForArea2D.fnGetAudioBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_audio_bus_override")
		defer methodName.Destroy()
		ptrsForArea2D.fnSetAudioBusOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_overriding_audio_bus")
		defer methodName.Destroy()
		ptrsForArea2D.fnIsOverridingAudioBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type Area2D struct {
	CollisionObject2D
}

func (me *Area2D) BaseClass() string {
	return "Area2D"
}

func NewArea2D() *Area2D {
	str := StringNameFromStr("Area2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Area2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type Area2DSpaceOverride int

const (
	Area2DSpaceOverrideSpaceOverrideDisabled       Area2DSpaceOverride = 0
	Area2DSpaceOverrideSpaceOverrideCombine        Area2DSpaceOverride = 1
	Area2DSpaceOverrideSpaceOverrideCombineReplace Area2DSpaceOverride = 2
	Area2DSpaceOverrideSpaceOverrideReplace        Area2DSpaceOverride = 3
	Area2DSpaceOverrideSpaceOverrideReplaceCombine Area2DSpaceOverride = 4
)

func (me *Area2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Area2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Area2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Area2D) SetGravitySpaceOverrideMode(space_override_mode Area2DSpaceOverride) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetGravitySpaceOverrideMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetGravitySpaceOverrideMode() Area2DSpaceOverride {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Area2DSpaceOverride

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetGravitySpaceOverrideMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Area2D) SetGravityIsPoint(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetGravityIsPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) IsGravityAPoint() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnIsGravityAPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetGravityPointUnitDistance(distance_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetGravityPointUnitDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetGravityPointUnitDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetGravityPointUnitDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetGravityPointCenter(center Vector2) {
	cargs := []gdc.ConstTypePtr{center.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetGravityPointCenter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetGravityPointCenter() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetGravityPointCenter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Area2D) SetGravityDirection(direction Vector2) {
	cargs := []gdc.ConstTypePtr{direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetGravityDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetGravityDirection() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetGravityDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Area2D) SetGravity(gravity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetGravity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetLinearDampSpaceOverrideMode(space_override_mode Area2DSpaceOverride) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetLinearDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetLinearDampSpaceOverrideMode() Area2DSpaceOverride {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Area2DSpaceOverride

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetLinearDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Area2D) SetAngularDampSpaceOverrideMode(space_override_mode Area2DSpaceOverride) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetAngularDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetAngularDampSpaceOverrideMode() Area2DSpaceOverride {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Area2DSpaceOverride

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetAngularDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Area2D) SetLinearDamp(linear_damp float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetLinearDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetLinearDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetLinearDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetAngularDamp(angular_damp float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetAngularDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetAngularDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetAngularDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetMonitoring(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetMonitoring), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) IsMonitoring() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnIsMonitoring), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetMonitorable(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetMonitorable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) IsMonitorable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnIsMonitorable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) GetOverlappingBodies() []Node2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetOverlappingBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Node2D](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Area2D) GetOverlappingAreas() []Area2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetOverlappingAreas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Area2D](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Area2D) HasOverlappingBodies() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnHasOverlappingBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) HasOverlappingAreas() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnHasOverlappingAreas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) OverlapsBody(body Node) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnOverlapsBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) OverlapsArea(area Node) bool {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnOverlapsArea), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Area2D) SetAudioBusName(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetAudioBusName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) GetAudioBusName() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnGetAudioBusName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Area2D) SetAudioBusOverride(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnSetAudioBusOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Area2D) IsOverridingAudioBus() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea2D.fnIsOverridingAudioBus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type Area2DBodyShapeEnteredSignalFn func(body_rid RID, body Node2D, body_shape_index int, local_shape_index int)

func (me *Area2D) ConnectBodyShapeEntered(subs SignalSubscribers, fn Area2DBodyShapeEnteredSignalFn) {
	sig := StringNameFromStr("body_shape_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectBodyShapeEntered(subs SignalSubscribers, fn Area2DBodyShapeEnteredSignalFn) {
	sig := StringNameFromStr("body_shape_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Area2DBodyShapeExitedSignalFn func(body_rid RID, body Node2D, body_shape_index int, local_shape_index int)

func (me *Area2D) ConnectBodyShapeExited(subs SignalSubscribers, fn Area2DBodyShapeExitedSignalFn) {
	sig := StringNameFromStr("body_shape_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectBodyShapeExited(subs SignalSubscribers, fn Area2DBodyShapeExitedSignalFn) {
	sig := StringNameFromStr("body_shape_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Area2DBodyEnteredSignalFn func(body Node2D)

func (me *Area2D) ConnectBodyEntered(subs SignalSubscribers, fn Area2DBodyEnteredSignalFn) {
	sig := StringNameFromStr("body_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectBodyEntered(subs SignalSubscribers, fn Area2DBodyEnteredSignalFn) {
	sig := StringNameFromStr("body_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Area2DBodyExitedSignalFn func(body Node2D)

func (me *Area2D) ConnectBodyExited(subs SignalSubscribers, fn Area2DBodyExitedSignalFn) {
	sig := StringNameFromStr("body_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectBodyExited(subs SignalSubscribers, fn Area2DBodyExitedSignalFn) {
	sig := StringNameFromStr("body_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Area2DAreaShapeEnteredSignalFn func(area_rid RID, area Area2D, area_shape_index int, local_shape_index int)

func (me *Area2D) ConnectAreaShapeEntered(subs SignalSubscribers, fn Area2DAreaShapeEnteredSignalFn) {
	sig := StringNameFromStr("area_shape_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectAreaShapeEntered(subs SignalSubscribers, fn Area2DAreaShapeEnteredSignalFn) {
	sig := StringNameFromStr("area_shape_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Area2DAreaShapeExitedSignalFn func(area_rid RID, area Area2D, area_shape_index int, local_shape_index int)

func (me *Area2D) ConnectAreaShapeExited(subs SignalSubscribers, fn Area2DAreaShapeExitedSignalFn) {
	sig := StringNameFromStr("area_shape_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectAreaShapeExited(subs SignalSubscribers, fn Area2DAreaShapeExitedSignalFn) {
	sig := StringNameFromStr("area_shape_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Area2DAreaEnteredSignalFn func(area Area2D)

func (me *Area2D) ConnectAreaEntered(subs SignalSubscribers, fn Area2DAreaEnteredSignalFn) {
	sig := StringNameFromStr("area_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectAreaEntered(subs SignalSubscribers, fn Area2DAreaEnteredSignalFn) {
	sig := StringNameFromStr("area_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Area2DAreaExitedSignalFn func(area Area2D)

func (me *Area2D) ConnectAreaExited(subs SignalSubscribers, fn Area2DAreaExitedSignalFn) {
	sig := StringNameFromStr("area_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area2D) DisconnectAreaExited(subs SignalSubscribers, fn Area2DAreaExitedSignalFn) {
	sig := StringNameFromStr("area_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
