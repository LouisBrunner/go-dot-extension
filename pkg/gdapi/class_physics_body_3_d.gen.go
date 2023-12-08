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

func  (me *PhysicsBody3D) MoveAndCollide(motion Vector3, test_only bool, safe_margin float32, recovery_as_collision bool, max_collisions int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody3D) TestMove(from Transform3D, motion Vector3, collision KinematicCollision3D, safe_margin float32, recovery_as_collision bool, max_collisions int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody3D) SetAxisLock(axis PhysicsServer3DBodyAxis, lock bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody3D) GetAxisLock(axis PhysicsServer3DBodyAxis, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody3D) GetCollisionExceptions() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody3D) AddCollisionExceptionWith(body Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody3D) RemoveCollisionExceptionWith(body Node, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
