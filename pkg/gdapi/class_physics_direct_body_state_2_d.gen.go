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

type ptrsForPhysicsDirectBodyState2DList struct {
	fnGetTotalGravity                      gdc.MethodBindPtr
	fnGetTotalLinearDamp                   gdc.MethodBindPtr
	fnGetTotalAngularDamp                  gdc.MethodBindPtr
	fnGetCenterOfMass                      gdc.MethodBindPtr
	fnGetCenterOfMassLocal                 gdc.MethodBindPtr
	fnGetInverseMass                       gdc.MethodBindPtr
	fnGetInverseInertia                    gdc.MethodBindPtr
	fnSetLinearVelocity                    gdc.MethodBindPtr
	fnGetLinearVelocity                    gdc.MethodBindPtr
	fnSetAngularVelocity                   gdc.MethodBindPtr
	fnGetAngularVelocity                   gdc.MethodBindPtr
	fnSetTransform                         gdc.MethodBindPtr
	fnGetTransform                         gdc.MethodBindPtr
	fnGetVelocityAtLocalPosition           gdc.MethodBindPtr
	fnApplyCentralImpulse                  gdc.MethodBindPtr
	fnApplyTorqueImpulse                   gdc.MethodBindPtr
	fnApplyImpulse                         gdc.MethodBindPtr
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
	fnGetContactLocalShape                 gdc.MethodBindPtr
	fnGetContactLocalVelocityAtPosition    gdc.MethodBindPtr
	fnGetContactCollider                   gdc.MethodBindPtr
	fnGetContactColliderPosition           gdc.MethodBindPtr
	fnGetContactColliderId                 gdc.MethodBindPtr
	fnGetContactColliderObject             gdc.MethodBindPtr
	fnGetContactColliderShape              gdc.MethodBindPtr
	fnGetContactColliderVelocityAtPosition gdc.MethodBindPtr
	fnGetContactImpulse                    gdc.MethodBindPtr
	fnGetStep                              gdc.MethodBindPtr
	fnIntegrateForces                      gdc.MethodBindPtr
	fnGetSpaceState                        gdc.MethodBindPtr
}

var ptrsForPhysicsDirectBodyState2D ptrsForPhysicsDirectBodyState2DList

func initPhysicsDirectBodyState2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsDirectBodyState2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_total_gravity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetTotalGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_total_linear_damp")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetTotalLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_total_angular_damp")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetTotalAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_center_of_mass")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_center_of_mass_local")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetCenterOfMassLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_inverse_mass")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetInverseMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_inverse_inertia")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetInverseInertia = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_linear_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnSetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_linear_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_angular_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnSetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_angular_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_transform")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("get_transform")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_velocity_at_local_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetVelocityAtLocalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2656412154))
	}
	{
		methodName := StringNameFromStr("apply_central_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnApplyCentralImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("apply_torque_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnApplyTorqueImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("apply_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnApplyImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288681949))
	}
	{
		methodName := StringNameFromStr("apply_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnApplyCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3862383994))
	}
	{
		methodName := StringNameFromStr("apply_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnApplyForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288681949))
	}
	{
		methodName := StringNameFromStr("apply_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnApplyTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("add_constant_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnAddConstantCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3862383994))
	}
	{
		methodName := StringNameFromStr("add_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnAddConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288681949))
	}
	{
		methodName := StringNameFromStr("add_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnAddConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnSetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnSetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sleep_state")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnSetSleepState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_sleeping")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnIsSleeping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_contact_count")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_contact_local_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactLocalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("get_contact_local_normal")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactLocalNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("get_contact_local_shape")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactLocalShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_contact_local_velocity_at_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactLocalVelocityAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("get_contact_collider")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 495598643))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactColliderPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_id")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactColliderId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_object")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactColliderObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3332903315))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_shape")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_contact_collider_velocity_at_position")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactColliderVelocityAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("get_contact_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetContactImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("get_step")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("integrate_forces")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnIntegrateForces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_space_state")
		defer methodName.Destroy()
		ptrsForPhysicsDirectBodyState2D.fnGetSpaceState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2506717822))
	}
}

type PhysicsDirectBodyState2D struct {
	Object
}

func (me *PhysicsDirectBodyState2D) BaseClass() string {
	return "PhysicsDirectBodyState2D"
}

func NewPhysicsDirectBodyState2D() *PhysicsDirectBodyState2D {
	str := StringNameFromStr("PhysicsDirectBodyState2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsDirectBodyState2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsDirectBodyState2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsDirectBodyState2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsDirectBodyState2D) GetTotalGravity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetTotalGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetTotalLinearDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetTotalLinearDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetTotalAngularDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetTotalAngularDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetCenterOfMass() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetCenterOfMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetCenterOfMassLocal() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetCenterOfMassLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetInverseMass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetInverseMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetInverseInertia() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetInverseInertia), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) SetLinearVelocity(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnSetLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) GetLinearVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) SetAngularVelocity(velocity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&velocity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnSetAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) GetAngularVelocity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) SetTransform(transform Transform2D) {
	cargs := []gdc.ConstTypePtr{transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) GetTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetVelocityAtLocalPosition(local_position Vector2) Vector2 {
	cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetVelocityAtLocalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) ApplyCentralImpulse(impulse Vector2) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnApplyCentralImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) ApplyTorqueImpulse(impulse float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&impulse)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnApplyTorqueImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) ApplyImpulse(impulse Vector2, position Vector2) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnApplyImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) ApplyCentralForce(force Vector2) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnApplyCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) ApplyForce(force Vector2, position Vector2) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnApplyForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) ApplyTorque(torque float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnApplyTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) AddConstantCentralForce(force Vector2) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnAddConstantCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) AddConstantForce(force Vector2, position Vector2) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnAddConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) AddConstantTorque(torque float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnAddConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) SetConstantForce(force Vector2) {
	cargs := []gdc.ConstTypePtr{force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnSetConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) GetConstantForce() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetConstantForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) SetConstantTorque(torque float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnSetConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) GetConstantTorque() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetConstantTorque), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) SetSleepState(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnSetSleepState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) IsSleeping() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnIsSleeping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetContactCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetContactLocalPosition(contact_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactLocalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetContactLocalNormal(contact_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactLocalNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetContactLocalShape(contact_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactLocalShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetContactLocalVelocityAtPosition(contact_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactLocalVelocityAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetContactCollider(contact_idx int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetContactColliderPosition(contact_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactColliderPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetContactColliderId(contact_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactColliderId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetContactColliderObject(contact_idx int64) Object {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactColliderObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetContactColliderShape(contact_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) GetContactColliderVelocityAtPosition(contact_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactColliderVelocityAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetContactImpulse(contact_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&contact_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetContactImpulse), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsDirectBodyState2D) GetStep() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsDirectBodyState2D) IntegrateForces() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnIntegrateForces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsDirectBodyState2D) GetSpaceState() PhysicsDirectSpaceState2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectSpaceState2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectBodyState2D.fnGetSpaceState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
