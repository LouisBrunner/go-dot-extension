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

type ptrsForPhysicsDirectBodyState3DList struct {
	fnGetTotalGravity                      gdc.MethodBindPtr
	fnGetTotalLinearDamp                   gdc.MethodBindPtr
	fnGetTotalAngularDamp                  gdc.MethodBindPtr
	fnGetCenterOfMass                      gdc.MethodBindPtr
	fnGetCenterOfMassLocal                 gdc.MethodBindPtr
	fnGetPrincipalInertiaAxes              gdc.MethodBindPtr
	fnGetInverseMass                       gdc.MethodBindPtr
	fnGetInverseInertia                    gdc.MethodBindPtr
	fnGetInverseInertiaTensor              gdc.MethodBindPtr
	fnSetLinearVelocity                    gdc.MethodBindPtr
	fnGetLinearVelocity                    gdc.MethodBindPtr
	fnSetAngularVelocity                   gdc.MethodBindPtr
	fnGetAngularVelocity                   gdc.MethodBindPtr
	fnSetTransform                         gdc.MethodBindPtr
	fnGetTransform                         gdc.MethodBindPtr
	fnGetVelocityAtLocalPosition           gdc.MethodBindPtr
	fnApplyCentralImpulse                  gdc.MethodBindPtr
	fnApplyImpulse                         gdc.MethodBindPtr
	fnApplyTorqueImpulse                   gdc.MethodBindPtr
	fnApplyCentralForce                    gdc.MethodBindPtr
	fnApplyForce                           gdc.MethodBindPtr
	fnApplyTorque                          gdc.MethodBindPtr
	fnAddConstantCentralForce              gdc.MethodBindPtr
	fnAddConstantForce                     gdc.MethodBindPtr
	fnAddConstantTorque                    gdc.MethodBindPtr
	fnSetConstantForce                     gdc.MethodBindPtr
	fnGetConstantForce                     gdc.MethodBindPtr
	fnSetConstantTorque                    gdc.MethodBindPtr
	fnGetConstantTorque                    gdc.MethodBindPtr
	fnSetSleepState                        gdc.MethodBindPtr
	fnIsSleeping                           gdc.MethodBindPtr
	fnGetContactCount                      gdc.MethodBindPtr
	fnGetContactLocalPosition              gdc.MethodBindPtr
	fnGetContactLocalNormal                gdc.MethodBindPtr
	fnGetContactImpulse                    gdc.MethodBindPtr
	fnGetContactLocalShape                 gdc.MethodBindPtr
	fnGetContactLocalVelocityAtPosition    gdc.MethodBindPtr
	fnGetContactCollider                   gdc.MethodBindPtr
	fnGetContactColliderPosition           gdc.MethodBindPtr
	fnGetContactColliderId                 gdc.MethodBindPtr
	fnGetContactColliderObject             gdc.MethodBindPtr
	fnGetContactColliderShape              gdc.MethodBindPtr
	fnGetContactColliderVelocityAtPosition gdc.MethodBindPtr
	fnGetStep                              gdc.MethodBindPtr
	fnIntegrateForces                      gdc.MethodBindPtr
	fnGetSpaceState                        gdc.MethodBindPtr
}

var ptrsForPhysicsDirectBodyState3D ptrsForPhysicsDirectBodyState3DList

