// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *SoftBody3D) GetPhysicsRid()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetCollisionMask(collision_mask int, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetCollisionLayer(collision_layer int, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetCollisionLayer()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetCollisionLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetCollisionLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetParentCollisionIgnore(parent_collision_ignore NodePath, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetParentCollisionIgnore()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetDisableMode(mode SoftBody3DDisableMode, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetDisableMode()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetCollisionExceptions()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) AddCollisionExceptionWith(body Node, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) RemoveCollisionExceptionWith(body Node, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetSimulationPrecision(simulation_precision int, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetSimulationPrecision()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetTotalMass(mass float32, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetTotalMass()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetLinearStiffness(linear_stiffness float32, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetLinearStiffness()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetPressureCoefficient(pressure_coefficient float32, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetPressureCoefficient()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetDampingCoefficient(damping_coefficient float32, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetDampingCoefficient()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetDragCoefficient(drag_coefficient float32, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetDragCoefficient()  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) GetPointTransform(point_index int, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetPointPinned(point_index int, pinned bool, attachment_path NodePath, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) IsPointPinned(point_index int, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) SetRayPickable(ray_pickable bool, )  {
  panic("TODO: implement")
}

func  (me *SoftBody3D) IsRayPickable()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
