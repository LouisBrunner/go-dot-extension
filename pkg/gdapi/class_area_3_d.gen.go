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

type ptrsForArea3DList struct {
  fnSetGravitySpaceOverrideMode gdc.MethodBindPtr
  fnGetGravitySpaceOverrideMode gdc.MethodBindPtr
  fnSetGravityIsPoint gdc.MethodBindPtr
  fnIsGravityAPoint gdc.MethodBindPtr
  fnSetGravityPointUnitDistance gdc.MethodBindPtr
  fnGetGravityPointUnitDistance gdc.MethodBindPtr
  fnSetGravityPointCenter gdc.MethodBindPtr
  fnGetGravityPointCenter gdc.MethodBindPtr
  fnSetGravityDirection gdc.MethodBindPtr
  fnGetGravityDirection gdc.MethodBindPtr
  fnSetGravity gdc.MethodBindPtr
  fnGetGravity gdc.MethodBindPtr
  fnSetLinearDampSpaceOverrideMode gdc.MethodBindPtr
  fnGetLinearDampSpaceOverrideMode gdc.MethodBindPtr
  fnSetAngularDampSpaceOverrideMode gdc.MethodBindPtr
  fnGetAngularDampSpaceOverrideMode gdc.MethodBindPtr
  fnSetAngularDamp gdc.MethodBindPtr
  fnGetAngularDamp gdc.MethodBindPtr
  fnSetLinearDamp gdc.MethodBindPtr
  fnGetLinearDamp gdc.MethodBindPtr
  fnSetPriority gdc.MethodBindPtr
  fnGetPriority gdc.MethodBindPtr
  fnSetWindForceMagnitude gdc.MethodBindPtr
  fnGetWindForceMagnitude gdc.MethodBindPtr
  fnSetWindAttenuationFactor gdc.MethodBindPtr
  fnGetWindAttenuationFactor gdc.MethodBindPtr
  fnSetWindSourcePath gdc.MethodBindPtr
  fnGetWindSourcePath gdc.MethodBindPtr
  fnSetMonitorable gdc.MethodBindPtr
  fnIsMonitorable gdc.MethodBindPtr
  fnSetMonitoring gdc.MethodBindPtr
  fnIsMonitoring gdc.MethodBindPtr
  fnGetOverlappingBodies gdc.MethodBindPtr
  fnGetOverlappingAreas gdc.MethodBindPtr
  fnHasOverlappingBodies gdc.MethodBindPtr
  fnHasOverlappingAreas gdc.MethodBindPtr
  fnOverlapsBody gdc.MethodBindPtr
  fnOverlapsArea gdc.MethodBindPtr
  fnSetAudioBusOverride gdc.MethodBindPtr
  fnIsOverridingAudioBus gdc.MethodBindPtr
  fnSetAudioBusName gdc.MethodBindPtr
  fnGetAudioBusName gdc.MethodBindPtr
  fnSetUseReverbBus gdc.MethodBindPtr
  fnIsUsingReverbBus gdc.MethodBindPtr
  fnSetReverbBusName gdc.MethodBindPtr
  fnGetReverbBusName gdc.MethodBindPtr
  fnSetReverbAmount gdc.MethodBindPtr
  fnGetReverbAmount gdc.MethodBindPtr
  fnSetReverbUniformity gdc.MethodBindPtr
  fnGetReverbUniformity gdc.MethodBindPtr
}

var ptrsForArea3D ptrsForArea3DList

