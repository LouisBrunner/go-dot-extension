// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *PhysicsBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsBody3D) MoveAndCollide(motion Vector3, test_only bool, safe_margin float32, recovery_as_collision bool, max_collisions int, ) KinematicCollision3D {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_and_collide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2825704414) // FIXME: should cache?
  var ret KinematicCollision3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(motion.AsCTypePtr()), gdc.ConstTypePtr(&test_only), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision), gdc.ConstTypePtr(&max_collisions), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsBody3D) TestMove(from Transform3D, motion Vector3, collision KinematicCollision3D, safe_margin float32, recovery_as_collision bool, max_collisions int, ) bool {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("test_move")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680299713) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(motion.AsCTypePtr()), gdc.ConstTypePtr(collision.AsCTypePtr()), gdc.ConstTypePtr(&safe_margin), gdc.ConstTypePtr(&recovery_as_collision), gdc.ConstTypePtr(&max_collisions), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsBody3D) SetAxisLock(axis PhysicsServer3DBodyAxis, lock bool, )  {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis_lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1787895195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&lock), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsBody3D) GetAxisLock(axis PhysicsServer3DBodyAxis, ) bool {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis_lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2264617709) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsBody3D) GetCollisionExceptions() PhysicsBody3D {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_exceptions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret PhysicsBody3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsBody3D) AddCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsBody3D) RemoveCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *PhysicsBody3D) GetPropAxisLockLinearX() bool {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) SetPropAxisLockLinearX(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) GetPropAxisLockLinearY() bool {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) SetPropAxisLockLinearY(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) GetPropAxisLockLinearZ() bool {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) SetPropAxisLockLinearZ(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) GetPropAxisLockAngularX() bool {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) SetPropAxisLockAngularX(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) GetPropAxisLockAngularY() bool {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) SetPropAxisLockAngularY(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) GetPropAxisLockAngularZ() bool {
  panic("TODO: implement")
}

func (me *PhysicsBody3D) SetPropAxisLockAngularZ(value bool) {
  panic("TODO: implement")
}