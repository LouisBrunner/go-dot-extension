// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type RigidBody3D struct {
  PhysicsBody3D
}

func (me *RigidBody3D) BaseClass() string {
  return "RigidBody3D"
}

func NewRigidBody3D() *RigidBody3D {
  str := StringNameFromStr("RigidBody3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RigidBody3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type RigidBody3DFreezeMode int
const (
  RigidBody3DFreezeModeFreezeModeStatic RigidBody3DFreezeMode = 0
  RigidBody3DFreezeModeFreezeModeKinematic RigidBody3DFreezeMode = 1
)

type RigidBody3DCenterOfMassMode int
const (
  RigidBody3DCenterOfMassModeCenterOfMassModeAuto RigidBody3DCenterOfMassMode = 0
  RigidBody3DCenterOfMassModeCenterOfMassModeCustom RigidBody3DCenterOfMassMode = 1
)

type RigidBody3DDampMode int
const (
  RigidBody3DDampModeDampModeCombine RigidBody3DDampMode = 0
  RigidBody3DDampModeDampModeReplace RigidBody3DDampMode = 1
)

func (me *RigidBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RigidBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RigidBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RigidBody3D) SetMass(mass float64, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetMass() float64 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetInertia(inertia Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inertia")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{inertia.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetInertia() Vector3 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inertia")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) SetCenterOfMassMode(mode RigidBody3DCenterOfMassMode, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_of_mass_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3625866032) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetCenterOfMassMode() RigidBody3DCenterOfMassMode {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 237405040) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody3DCenterOfMassMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody3D) SetCenterOfMass(center_of_mass Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_of_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{center_of_mass.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetCenterOfMass() Vector3 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1784508650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{physics_material_override.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetPhysicsMaterialOverride() PhysicsMaterial {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2521850424) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPhysicsMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) SetLinearVelocity(linear_velocity Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{linear_velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetLinearVelocity() Vector3 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) SetAngularVelocity(angular_velocity Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{angular_velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetAngularVelocity() Vector3 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) GetInverseInertiaTensor() Basis {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inverse_inertia_tensor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2716978435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBasis()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) SetGravityScale(gravity_scale float64, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetGravityScale() float64 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetLinearDampMode(linear_damp_mode RigidBody3DDampMode, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1802035050) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetLinearDampMode() RigidBody3DDampMode {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1366206940) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody3DDampMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody3D) SetAngularDampMode(angular_damp_mode RigidBody3DDampMode, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1802035050) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetAngularDampMode() RigidBody3DDampMode {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1366206940) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody3DDampMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody3D) SetLinearDamp(linear_damp float64, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetLinearDamp() float64 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetAngularDamp(angular_damp float64, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetAngularDamp() float64 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetMaxContactsReported(amount int64, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetMaxContactsReported() int64 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) GetContactCount() int64 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetUseCustomIntegrator(enable bool, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_custom_integrator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) IsUsingCustomIntegrator() bool {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_custom_integrator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetContactMonitor(enabled bool, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_contact_monitor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) IsContactMonitorEnabled() bool {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_contact_monitor_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetUseContinuousCollisionDetection(enable bool, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_continuous_collision_detection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) IsUsingContinuousCollisionDetection() bool {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_continuous_collision_detection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetAxisVelocity(axis_velocity Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{axis_velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) ApplyCentralImpulse(impulse Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) ApplyImpulse(impulse Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2754756483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) ApplyTorqueImpulse(impulse Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) ApplyCentralForce(force Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) ApplyForce(force Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2754756483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) ApplyTorque(torque Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{torque.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) AddConstantCentralForce(force Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) AddConstantForce(force Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2754756483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) AddConstantTorque(torque Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{torque.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) SetConstantForce(force Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetConstantForce() Vector3 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) SetConstantTorque(torque Vector3, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{torque.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetConstantTorque() Vector3 {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody3D) SetSleeping(sleeping bool, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sleeping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sleeping) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) IsSleeping() bool {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sleeping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetCanSleep(able_to_sleep bool, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_can_sleep")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&able_to_sleep) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) IsAbleToSleep() bool {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_able_to_sleep")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetLockRotationEnabled(lock_rotation bool, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lock_rotation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lock_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) IsLockRotationEnabled() bool {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_lock_rotation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetFreezeEnabled(freeze_mode bool, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_freeze_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) IsFreezeEnabled() bool {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_freeze_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody3D) SetFreezeMode(freeze_mode RigidBody3DFreezeMode, )  {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_freeze_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1319914653) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody3D) GetFreezeMode() RigidBody3DFreezeMode {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_freeze_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2008423905) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody3DFreezeMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody3D) GetCollidingBodies() []Node3D {
  classNameV := StringNameFromStr("RigidBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_colliding_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Node3D](ret)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type RigidBody3DBodyShapeEnteredSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int, )

func (me *RigidBody3D) ConnectBodyShapeEntered(subs SignalSubscribers, fn RigidBody3DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody3D) DisconnectBodyShapeEntered(subs SignalSubscribers, fn RigidBody3DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody3DBodyShapeExitedSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int, )

func (me *RigidBody3D) ConnectBodyShapeExited(subs SignalSubscribers, fn RigidBody3DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody3D) DisconnectBodyShapeExited(subs SignalSubscribers, fn RigidBody3DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody3DBodyEnteredSignalFn func(body Node, )

func (me *RigidBody3D) ConnectBodyEntered(subs SignalSubscribers, fn RigidBody3DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody3D) DisconnectBodyEntered(subs SignalSubscribers, fn RigidBody3DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody3DBodyExitedSignalFn func(body Node, )

func (me *RigidBody3D) ConnectBodyExited(subs SignalSubscribers, fn RigidBody3DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody3D) DisconnectBodyExited(subs SignalSubscribers, fn RigidBody3DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody3DSleepingStateChangedSignalFn func()

func (me *RigidBody3D) ConnectSleepingStateChanged(subs SignalSubscribers, fn RigidBody3DSleepingStateChangedSignalFn) {
  sig := StringNameFromStr("sleeping_state_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody3D) DisconnectSleepingStateChanged(subs SignalSubscribers, fn RigidBody3DSleepingStateChangedSignalFn) {
  sig := StringNameFromStr("sleeping_state_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
