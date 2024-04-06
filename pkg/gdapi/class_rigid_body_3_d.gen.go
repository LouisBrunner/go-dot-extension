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

type ptrsForRigidBody3DList struct {
	fnXIntegrateForces                    gdc.MethodBindPtr
	fnSetMass                             gdc.MethodBindPtr
	fnGetMass                             gdc.MethodBindPtr
	fnSetInertia                          gdc.MethodBindPtr
	fnGetInertia                          gdc.MethodBindPtr
	fnSetCenterOfMassMode                 gdc.MethodBindPtr
	fnGetCenterOfMassMode                 gdc.MethodBindPtr
	fnSetCenterOfMass                     gdc.MethodBindPtr
	fnGetCenterOfMass                     gdc.MethodBindPtr
	fnSetPhysicsMaterialOverride          gdc.MethodBindPtr
	fnGetPhysicsMaterialOverride          gdc.MethodBindPtr
	fnSetLinearVelocity                   gdc.MethodBindPtr
	fnGetLinearVelocity                   gdc.MethodBindPtr
	fnSetAngularVelocity                  gdc.MethodBindPtr
	fnGetAngularVelocity                  gdc.MethodBindPtr
	fnGetInverseInertiaTensor             gdc.MethodBindPtr
	fnSetGravityScale                     gdc.MethodBindPtr
	fnGetGravityScale                     gdc.MethodBindPtr
	fnSetLinearDampMode                   gdc.MethodBindPtr
	fnGetLinearDampMode                   gdc.MethodBindPtr
	fnSetAngularDampMode                  gdc.MethodBindPtr
	fnGetAngularDampMode                  gdc.MethodBindPtr
	fnSetLinearDamp                       gdc.MethodBindPtr
	fnGetLinearDamp                       gdc.MethodBindPtr
	fnSetAngularDamp                      gdc.MethodBindPtr
	fnGetAngularDamp                      gdc.MethodBindPtr
	fnSetMaxContactsReported              gdc.MethodBindPtr
	fnGetMaxContactsReported              gdc.MethodBindPtr
	fnGetContactCount                     gdc.MethodBindPtr
	fnSetUseCustomIntegrator              gdc.MethodBindPtr
	fnIsUsingCustomIntegrator             gdc.MethodBindPtr
	fnSetContactMonitor                   gdc.MethodBindPtr
	fnIsContactMonitorEnabled             gdc.MethodBindPtr
	fnSetUseContinuousCollisionDetection  gdc.MethodBindPtr
	fnIsUsingContinuousCollisionDetection gdc.MethodBindPtr
	fnSetAxisVelocity                     gdc.MethodBindPtr
	fnApplyCentralImpulse                 gdc.MethodBindPtr
	fnApplyImpulse                        gdc.MethodBindPtr
	fnApplyTorqueImpulse                  gdc.MethodBindPtr
	fnApplyCentralForce                   gdc.MethodBindPtr
	fnApplyForce                          gdc.MethodBindPtr
	fnApplyTorque                         gdc.MethodBindPtr
	fnAddConstantCentralForce             gdc.MethodBindPtr
	fnAddConstantForce                    gdc.MethodBindPtr
	fnAddConstantTorque                   gdc.MethodBindPtr
	fnSetConstantForce                    gdc.MethodBindPtr
	fnGetConstantForce                    gdc.MethodBindPtr
	fnSetConstantTorque                   gdc.MethodBindPtr
	fnGetConstantTorque                   gdc.MethodBindPtr
	fnSetSleeping                         gdc.MethodBindPtr
	fnIsSleeping                          gdc.MethodBindPtr
	fnSetCanSleep                         gdc.MethodBindPtr
	fnIsAbleToSleep                       gdc.MethodBindPtr
	fnSetLockRotationEnabled              gdc.MethodBindPtr
	fnIsLockRotationEnabled               gdc.MethodBindPtr
	fnSetFreezeEnabled                    gdc.MethodBindPtr
	fnIsFreezeEnabled                     gdc.MethodBindPtr
	fnSetFreezeMode                       gdc.MethodBindPtr
	fnGetFreezeMode                       gdc.MethodBindPtr
	fnGetCollidingBodies                  gdc.MethodBindPtr
}

var ptrsForRigidBody3D ptrsForRigidBody3DList

func initRigidBody3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RigidBody3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mass")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mass")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_inertia")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetInertia = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_inertia")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetInertia = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_center_of_mass_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetCenterOfMassMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3625866032))
	}
	{
		methodName := StringNameFromStr("get_center_of_mass_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetCenterOfMassMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 237405040))
	}
	{
		methodName := StringNameFromStr("set_center_of_mass")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_center_of_mass")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_physics_material_override")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1784508650))
	}
	{
		methodName := StringNameFromStr("get_physics_material_override")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2521850424))
	}
	{
		methodName := StringNameFromStr("set_linear_velocity")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_linear_velocity")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_angular_velocity")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_angular_velocity")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_inverse_inertia_tensor")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetInverseInertiaTensor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2716978435))
	}
	{
		methodName := StringNameFromStr("set_gravity_scale")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetGravityScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_gravity_scale")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetGravityScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_linear_damp_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetLinearDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1802035050))
	}
	{
		methodName := StringNameFromStr("get_linear_damp_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetLinearDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1366206940))
	}
	{
		methodName := StringNameFromStr("set_angular_damp_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetAngularDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1802035050))
	}
	{
		methodName := StringNameFromStr("get_angular_damp_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetAngularDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1366206940))
	}
	{
		methodName := StringNameFromStr("set_linear_damp")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_linear_damp")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_angular_damp")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_angular_damp")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_contacts_reported")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_contacts_reported")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_contact_count")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetContactCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_use_custom_integrator")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetUseCustomIntegrator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_custom_integrator")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnIsUsingCustomIntegrator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_contact_monitor")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetContactMonitor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_contact_monitor_enabled")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnIsContactMonitorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_continuous_collision_detection")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetUseContinuousCollisionDetection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_continuous_collision_detection")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnIsUsingContinuousCollisionDetection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_axis_velocity")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetAxisVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("apply_central_impulse")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnApplyCentralImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("apply_impulse")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnApplyImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2754756483))
	}
	{
		methodName := StringNameFromStr("apply_torque_impulse")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnApplyTorqueImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("apply_central_force")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnApplyCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("apply_force")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnApplyForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2754756483))
	}
	{
		methodName := StringNameFromStr("apply_torque")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnApplyTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("add_constant_central_force")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnAddConstantCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("add_constant_force")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnAddConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2754756483))
	}
	{
		methodName := StringNameFromStr("add_constant_torque")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnAddConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_constant_force")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_constant_force")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_constant_torque")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_constant_torque")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_sleeping")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetSleeping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_sleeping")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnIsSleeping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_can_sleep")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetCanSleep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_able_to_sleep")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnIsAbleToSleep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_lock_rotation_enabled")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetLockRotationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_lock_rotation_enabled")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnIsLockRotationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_freeze_enabled")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetFreezeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_freeze_enabled")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnIsFreezeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_freeze_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnSetFreezeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1319914653))
	}
	{
		methodName := StringNameFromStr("get_freeze_mode")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetFreezeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2008423905))
	}
	{
		methodName := StringNameFromStr("get_colliding_bodies")
		defer methodName.Destroy()
		ptrsForRigidBody3D.fnGetCollidingBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}

}

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
	RigidBody3DFreezeModeFreezeModeStatic    RigidBody3DFreezeMode = 0
	RigidBody3DFreezeModeFreezeModeKinematic RigidBody3DFreezeMode = 1
)

type RigidBody3DCenterOfMassMode int

