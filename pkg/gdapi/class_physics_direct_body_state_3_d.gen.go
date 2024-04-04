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

func  (me *PhysicsDirectBodyState3D) GetTotalGravity() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetTotalLinearDamp() float64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsDirectBodyState3D) GetTotalAngularDamp() float64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsDirectBodyState3D) GetCenterOfMass() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) GetCenterOfMassLocal() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetPrincipalInertiaAxes() Basis {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_principal_inertia_axes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2716978435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBasis()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetInverseMass() float64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inverse_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsDirectBodyState3D) GetInverseInertia() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inverse_inertia")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetInverseInertiaTensor() Basis {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) SetLinearVelocity(velocity Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) GetLinearVelocity() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) SetAngularVelocity(velocity Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) GetAngularVelocity() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) SetTransform(transform Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) GetTransform() Transform3D {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetVelocityAtLocalPosition(local_position Vector3, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity_at_local_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 192990374) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) ApplyCentralImpulse(impulse Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2007698547) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) ApplyImpulse(impulse Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2754756483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) ApplyTorqueImpulse(impulse Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) ApplyCentralForce(force Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2007698547) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) ApplyForce(force Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2754756483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) ApplyTorque(torque Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{torque.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) AddConstantCentralForce(force Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2007698547) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) AddConstantForce(force Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2754756483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) AddConstantTorque(torque Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{torque.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) SetConstantForce(force Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) GetConstantForce() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) SetConstantTorque(torque Vector3, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{torque.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) GetConstantTorque() Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) SetSleepState(enabled bool, )  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sleep_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) IsSleeping() bool {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) GetContactCount() int64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
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

func  (me *PhysicsDirectBodyState3D) GetContactLocalPosition(contact_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetContactLocalNormal(contact_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetContactImpulse(contact_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetContactLocalShape(contact_idx int64, ) int64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsDirectBodyState3D) GetContactLocalVelocityAtPosition(contact_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_local_velocity_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetContactCollider(contact_idx int64, ) RID {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 495598643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderPosition(contact_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderId(contact_idx int64, ) int64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderObject(contact_idx int64, ) Object {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3332903315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderShape(contact_idx int64, ) int64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderVelocityAtPosition(contact_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_collider_velocity_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&contact_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&contact_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectBodyState3D) GetStep() float64 {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsDirectBodyState3D) IntegrateForces()  {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("integrate_forces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsDirectBodyState3D) GetSpaceState() PhysicsDirectSpaceState3D {
  classNameV := StringNameFromStr("PhysicsDirectBodyState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_space_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2069328350) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPhysicsDirectSpaceState3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
