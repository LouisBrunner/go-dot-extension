// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RigidBody2D struct {
  obj gdc.ObjectPtr
}

func (me *RigidBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RigidBody2D) BaseClass() string {
  return "RigidBody2D"
}



// Enums

type RigidBody2DFreezeMode int
const (
  RigidBody2DFreezeModeFreezeModeStatic RigidBody2DFreezeMode = 0
  RigidBody2DFreezeModeFreezeModeKinematic RigidBody2DFreezeMode = 1
)

type RigidBody2DCenterOfMassMode int
const (
  RigidBody2DCenterOfMassModeCenterOfMassModeAuto RigidBody2DCenterOfMassMode = 0
  RigidBody2DCenterOfMassModeCenterOfMassModeCustom RigidBody2DCenterOfMassMode = 1
)

type RigidBody2DDampMode int
const (
  RigidBody2DDampModeDampModeCombine RigidBody2DDampMode = 0
  RigidBody2DDampModeDampModeReplace RigidBody2DDampMode = 1
)

type RigidBody2DCCDMode int
const (
  RigidBody2DCCDModeCcdModeDisabled RigidBody2DCCDMode = 0
  RigidBody2DCCDModeCcdModeCastRay RigidBody2DCCDMode = 1
  RigidBody2DCCDModeCcdModeCastShape RigidBody2DCCDMode = 2
)

func (me *RigidBody2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RigidBody2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RigidBody2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RigidBody2D) SetMass(mass float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetMass() float32 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) GetInertia() float32 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inertia")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetInertia(inertia float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inertia")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inertia), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) SetCenterOfMassMode(mode RigidBody2DCenterOfMassMode, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_of_mass_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1757235706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetCenterOfMassMode() RigidBody2DCenterOfMassMode {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3277132817) // FIXME: should cache?
  var ret RigidBody2DCenterOfMassMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetCenterOfMass(center_of_mass Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_of_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(center_of_mass.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetCenterOfMass() Vector2 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1784508650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(physics_material_override.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetPhysicsMaterialOverride() PhysicsMaterial {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2521850424) // FIXME: should cache?
  var ret PhysicsMaterial
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetGravityScale(gravity_scale float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetGravityScale() float32 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetLinearDampMode(linear_damp_mode RigidBody2DDampMode, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3406533708) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetLinearDampMode() RigidBody2DDampMode {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2970511462) // FIXME: should cache?
  var ret RigidBody2DDampMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetAngularDampMode(angular_damp_mode RigidBody2DDampMode, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3406533708) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetAngularDampMode() RigidBody2DDampMode {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2970511462) // FIXME: should cache?
  var ret RigidBody2DDampMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetLinearDamp(linear_damp float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetLinearDamp() float32 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetAngularDamp(angular_damp float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetAngularDamp() float32 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetLinearVelocity(linear_velocity Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(linear_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetLinearVelocity() Vector2 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetAngularVelocity(angular_velocity float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_velocity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetAngularVelocity() float32 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetMaxContactsReported(amount int, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetMaxContactsReported() int {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) GetContactCount() int {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetUseCustomIntegrator(enable bool, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_custom_integrator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) IsUsingCustomIntegrator() bool {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_custom_integrator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetContactMonitor(enabled bool, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_contact_monitor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) IsContactMonitorEnabled() bool {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_contact_monitor_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetContinuousCollisionDetectionMode(mode RigidBody2DCCDMode, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_continuous_collision_detection_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1000241384) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetContinuousCollisionDetectionMode() RigidBody2DCCDMode {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_continuous_collision_detection_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 815214376) // FIXME: should cache?
  var ret RigidBody2DCCDMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetAxisVelocity(axis_velocity Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(axis_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) ApplyCentralImpulse(impulse Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3862383994) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(impulse.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) ApplyImpulse(impulse Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4288681949) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(impulse.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) ApplyTorqueImpulse(torque float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) ApplyCentralForce(force Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) ApplyForce(force Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4288681949) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) ApplyTorque(torque float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) AddConstantCentralForce(force Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) AddConstantForce(force Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4288681949) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) AddConstantTorque(torque float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) SetConstantForce(force Vector2, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetConstantForce() Vector2 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetConstantTorque(torque float32, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetConstantTorque() float32 {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetSleeping(sleeping bool, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sleeping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sleeping), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) IsSleeping() bool {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sleeping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetCanSleep(able_to_sleep bool, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_can_sleep")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&able_to_sleep), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) IsAbleToSleep() bool {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_able_to_sleep")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetLockRotationEnabled(lock_rotation bool, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lock_rotation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lock_rotation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) IsLockRotationEnabled() bool {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_lock_rotation_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetFreezeEnabled(freeze_mode bool, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_freeze_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) IsFreezeEnabled() bool {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_freeze_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) SetFreezeMode(freeze_mode RigidBody2DFreezeMode, )  {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_freeze_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1705112154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RigidBody2D) GetFreezeMode() RigidBody2DFreezeMode {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_freeze_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2016872314) // FIXME: should cache?
  var ret RigidBody2DFreezeMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RigidBody2D) GetCollidingBodies() Node2D {
  classNameV := StringNameFromStr("RigidBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_colliding_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Node2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type RigidBody2DBodyShapeEnteredSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int, )

func (me *RigidBody2D) ConnectBodyShapeEntered(subs SignalSubscribers, fn RigidBody2DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyShapeEntered(subs SignalSubscribers, fn RigidBody2DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type RigidBody2DBodyShapeExitedSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int, )

func (me *RigidBody2D) ConnectBodyShapeExited(subs SignalSubscribers, fn RigidBody2DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyShapeExited(subs SignalSubscribers, fn RigidBody2DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type RigidBody2DBodyEnteredSignalFn func(body Node, )

func (me *RigidBody2D) ConnectBodyEntered(subs SignalSubscribers, fn RigidBody2DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyEntered(subs SignalSubscribers, fn RigidBody2DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type RigidBody2DBodyExitedSignalFn func(body Node, )

func (me *RigidBody2D) ConnectBodyExited(subs SignalSubscribers, fn RigidBody2DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyExited(subs SignalSubscribers, fn RigidBody2DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type RigidBody2DSleepingStateChangedSignalFn func()

func (me *RigidBody2D) ConnectSleepingStateChanged(subs SignalSubscribers, fn RigidBody2DSleepingStateChangedSignalFn) {
  sig := StringNameFromStr("sleeping_state_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectSleepingStateChanged(subs SignalSubscribers, fn RigidBody2DSleepingStateChangedSignalFn) {
  sig := StringNameFromStr("sleeping_state_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
