// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type PhysicsBody2D struct {
  CollisionObject2D
}

func (me *PhysicsBody2D) BaseClass() string {
  return "PhysicsBody2D"
}

func NewPhysicsBody2D() *PhysicsBody2D {
  str := StringNameFromStr("PhysicsBody2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsBody2D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *PhysicsBody2D) MoveAndCollide(motion Vector2, test_only bool, safe_margin float64, recovery_as_collision bool, ) KinematicCollision2D {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_and_collide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3681923724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{motion.AsCTypePtr(), gdc.ConstTypePtr(&test_only) , gdc.ConstTypePtr(&safe_margin) , gdc.ConstTypePtr(&recovery_as_collision) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewKinematicCollision2D()
  pinner.Pin(&test_only)
  pinner.Pin(&safe_margin)
  pinner.Pin(&recovery_as_collision)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsBody2D) TestMove(from Transform2D, motion Vector2, collision KinematicCollision2D, safe_margin float64, recovery_as_collision bool, ) bool {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("test_move")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3324464701) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), motion.AsCTypePtr(), collision.AsCTypePtr(), gdc.ConstTypePtr(&safe_margin) , gdc.ConstTypePtr(&recovery_as_collision) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&safe_margin)
  pinner.Pin(&recovery_as_collision)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsBody2D) GetCollisionExceptions() []PhysicsBody2D {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_exceptions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[PhysicsBody2D](ret)
}

func  (me *PhysicsBody2D) AddCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsBody2D) RemoveCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