func initPhysicsDirectBodyState3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsDirectBodyState3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_total_gravity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetTotalGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_total_linear_damp")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetTotalLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_total_angular_damp")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetTotalAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_center_of_mass")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_center_of_mass_local")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetCenterOfMassLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_principal_inertia_axes")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetPrincipalInertiaAxes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2716978435))
	}
	{
		methodName := StringNameFromStr("get_inverse_mass")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetInverseMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_inverse_inertia")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetInverseInertia = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_inverse_inertia_tensor")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetInverseInertiaTensor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2716978435))
	}
	{
		methodName := StringNameFromStr("set_linear_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnSetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_linear_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_angular_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnSetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_angular_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_transform")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_transform")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("get_velocity_at_local_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetVelocityAtLocalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 192990374))
	}
	{
		methodName := StringNameFromStr("apply_central_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnApplyCentralImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2007698547))
	}
	{
		methodName := StringNameFromStr("apply_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnApplyImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2754756483))
	}
	{
		methodName := StringNameFromStr("apply_torque_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnApplyTorqueImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("apply_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnApplyCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2007698547))
	}
	{
		methodName := StringNameFromStr("apply_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnApplyForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2754756483))
	}
	{
		methodName := StringNameFromStr("apply_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnApplyTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("add_constant_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnAddConstantCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2007698547))
	}
	{
		methodName := StringNameFromStr("add_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnAddConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2754756483))
	}
	{
		methodName := StringNameFromStr("add_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnAddConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("set_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnSetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnSetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_sleep_state")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnSetSleepState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_sleeping")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnIsSleeping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_contact_count")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_contact_local_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactLocalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_contact_local_normal")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactLocalNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_contact_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_contact_local_shape")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactLocalShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_contact_local_velocity_at_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactLocalVelocityAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_contact_collider")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 495598643))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactColliderPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_id")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactColliderId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_object")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactColliderObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3332903315))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_shape")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_velocity_at_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetContactColliderVelocityAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_step")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("integrate_forces")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnIntegrateForces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_space_state")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState3D.fnGetSpaceState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2069328350))
	}

}

type PhysicsDirectBodyState3D struct {
	Object
}

func (me *PhysicsDirectBodyState3D) BaseClass() string {
	return "PhysicsDirectBodyState3D"
}

func NewPhysicsDirectBodyState3D() *PhysicsDirectBodyState3D {
	str := StringNameFromStr("PhysicsDirectBodyState3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsDirectBodyState3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsDirectBodyState3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsDirectBodyState3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsDirectBodyState3D) GetTotalGravity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetTotalGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetTotalLinearDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetTotalLinearDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetTotalAngularDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetTotalAngularDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetCenterOfMass() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetCenterOfMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetCenterOfMassLocal() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetCenterOfMassLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetPrincipalInertiaAxes() Basis {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBasis()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetPrincipalInertiaAxes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetInverseMass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetInverseMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetInverseInertia() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetInverseInertia), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetInverseInertiaTensor() Basis {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBasis()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetInverseInertiaTensor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) SetLinearVelocity(velocity Vector3) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnSetLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) GetLinearVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) SetAngularVelocity(velocity Vector3) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnSetAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) GetAngularVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) SetTransform(transform Transform3D) {
	cargs := []gdc.ConstTypePtr{transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) GetTransform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetVelocityAtLocalPosition(local_position Vector3) Vector3 {
	cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetVelocityAtLocalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) ApplyCentralImpulse(impulse Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnApplyCentralImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) ApplyImpulse(impulse Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnApplyImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) ApplyTorqueImpulse(impulse Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnApplyTorqueImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) ApplyCentralForce(force Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnApplyCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) ApplyForce(force Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnApplyForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) ApplyTorque(torque Vector3) {
	cargs := []gdc.ConstTypePtr{torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnApplyTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) AddConstantCentralForce(force Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnAddConstantCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) AddConstantForce(force Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnAddConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) AddConstantTorque(torque Vector3) {
	cargs := []gdc.ConstTypePtr{torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnAddConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) SetConstantForce(force Vector3) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnSetConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) GetConstantForce() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetConstantForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) SetConstantTorque(torque Vector3) {
	cargs := []gdc.ConstTypePtr{torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnSetConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) GetConstantTorque() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetConstantTorque), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) SetSleepState(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnSetSleepState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) IsSleeping() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnIsSleeping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetContactCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetContactLocalPosition(contact_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactLocalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetContactLocalNormal(contact_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactLocalNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetContactImpulse(contact_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactImpulse), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetContactLocalShape(contact_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactLocalShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetContactLocalVelocityAtPosition(contact_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactLocalVelocityAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetContactCollider(contact_idx int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetContactColliderPosition(contact_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactColliderPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetContactColliderId(contact_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactColliderId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetContactColliderObject(contact_idx int64) Object {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactColliderObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetContactColliderShape(contact_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) GetContactColliderVelocityAtPosition(contact_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetContactColliderVelocityAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState3D) GetStep() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState3D) IntegrateForces() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnIntegrateForces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState3D) GetSpaceState() PhysicsDirectSpaceState3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectSpaceState3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState3D.fnGetSpaceState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
