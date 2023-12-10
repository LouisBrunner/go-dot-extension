// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsBody2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsBody2D) BaseClass() string {
  return "PhysicsBody2D"
}



// Enums

func (me *PhysicsBody2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsBody2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsBody2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsBody2D) MoveAndCollide(motion Vector2, test_only bool, safe_margin float32, recovery_as_collision bool, ) KinematicCollision2D {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_and_collide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1529961754) // FIXME: should cache?
  var ret KinematicCollision2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(motion.AsCTypePtr()), gdc.ConstTypePtr(&test_only), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsBody2D) TestMove(from Transform2D, motion Vector2, collision KinematicCollision2D, safe_margin float32, recovery_as_collision bool, ) bool {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("test_move")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1369208982) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(motion.AsCTypePtr()), gdc.ConstTypePtr(collision.AsCTypePtr()), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsBody2D) GetCollisionExceptions() PhysicsBody2D {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_exceptions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret PhysicsBody2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsBody2D) AddCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsBody2D) RemoveCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties