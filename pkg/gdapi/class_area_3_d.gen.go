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
  panic("TODO: implement")
}

func  (me *Area3D) GetGravitySpaceOverrideMode()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetGravityIsPoint(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area3D) IsGravityAPoint()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetGravityPointUnitDistance(distance_scale float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetGravityPointUnitDistance()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetGravityPointCenter(center Vector3, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetGravityPointCenter()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetGravityDirection(direction Vector3, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetGravityDirection()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetGravity(gravity float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetGravity()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetLinearDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetLinearDampSpaceOverrideMode()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetAngularDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetAngularDampSpaceOverrideMode()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetAngularDamp(angular_damp float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetAngularDamp()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetLinearDamp(linear_damp float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetLinearDamp()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetPriority()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetWindForceMagnitude(wind_force_magnitude float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetWindForceMagnitude()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetWindAttenuationFactor(wind_attenuation_factor float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetWindAttenuationFactor()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetWindSourcePath(wind_source_path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetWindSourcePath()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetMonitorable(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area3D) IsMonitorable()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetMonitoring(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area3D) IsMonitoring()  {
  panic("TODO: implement")
}

func  (me *Area3D) GetOverlappingBodies()  {
  panic("TODO: implement")
}

func  (me *Area3D) GetOverlappingAreas()  {
  panic("TODO: implement")
}

func  (me *Area3D) HasOverlappingBodies()  {
  panic("TODO: implement")
}

func  (me *Area3D) HasOverlappingAreas()  {
  panic("TODO: implement")
}

func  (me *Area3D) OverlapsBody(body Node, )  {
  panic("TODO: implement")
}

func  (me *Area3D) OverlapsArea(area Node, )  {
  panic("TODO: implement")
}

func  (me *Area3D) SetAudioBusOverride(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area3D) IsOverridingAudioBus()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetAudioBusName(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetAudioBusName()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetUseReverbBus(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area3D) IsUsingReverbBus()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetReverbBusName(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetReverbBusName()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetReverbAmount(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetReverbAmount()  {
  panic("TODO: implement")
}

func  (me *Area3D) SetReverbUniformity(amount float32, )  {
  panic("TODO: implement")
}

func  (me *Area3D) GetReverbUniformity()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
