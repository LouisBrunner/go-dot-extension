// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsTestMotionParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsTestMotionParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsTestMotionParameters3D) BaseClass() string {
  return "PhysicsTestMotionParameters3D"
}



// Enums

func (me *PhysicsTestMotionParameters3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsTestMotionParameters3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsTestMotionParameters3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsTestMotionParameters3D) GetFrom()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetFrom(from Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) GetMotion()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetMotion(motion Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) GetMargin()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) GetMaxCollisions()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetMaxCollisions(max_collisions int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) IsCollideSeparationRayEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetCollideSeparationRayEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) GetExcludeBodies()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetExcludeBodies(exclude_list RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) GetExcludeObjects()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetExcludeObjects(exclude_list int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) IsRecoveryAsCollisionEnabled()  {
  panic("TODO: implement")
}

func  (me *PhysicsTestMotionParameters3D) SetRecoveryAsCollisionEnabled(enabled bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
