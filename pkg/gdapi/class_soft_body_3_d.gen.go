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

type SoftBody3DDisableMode int
const (
  SoftBody3DDisableModeDisableModeRemove SoftBody3DDisableMode = 0
  SoftBody3DDisableModeDisableModeKeepActive SoftBody3DDisableMode = 1
)

func  (me *SoftBody3D) GetPhysicsRid() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetCollisionMask(collision_mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetCollisionMask() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetCollisionLayer(collision_layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetCollisionLayer() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetCollisionMaskValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetCollisionMaskValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetCollisionLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetCollisionLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetParentCollisionIgnore(parent_collision_ignore NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetParentCollisionIgnore() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetDisableMode(mode SoftBody3DDisableMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetDisableMode() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetCollisionExceptions() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) AddCollisionExceptionWith(body Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) RemoveCollisionExceptionWith(body Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetSimulationPrecision(simulation_precision int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetSimulationPrecision() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetTotalMass(mass float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetTotalMass() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetLinearStiffness(linear_stiffness float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetLinearStiffness() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetPressureCoefficient(pressure_coefficient float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetPressureCoefficient() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetDampingCoefficient(damping_coefficient float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetDampingCoefficient() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetDragCoefficient(drag_coefficient float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetDragCoefficient() { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) GetPointTransform(point_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetPointPinned(point_index int, pinned bool, attachment_path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) IsPointPinned(point_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) SetRayPickable(ray_pickable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SoftBody3D) IsRayPickable() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
