// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsBody2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsBody2D) BaseClass() string {
  return "PhysicsBody2D"
}

func  (me *PhysicsBody2D) MoveAndCollide(motion Vector2, test_only bool, safe_margin float32, recovery_as_collision bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody2D) TestMove(from Transform2D, motion Vector2, collision KinematicCollision2D, safe_margin float32, recovery_as_collision bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody2D) GetCollisionExceptions() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody2D) AddCollisionExceptionWith(body Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsBody2D) RemoveCollisionExceptionWith(body Node, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
