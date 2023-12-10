// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SoftBody3D struct {
  obj gdc.ObjectPtr
}

func (me *SoftBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SoftBody3D) BaseClass() string {
  return "SoftBody3D"
}



// Enums

type SoftBody3DDisableMode int
const (
  SoftBody3DDisableModeDisableModeRemove SoftBody3DDisableMode = 0
  SoftBody3DDisableModeDisableModeKeepActive SoftBody3DDisableMode = 1
)

func (me *SoftBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SoftBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SoftBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SoftBody3D) GetPhysicsRid() RID {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetCollisionMask(collision_mask int, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetCollisionMask() int {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetCollisionLayer(collision_layer int, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetCollisionLayer() int {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetCollisionLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetCollisionLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetParentCollisionIgnore(parent_collision_ignore NodePath, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parent_collision_ignore")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parent_collision_ignore.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetParentCollisionIgnore() NodePath {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent_collision_ignore")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetDisableMode(mode SoftBody3DDisableMode, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1104158384) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetDisableMode() SoftBody3DDisableMode {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_disable_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4135042476) // FIXME: should cache?
  var ret SoftBody3DDisableMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) GetCollisionExceptions() PhysicsBody3D {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_exceptions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret PhysicsBody3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) AddCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) RemoveCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) SetSimulationPrecision(simulation_precision int, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_simulation_precision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&simulation_precision), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetSimulationPrecision() int {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_simulation_precision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetTotalMass(mass float32, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_total_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetTotalMass() float32 {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetLinearStiffness(linear_stiffness float32, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_stiffness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetLinearStiffness() float32 {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetPressureCoefficient(pressure_coefficient float32, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressure_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure_coefficient), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetPressureCoefficient() float32 {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pressure_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetDampingCoefficient(damping_coefficient float32, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_damping_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&damping_coefficient), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetDampingCoefficient() float32 {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_damping_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetDragCoefficient(drag_coefficient float32, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&drag_coefficient), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) GetDragCoefficient() float32 {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drag_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) GetPointTransform(point_index int, ) Vector3 {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 871989493) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetPointPinned(point_index int, pinned bool, attachment_path NodePath, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_pinned")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814935226) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_index), gdc.ConstTypePtr(&pinned), gdc.ConstTypePtr(attachment_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) IsPointPinned(point_index int, ) bool {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_point_pinned")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SoftBody3D) SetRayPickable(ray_pickable bool, )  {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ray_pickable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SoftBody3D) IsRayPickable() bool {
  classNameV := StringNameFromStr("SoftBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SoftBody3D) GetPropCollisionLayer() int {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropCollisionLayer(value int) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropCollisionMask() int {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropCollisionMask(value int) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropParentCollisionIgnore() NodePath {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropParentCollisionIgnore(value NodePath) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropSimulationPrecision() int {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropSimulationPrecision(value int) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropTotalMass() float32 {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropTotalMass(value float32) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropLinearStiffness() float32 {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropLinearStiffness(value float32) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropPressureCoefficient() float32 {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropPressureCoefficient(value float32) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropDampingCoefficient() float32 {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropDampingCoefficient(value float32) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropDragCoefficient() float32 {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropDragCoefficient(value float32) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropRayPickable() bool {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropRayPickable(value bool) {
  panic("TODO: implement")
}

func (me *SoftBody3D) GetPropDisableMode() int {
  panic("TODO: implement")
}

func (me *SoftBody3D) SetPropDisableMode(value int) {
  panic("TODO: implement")
}