// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Area3D struct {
  obj gdc.ObjectPtr
}

func (me *Area3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Area3D) BaseClass() string {
  return "Area3D"
}

type Area3DSpaceOverride int
const (
  Area3DSpaceOverrideSpaceOverrideDisabled Area3DSpaceOverride = 0
  Area3DSpaceOverrideSpaceOverrideCombine Area3DSpaceOverride = 1
  Area3DSpaceOverrideSpaceOverrideCombineReplace Area3DSpaceOverride = 2
  Area3DSpaceOverrideSpaceOverrideReplace Area3DSpaceOverride = 3
  Area3DSpaceOverrideSpaceOverrideReplaceCombine Area3DSpaceOverride = 4
)

func  (me *Area3D) SetGravitySpaceOverrideMode(space_override_mode Area3DSpaceOverride, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetGravitySpaceOverrideMode() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetGravityIsPoint(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) IsGravityAPoint() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetGravityPointUnitDistance(distance_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetGravityPointUnitDistance() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetGravityPointCenter(center Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetGravityPointCenter() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetGravityDirection(direction Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetGravityDirection() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetGravity(gravity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetGravity() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetLinearDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetLinearDampSpaceOverrideMode() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetAngularDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetAngularDampSpaceOverrideMode() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetAngularDamp(angular_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetAngularDamp() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetLinearDamp(linear_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetLinearDamp() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetPriority(priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetPriority() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetWindForceMagnitude(wind_force_magnitude float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetWindForceMagnitude() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetWindAttenuationFactor(wind_attenuation_factor float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetWindAttenuationFactor() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetWindSourcePath(wind_source_path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetWindSourcePath() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetMonitorable(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) IsMonitorable() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetMonitoring(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) IsMonitoring() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetOverlappingBodies() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetOverlappingAreas() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) HasOverlappingBodies() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) HasOverlappingAreas() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) OverlapsBody(body Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) OverlapsArea(area Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetAudioBusOverride(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) IsOverridingAudioBus() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetAudioBusName(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetAudioBusName() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetUseReverbBus(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) IsUsingReverbBus() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetReverbBusName(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetReverbBusName() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetReverbAmount(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetReverbAmount() { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) SetReverbUniformity(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Area3D) GetReverbUniformity() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
