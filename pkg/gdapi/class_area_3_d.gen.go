// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Area3D struct {
  obj gdc.ObjectPtr
}

func (me *Area3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Area3D) BaseClass() string {
  return "Area3D"
}



// Enums

type Area3DSpaceOverride int
const (
  Area3DSpaceOverrideSpaceOverrideDisabled Area3DSpaceOverride = 0
  Area3DSpaceOverrideSpaceOverrideCombine Area3DSpaceOverride = 1
  Area3DSpaceOverrideSpaceOverrideCombineReplace Area3DSpaceOverride = 2
  Area3DSpaceOverrideSpaceOverrideReplace Area3DSpaceOverride = 3
  Area3DSpaceOverrideSpaceOverrideReplaceCombine Area3DSpaceOverride = 4
)

func (me *Area3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Area3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Area3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Area3D) SetGravitySpaceOverrideMode(space_override_mode Area3DSpaceOverride, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311433571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetGravitySpaceOverrideMode() Area3DSpaceOverride {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 958191869) // FIXME: should cache?
  var ret Area3DSpaceOverride
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetGravityIsPoint(enable bool, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_is_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) IsGravityAPoint() bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_gravity_a_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetGravityPointUnitDistance(distance_scale float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_point_unit_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetGravityPointUnitDistance() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_point_unit_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetGravityPointCenter(center Vector3, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_point_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(center.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetGravityPointCenter() Vector3 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_point_center")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetGravityDirection(direction Vector3, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(direction.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetGravityDirection() Vector3 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetGravity(gravity float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetGravity() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetLinearDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311433571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetLinearDampSpaceOverrideMode() Area3DSpaceOverride {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 958191869) // FIXME: should cache?
  var ret Area3DSpaceOverride
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetAngularDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311433571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetAngularDampSpaceOverrideMode() Area3DSpaceOverride {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp_space_override_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 958191869) // FIXME: should cache?
  var ret Area3DSpaceOverride
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetAngularDamp(angular_damp float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetAngularDamp() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetLinearDamp(linear_damp float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetLinearDamp() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetPriority(priority int, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetPriority() int {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetWindForceMagnitude(wind_force_magnitude float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wind_force_magnitude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wind_force_magnitude), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetWindForceMagnitude() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wind_force_magnitude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetWindAttenuationFactor(wind_attenuation_factor float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wind_attenuation_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wind_attenuation_factor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetWindAttenuationFactor() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wind_attenuation_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetWindSourcePath(wind_source_path NodePath, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wind_source_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(wind_source_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetWindSourcePath() NodePath {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wind_source_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetMonitorable(enable bool, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_monitorable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) IsMonitorable() bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_monitorable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetMonitoring(enable bool, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_monitoring")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) IsMonitoring() bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_monitoring")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) GetOverlappingBodies() Node3D {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_overlapping_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Node3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) GetOverlappingAreas() Area3D {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_overlapping_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Area3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) HasOverlappingBodies() bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_overlapping_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) HasOverlappingAreas() bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_overlapping_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) OverlapsBody(body Node, ) bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("overlaps_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093956946) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) OverlapsArea(area Node, ) bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("overlaps_area")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3093956946) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetAudioBusOverride(enable bool, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_bus_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) IsOverridingAudioBus() bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_overriding_audio_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetAudioBusName(name StringName, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetAudioBusName() StringName {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_audio_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetUseReverbBus(enable bool, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_reverb_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) IsUsingReverbBus() bool {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_reverb_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetReverbBusName(name StringName, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reverb_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetReverbBusName() StringName {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reverb_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetReverbAmount(amount float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reverb_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetReverbAmount() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reverb_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Area3D) SetReverbUniformity(amount float32, )  {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reverb_uniformity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Area3D) GetReverbUniformity() float32 {
  classNameV := StringNameFromStr("Area3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reverb_uniformity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Area3D) GetPropMonitoring() bool {
  panic("TODO: implement")
}

func (me *Area3D) SetPropMonitoring(value bool) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropMonitorable() bool {
  panic("TODO: implement")
}

func (me *Area3D) SetPropMonitorable(value bool) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropPriority() int {
  panic("TODO: implement")
}

func (me *Area3D) SetPropPriority(value int) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropGravitySpaceOverride() int {
  panic("TODO: implement")
}

func (me *Area3D) SetPropGravitySpaceOverride(value int) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropGravityPoint() bool {
  panic("TODO: implement")
}

func (me *Area3D) SetPropGravityPoint(value bool) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropGravityPointUnitDistance() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropGravityPointUnitDistance(value float32) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropGravityPointCenter() Vector3 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropGravityPointCenter(value Vector3) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropGravityDirection() Vector3 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropGravityDirection(value Vector3) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropGravity() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropGravity(value float32) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropLinearDampSpaceOverride() int {
  panic("TODO: implement")
}

func (me *Area3D) SetPropLinearDampSpaceOverride(value int) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropLinearDamp() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropLinearDamp(value float32) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropAngularDampSpaceOverride() int {
  panic("TODO: implement")
}

func (me *Area3D) SetPropAngularDampSpaceOverride(value int) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropAngularDamp() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropAngularDamp(value float32) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropWindForceMagnitude() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropWindForceMagnitude(value float32) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropWindAttenuationFactor() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropWindAttenuationFactor(value float32) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropWindSourcePath() NodePath {
  panic("TODO: implement")
}

func (me *Area3D) SetPropWindSourcePath(value NodePath) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropAudioBusOverride() bool {
  panic("TODO: implement")
}

func (me *Area3D) SetPropAudioBusOverride(value bool) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropAudioBusName() StringName {
  panic("TODO: implement")
}

func (me *Area3D) SetPropAudioBusName(value StringName) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropReverbBusEnabled() bool {
  panic("TODO: implement")
}

func (me *Area3D) SetPropReverbBusEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropReverbBusName() StringName {
  panic("TODO: implement")
}

func (me *Area3D) SetPropReverbBusName(value StringName) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropReverbBusAmount() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropReverbBusAmount(value float32) {
  panic("TODO: implement")
}

func (me *Area3D) GetPropReverbBusUniformity() float32 {
  panic("TODO: implement")
}

func (me *Area3D) SetPropReverbBusUniformity(value float32) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API