// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Area2D struct {
  obj gdc.ObjectPtr
}

func (me *Area2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Area2D) BaseClass() string {
  return "Area2D"
}



// Enums

type Area2DSpaceOverride int
const (
  Area2DSpaceOverrideSpaceOverrideDisabled Area2DSpaceOverride = 0
  Area2DSpaceOverrideSpaceOverrideCombine Area2DSpaceOverride = 1
  Area2DSpaceOverrideSpaceOverrideCombineReplace Area2DSpaceOverride = 2
  Area2DSpaceOverrideSpaceOverrideReplace Area2DSpaceOverride = 3
  Area2DSpaceOverrideSpaceOverrideReplaceCombine Area2DSpaceOverride = 4
)

func (me *Area2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Area2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Area2D) SetGravitySpaceOverrideMode(space_override_mode Area2DSpaceOverride, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetGravitySpaceOverrideMode()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetGravityIsPoint(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area2D) IsGravityAPoint()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetGravityPointUnitDistance(distance_scale float32, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetGravityPointUnitDistance()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetGravityPointCenter(center Vector2, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetGravityPointCenter()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetGravityDirection(direction Vector2, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetGravityDirection()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetGravity(gravity float32, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetGravity()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetLinearDampSpaceOverrideMode(space_override_mode Area2DSpaceOverride, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetLinearDampSpaceOverrideMode()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetAngularDampSpaceOverrideMode(space_override_mode Area2DSpaceOverride, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetAngularDampSpaceOverrideMode()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetLinearDamp(linear_damp float32, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetLinearDamp()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetAngularDamp(angular_damp float32, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetAngularDamp()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetPriority()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetMonitoring(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area2D) IsMonitoring()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetMonitorable(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area2D) IsMonitorable()  {
  panic("TODO: implement")
}

func  (me *Area2D) GetOverlappingBodies()  {
  panic("TODO: implement")
}

func  (me *Area2D) GetOverlappingAreas()  {
  panic("TODO: implement")
}

func  (me *Area2D) HasOverlappingBodies()  {
  panic("TODO: implement")
}

func  (me *Area2D) HasOverlappingAreas()  {
  panic("TODO: implement")
}

func  (me *Area2D) OverlapsBody(body Node, )  {
  panic("TODO: implement")
}

func  (me *Area2D) OverlapsArea(area Node, )  {
  panic("TODO: implement")
}

func  (me *Area2D) SetAudioBusName(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Area2D) GetAudioBusName()  {
  panic("TODO: implement")
}

func  (me *Area2D) SetAudioBusOverride(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Area2D) IsOverridingAudioBus()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
