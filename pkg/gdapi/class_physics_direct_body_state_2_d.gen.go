// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsDirectBodyState2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState2D) BaseClass() string {
  return "PhysicsDirectBodyState2D"
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

func  (me *PhysicsDirectBodyState2D) GetTotalGravity() Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetTotalLinearDamp() float32 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetTotalAngularDamp() float32 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetCenterOfMass() Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetCenterOfMassLocal() Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetInverseMass() float32 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inverse_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetInverseInertia() float32 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inverse_inertia")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) SetLinearVelocity(velocity Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) GetLinearVelocity() Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) SetAngularVelocity(velocity float32, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&velocity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) GetAngularVelocity() float32 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) SetTransform(transform Transform2D, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) GetTransform() Transform2D {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetVelocityAtLocalPosition(local_position Vector2, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_at_local_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) ApplyCentralImpulse(impulse Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(impulse.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) ApplyTorqueImpulse(impulse float32, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&impulse), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) ApplyImpulse(impulse Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 496058220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(impulse.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) ApplyCentralForce(force Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3862383994) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) ApplyForce(force Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 496058220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) ApplyTorque(torque float32, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) AddConstantCentralForce(force Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3862383994) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) AddConstantForce(force Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 496058220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) AddConstantTorque(torque float32, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) SetConstantForce(force Vector2, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) GetConstantForce() Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) SetConstantTorque(torque float32, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) GetConstantTorque() float32 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) SetSleepState(enabled bool, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sleep_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) IsSleeping() bool {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sleeping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactCount() int {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalPosition(contact_idx int, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalNormal(contact_idx int, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalShape(contact_idx int, ) int {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalVelocityAtPosition(contact_idx int, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_velocity_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactCollider(contact_idx int, ) RID {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderPosition(contact_idx int, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderId(contact_idx int, ) int {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderObject(contact_idx int, ) Object {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3332903315) // FIXME: should cache?
  var ret Object
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderShape(contact_idx int, ) int {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderVelocityAtPosition(contact_idx int, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_velocity_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetContactImpulse(contact_idx int, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) GetStep() float32 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectBodyState2D) IntegrateForces()  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("integrate_forces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsDirectBodyState2D) GetSpaceState() PhysicsDirectSpaceState2D {
  classNameV := StringNameFromStr("PhysicsDirectBodyState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_space_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2506717822) // FIXME: should cache?
  var ret PhysicsDirectSpaceState2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PhysicsDirectBodyState2D) GetPropStep() float32 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropStep(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropInverseMass() float32 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropInverseMass(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropInverseInertia() float32 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropInverseInertia(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropTotalAngularDamp() float32 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropTotalAngularDamp(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropTotalLinearDamp() float32 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropTotalLinearDamp(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropTotalGravity() Vector2 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropTotalGravity(value Vector2) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropCenterOfMass() Vector2 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropCenterOfMass(value Vector2) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropCenterOfMassLocal() Vector2 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropCenterOfMassLocal(value Vector2) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropAngularVelocity() float32 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropAngularVelocity(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropLinearVelocity() Vector2 {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropLinearVelocity(value Vector2) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropSleeping() bool {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropSleeping(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) GetPropTransform() Transform2D {
  panic("TODO: implement")
}

func (me *PhysicsDirectBodyState2D) SetPropTransform(value Transform2D) {
  panic("TODO: implement")
}