const (
	RigidBody3DCenterOfMassModeCenterOfMassModeAuto   RigidBody3DCenterOfMassMode = 0
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

func (me *RigidBody3D) SetMass(mass float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetMass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetInertia(inertia Vector3) {
	cargs := []gdc.ConstTypePtr{inertia.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetInertia), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetInertia() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetInertia), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) SetCenterOfMassMode(mode RigidBody3DCenterOfMassMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetCenterOfMassMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetCenterOfMassMode() RigidBody3DCenterOfMassMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RigidBody3DCenterOfMassMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetCenterOfMassMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RigidBody3D) SetCenterOfMass(center_of_mass Vector3) {
	cargs := []gdc.ConstTypePtr{center_of_mass.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetCenterOfMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetCenterOfMass() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetCenterOfMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial) {
	cargs := []gdc.ConstTypePtr{physics_material_override.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetPhysicsMaterialOverride() PhysicsMaterial {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) SetLinearVelocity(linear_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{linear_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetLinearVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) SetAngularVelocity(angular_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{angular_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetAngularVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) GetInverseInertiaTensor() Basis {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBasis()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetInverseInertiaTensor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) SetGravityScale(gravity_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetGravityScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetGravityScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetGravityScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetLinearDampMode(linear_damp_mode RigidBody3DDampMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetLinearDampMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetLinearDampMode() RigidBody3DDampMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RigidBody3DDampMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetLinearDampMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RigidBody3D) SetAngularDampMode(angular_damp_mode RigidBody3DDampMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetAngularDampMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetAngularDampMode() RigidBody3DDampMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RigidBody3DDampMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetAngularDampMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RigidBody3D) SetLinearDamp(linear_damp float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetLinearDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetLinearDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetLinearDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetAngularDamp(angular_damp float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetAngularDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetAngularDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetAngularDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetMaxContactsReported(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetMaxContactsReported), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetMaxContactsReported() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetMaxContactsReported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) GetContactCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetContactCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetUseCustomIntegrator(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetUseCustomIntegrator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) IsUsingCustomIntegrator() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnIsUsingCustomIntegrator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetContactMonitor(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetContactMonitor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) IsContactMonitorEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnIsContactMonitorEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetUseContinuousCollisionDetection(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetUseContinuousCollisionDetection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) IsUsingContinuousCollisionDetection() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnIsUsingContinuousCollisionDetection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetAxisVelocity(axis_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{axis_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetAxisVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) ApplyCentralImpulse(impulse Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnApplyCentralImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) ApplyImpulse(impulse Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnApplyImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) ApplyTorqueImpulse(impulse Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnApplyTorqueImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) ApplyCentralForce(force Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnApplyCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) ApplyForce(force Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnApplyForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) ApplyTorque(torque Vector3) {
	cargs := []gdc.ConstTypePtr{torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnApplyTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) AddConstantCentralForce(force Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnAddConstantCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) AddConstantForce(force Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnAddConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) AddConstantTorque(torque Vector3) {
	cargs := []gdc.ConstTypePtr{torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnAddConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) SetConstantForce(force Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetConstantForce() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetConstantForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) SetConstantTorque(torque Vector3) {
	cargs := []gdc.ConstTypePtr{torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetConstantTorque() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetConstantTorque), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RigidBody3D) SetSleeping(sleeping bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sleeping)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetSleeping), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) IsSleeping() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnIsSleeping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetCanSleep(able_to_sleep bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&able_to_sleep)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetCanSleep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) IsAbleToSleep() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnIsAbleToSleep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetLockRotationEnabled(lock_rotation bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lock_rotation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetLockRotationEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) IsLockRotationEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnIsLockRotationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetFreezeEnabled(freeze_mode bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetFreezeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) IsFreezeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnIsFreezeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RigidBody3D) SetFreezeMode(freeze_mode RigidBody3DFreezeMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnSetFreezeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RigidBody3D) GetFreezeMode() RigidBody3DFreezeMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RigidBody3DFreezeMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetFreezeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RigidBody3D) GetCollidingBodies() []Node3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody3D.fnGetCollidingBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Node3D](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type RigidBody3DBodyShapeEnteredSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int)

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

type RigidBody3DBodyShapeExitedSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int)

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

type RigidBody3DBodyEnteredSignalFn func(body Node)

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

type RigidBody3DBodyExitedSignalFn func(body Node)

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
