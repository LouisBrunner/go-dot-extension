// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicalBone3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicalBone3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicalBone3D) BaseClass() string {
  return "PhysicalBone3D"
}



// Enums

type PhysicalBone3DDampMode int
const (
  PhysicalBone3DDampModeDampModeCombine PhysicalBone3DDampMode = 0
  PhysicalBone3DDampModeDampModeReplace PhysicalBone3DDampMode = 1
)

type PhysicalBone3DJointType int
const (
  PhysicalBone3DJointTypeJointTypeNone PhysicalBone3DJointType = 0
  PhysicalBone3DJointTypeJointTypePin PhysicalBone3DJointType = 1
  PhysicalBone3DJointTypeJointTypeCone PhysicalBone3DJointType = 2
  PhysicalBone3DJointTypeJointTypeHinge PhysicalBone3DJointType = 3
  PhysicalBone3DJointTypeJointTypeSlider PhysicalBone3DJointType = 4
  PhysicalBone3DJointTypeJointType6Dof PhysicalBone3DJointType = 5
)

func (me *PhysicalBone3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicalBone3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicalBone3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicalBone3D) ApplyCentralImpulse(impulse Vector3, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_central_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(impulse.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) ApplyImpulse(impulse Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1002852006) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(impulse.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) SetJointType(joint_type PhysicalBone3DJointType, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2289552604) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetJointType() PhysicalBone3DJointType {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 931347320) // FIXME: should cache?
  var ret PhysicalBone3DJointType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetJointOffset(offset Transform3D, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetJointOffset() Transform3D {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetJointRotation(euler Vector3, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(euler.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetJointRotation() Vector3 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetBodyOffset(offset Transform3D, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_body_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetBodyOffset() Transform3D {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_body_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) GetSimulatePhysics() bool {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_simulate_physics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) IsSimulatingPhysics() bool {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_simulating_physics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) GetBoneId() int {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetMass(mass float32, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetMass() float32 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetFriction(friction float32, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetFriction() float32 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetBounce(bounce float32, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetBounce() float32 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetGravityScale(gravity_scale float32, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetGravityScale() float32 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetLinearDampMode(linear_damp_mode PhysicalBone3DDampMode, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1244972221) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetLinearDampMode() PhysicalBone3DDampMode {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205884699) // FIXME: should cache?
  var ret PhysicalBone3DDampMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetAngularDampMode(angular_damp_mode PhysicalBone3DDampMode, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1244972221) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetAngularDampMode() PhysicalBone3DDampMode {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 205884699) // FIXME: should cache?
  var ret PhysicalBone3DDampMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetLinearDamp(linear_damp float32, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetLinearDamp() float32 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetAngularDamp(angular_damp float32, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetAngularDamp() float32 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_damp")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetLinearVelocity(linear_velocity Vector3, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(linear_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetLinearVelocity() Vector3 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetAngularVelocity(angular_velocity Vector3, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(angular_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) GetAngularVelocity() Vector3 {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetUseCustomIntegrator(enable bool, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_custom_integrator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) IsUsingCustomIntegrator() bool {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_custom_integrator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone3D) SetCanSleep(able_to_sleep bool, )  {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_can_sleep")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&able_to_sleep), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone3D) IsAbleToSleep() bool {
  classNameV := StringNameFromStr("PhysicalBone3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_able_to_sleep")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PhysicalBone3D) GetPropJointType() int {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropJointType(value int) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropJointOffset() Transform3D {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropJointOffset(value Transform3D) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropJointRotation() Vector3 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropJointRotation(value Vector3) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropBodyOffset() Transform3D {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropBodyOffset(value Transform3D) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropMass() float32 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropMass(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropFriction() float32 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropFriction(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropBounce() float32 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropBounce(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropGravityScale() float32 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropGravityScale(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropCustomIntegrator() bool {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropCustomIntegrator(value bool) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropLinearDampMode() int {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropLinearDampMode(value int) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropLinearDamp() float32 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropLinearDamp(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropAngularDampMode() int {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropAngularDampMode(value int) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropAngularDamp() float32 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropAngularDamp(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropLinearVelocity() Vector3 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropLinearVelocity(value Vector3) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropAngularVelocity() Vector3 {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropAngularVelocity(value Vector3) {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) GetPropCanSleep() bool {
  panic("TODO: implement")
}

func (me *PhysicalBone3D) SetPropCanSleep(value bool) {
  panic("TODO: implement")
}