func initArea3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Area3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_gravity_space_override_mode")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetGravitySpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311433571))
  }
  {
    methodName := StringNameFromStr("get_gravity_space_override_mode")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetGravitySpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 958191869))
  }
  {
    methodName := StringNameFromStr("set_gravity_is_point")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetGravityIsPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_gravity_a_point")
    defer methodName.Destroy()
    ptrsForArea3D.fnIsGravityAPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_gravity_point_unit_distance")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetGravityPointUnitDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_gravity_point_unit_distance")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetGravityPointUnitDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_gravity_point_center")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetGravityPointCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_gravity_point_center")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetGravityPointCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_gravity_direction")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetGravityDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_gravity_direction")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetGravityDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_gravity")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_gravity")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_linear_damp_space_override_mode")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetLinearDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311433571))
  }
  {
    methodName := StringNameFromStr("get_linear_damp_space_override_mode")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetLinearDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 958191869))
  }
  {
    methodName := StringNameFromStr("set_angular_damp_space_override_mode")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetAngularDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311433571))
  }
  {
    methodName := StringNameFromStr("get_angular_damp_space_override_mode")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetAngularDampSpaceOverrideMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 958191869))
  }
  {
    methodName := StringNameFromStr("set_angular_damp")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_angular_damp")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_linear_damp")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_linear_damp")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_priority")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_priority")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_wind_force_magnitude")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetWindForceMagnitude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_wind_force_magnitude")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetWindForceMagnitude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_wind_attenuation_factor")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetWindAttenuationFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_wind_attenuation_factor")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetWindAttenuationFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_wind_source_path")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetWindSourcePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_wind_source_path")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetWindSourcePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_monitorable")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetMonitorable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_monitorable")
    defer methodName.Destroy()
    ptrsForArea3D.fnIsMonitorable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_monitoring")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetMonitoring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_monitoring")
    defer methodName.Destroy()
    ptrsForArea3D.fnIsMonitoring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_overlapping_bodies")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetOverlappingBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("get_overlapping_areas")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetOverlappingAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("has_overlapping_bodies")
    defer methodName.Destroy()
    ptrsForArea3D.fnHasOverlappingBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("has_overlapping_areas")
    defer methodName.Destroy()
    ptrsForArea3D.fnHasOverlappingAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("overlaps_body")
    defer methodName.Destroy()
    ptrsForArea3D.fnOverlapsBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093956946))
  }
  {
    methodName := StringNameFromStr("overlaps_area")
    defer methodName.Destroy()
    ptrsForArea3D.fnOverlapsArea = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093956946))
  }
  {
    methodName := StringNameFromStr("set_audio_bus_override")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetAudioBusOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_overriding_audio_bus")
    defer methodName.Destroy()
    ptrsForArea3D.fnIsOverridingAudioBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_audio_bus_name")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetAudioBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_audio_bus_name")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetAudioBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_use_reverb_bus")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetUseReverbBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_reverb_bus")
    defer methodName.Destroy()
    ptrsForArea3D.fnIsUsingReverbBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_reverb_bus_name")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetReverbBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_reverb_bus_name")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetReverbBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_reverb_amount")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetReverbAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_reverb_amount")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetReverbAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_reverb_uniformity")
    defer methodName.Destroy()
    ptrsForArea3D.fnSetReverbUniformity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_reverb_uniformity")
    defer methodName.Destroy()
    ptrsForArea3D.fnGetReverbUniformity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type Area3D struct {
  CollisionObject3D
}

func (me *Area3D) BaseClass() string {
  return "Area3D"
}

func NewArea3D() *Area3D {
  str := StringNameFromStr("Area3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Area3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetGravitySpaceOverrideMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetGravitySpaceOverrideMode() Area3DSpaceOverride {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Area3DSpaceOverride

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetGravitySpaceOverrideMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Area3D) SetGravityIsPoint(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetGravityIsPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) IsGravityAPoint() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnIsGravityAPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetGravityPointUnitDistance(distance_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetGravityPointUnitDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetGravityPointUnitDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetGravityPointUnitDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetGravityPointCenter(center Vector3, )  {
  cargs := []gdc.ConstTypePtr{center.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetGravityPointCenter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetGravityPointCenter() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetGravityPointCenter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Area3D) SetGravityDirection(direction Vector3, )  {
  cargs := []gdc.ConstTypePtr{direction.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetGravityDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetGravityDirection() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetGravityDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Area3D) SetGravity(gravity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetGravity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetLinearDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetLinearDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetLinearDampSpaceOverrideMode() Area3DSpaceOverride {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Area3DSpaceOverride

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetLinearDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Area3D) SetAngularDampSpaceOverrideMode(space_override_mode Area3DSpaceOverride, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&space_override_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetAngularDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetAngularDampSpaceOverrideMode() Area3DSpaceOverride {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Area3DSpaceOverride

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetAngularDampSpaceOverrideMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Area3D) SetAngularDamp(angular_damp float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetAngularDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetAngularDamp() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetAngularDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetLinearDamp(linear_damp float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetLinearDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetLinearDamp() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetLinearDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetPriority(priority int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetPriority() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetWindForceMagnitude(wind_force_magnitude float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wind_force_magnitude) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetWindForceMagnitude), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetWindForceMagnitude() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetWindForceMagnitude), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetWindAttenuationFactor(wind_attenuation_factor float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wind_attenuation_factor) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetWindAttenuationFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetWindAttenuationFactor() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetWindAttenuationFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetWindSourcePath(wind_source_path NodePath, )  {
  cargs := []gdc.ConstTypePtr{wind_source_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetWindSourcePath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetWindSourcePath() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetWindSourcePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Area3D) SetMonitorable(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetMonitorable), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) IsMonitorable() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnIsMonitorable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetMonitoring(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetMonitoring), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) IsMonitoring() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnIsMonitoring), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) GetOverlappingBodies() []Node3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetOverlappingBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Node3D](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Area3D) GetOverlappingAreas() []Area3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetOverlappingAreas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Area3D](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Area3D) HasOverlappingBodies() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnHasOverlappingBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) HasOverlappingAreas() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnHasOverlappingAreas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) OverlapsBody(body Node, ) bool {
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnOverlapsBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) OverlapsArea(area Node, ) bool {
  cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnOverlapsArea), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetAudioBusOverride(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetAudioBusOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) IsOverridingAudioBus() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnIsOverridingAudioBus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetAudioBusName(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetAudioBusName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetAudioBusName() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetAudioBusName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Area3D) SetUseReverbBus(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetUseReverbBus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) IsUsingReverbBus() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnIsUsingReverbBus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetReverbBusName(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetReverbBusName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetReverbBusName() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetReverbBusName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Area3D) SetReverbAmount(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetReverbAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetReverbAmount() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetReverbAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Area3D) SetReverbUniformity(amount float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnSetReverbUniformity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Area3D) GetReverbUniformity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArea3D.fnGetReverbUniformity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type Area3DBodyShapeEnteredSignalFn func(body_rid RID, body Node3D, body_shape_index int, local_shape_index int, )

