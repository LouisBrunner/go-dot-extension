// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VehicleWheel3D struct {
  obj gdc.ObjectPtr
}

func (me *VehicleWheel3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VehicleWheel3D) BaseClass() string {
  return "VehicleWheel3D"
}



// Enums

func (me *VehicleWheel3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VehicleWheel3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VehicleWheel3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VehicleWheel3D) SetRadius(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetRadius() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetSuspensionRestLength(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suspension_rest_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetSuspensionRestLength() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suspension_rest_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetSuspensionTravel(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suspension_travel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetSuspensionTravel() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suspension_travel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetSuspensionStiffness(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suspension_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetSuspensionStiffness() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suspension_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetSuspensionMaxForce(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suspension_max_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetSuspensionMaxForce() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suspension_max_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetDampingCompression(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_damping_compression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetDampingCompression() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_damping_compression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetDampingRelaxation(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_damping_relaxation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetDampingRelaxation() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_damping_relaxation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetUseAsTraction(enable bool, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_as_traction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) IsUsedAsTraction() bool {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_used_as_traction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetUseAsSteering(enable bool, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_as_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) IsUsedAsSteering() bool {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_used_as_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetFrictionSlip(length float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_friction_slip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetFrictionSlip() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_friction_slip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) IsInContact() bool {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_contact")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) GetContactBody() Node3D {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_contact_body")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 151077316) // FIXME: should cache?
  var ret Node3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetRollInfluence(roll_influence float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_roll_influence")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&roll_influence), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetRollInfluence() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_roll_influence")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) GetSkidinfo() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skidinfo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) GetRpm() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rpm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetEngineForce(engine_force float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_engine_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&engine_force), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetEngineForce() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_engine_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetBrake(brake float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_brake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brake), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetBrake() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_brake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleWheel3D) SetSteering(steering float32, )  {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&steering), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleWheel3D) GetSteering() float32 {
  classNameV := StringNameFromStr("VehicleWheel3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VehicleWheel3D) GetPropEngineForce() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropEngineForce(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropBrake() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropBrake(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropSteering() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropSteering(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropUseAsTraction() bool {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropUseAsTraction(value bool) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropUseAsSteering() bool {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropUseAsSteering(value bool) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropWheelRollInfluence() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropWheelRollInfluence(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropWheelRadius() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropWheelRadius(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropWheelRestLength() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropWheelRestLength(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropWheelFrictionSlip() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropWheelFrictionSlip(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropSuspensionTravel() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropSuspensionTravel(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropSuspensionStiffness() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropSuspensionStiffness(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropSuspensionMaxForce() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropSuspensionMaxForce(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropDampingCompression() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropDampingCompression(value float32) {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) GetPropDampingRelaxation() float32 {
  panic("TODO: implement")
}

func (me *VehicleWheel3D) SetPropDampingRelaxation(value float32) {
  panic("TODO: implement")
}