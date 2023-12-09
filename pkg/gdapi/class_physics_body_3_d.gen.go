// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsBody3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsBody3D) BaseClass() string {
  return "PhysicsBody3D"
}



// Enums

func (me *PhysicsBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsBody3D) MoveAndCollide(motion Vector3, test_only bool, safe_margin float32, recovery_as_collision bool, max_collisions int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsBody3D) TestMove(from Transform3D, motion Vector3, collision KinematicCollision3D, safe_margin float32, recovery_as_collision bool, max_collisions int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsBody3D) SetAxisLock(axis PhysicsServer3DBodyAxis, lock bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsBody3D) GetAxisLock(axis PhysicsServer3DBodyAxis, )  {
  panic("TODO: implement")
}

func  (me *PhysicsBody3D) GetCollisionExceptions()  {
  panic("TODO: implement")
}

func  (me *PhysicsBody3D) AddCollisionExceptionWith(body Node, )  {
  panic("TODO: implement")
}

func  (me *PhysicsBody3D) RemoveCollisionExceptionWith(body Node, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