func (me *Area3D) ConnectBodyShapeEntered(subs SignalSubscribers, fn Area3DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectBodyShapeEntered(subs SignalSubscribers, fn Area3DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Area3DBodyShapeExitedSignalFn func(body_rid RID, body Node3D, body_shape_index int, local_shape_index int, )

func (me *Area3D) ConnectBodyShapeExited(subs SignalSubscribers, fn Area3DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectBodyShapeExited(subs SignalSubscribers, fn Area3DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Area3DBodyEnteredSignalFn func(body Node3D, )

func (me *Area3D) ConnectBodyEntered(subs SignalSubscribers, fn Area3DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectBodyEntered(subs SignalSubscribers, fn Area3DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Area3DBodyExitedSignalFn func(body Node3D, )

func (me *Area3D) ConnectBodyExited(subs SignalSubscribers, fn Area3DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectBodyExited(subs SignalSubscribers, fn Area3DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Area3DAreaShapeEnteredSignalFn func(area_rid RID, area Area3D, area_shape_index int, local_shape_index int, )

func (me *Area3D) ConnectAreaShapeEntered(subs SignalSubscribers, fn Area3DAreaShapeEnteredSignalFn) {
  sig := StringNameFromStr("area_shape_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectAreaShapeEntered(subs SignalSubscribers, fn Area3DAreaShapeEnteredSignalFn) {
  sig := StringNameFromStr("area_shape_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Area3DAreaShapeExitedSignalFn func(area_rid RID, area Area3D, area_shape_index int, local_shape_index int, )

func (me *Area3D) ConnectAreaShapeExited(subs SignalSubscribers, fn Area3DAreaShapeExitedSignalFn) {
  sig := StringNameFromStr("area_shape_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectAreaShapeExited(subs SignalSubscribers, fn Area3DAreaShapeExitedSignalFn) {
  sig := StringNameFromStr("area_shape_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Area3DAreaEnteredSignalFn func(area Area3D, )

func (me *Area3D) ConnectAreaEntered(subs SignalSubscribers, fn Area3DAreaEnteredSignalFn) {
  sig := StringNameFromStr("area_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectAreaEntered(subs SignalSubscribers, fn Area3DAreaEnteredSignalFn) {
  sig := StringNameFromStr("area_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Area3DAreaExitedSignalFn func(area Area3D, )

func (me *Area3D) ConnectAreaExited(subs SignalSubscribers, fn Area3DAreaExitedSignalFn) {
  sig := StringNameFromStr("area_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Area3D) DisconnectAreaExited(subs SignalSubscribers, fn Area3DAreaExitedSignalFn) {
  sig := StringNameFromStr("area_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
