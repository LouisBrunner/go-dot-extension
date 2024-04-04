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

type PhysicsBody3D struct {
  CollisionObject3D
}

func (me *PhysicsBody3D) BaseClass() string {
  return "PhysicsBody3D"
}

func NewPhysicsBody3D() *PhysicsBody3D {
  str := StringNameFromStr("PhysicsBody3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsBody3D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *PhysicsBody3D) MoveAndCollide(motion Vector3, test_only bool, safe_margin float64, recovery_as_collision bool, max_collisions int64, ) KinematicCollision3D {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_and_collide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3208792678) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{motion.AsCTypePtr(), gdc.ConstTypePtr(&test_only) , gdc.ConstTypePtr(&safe_margin) , gdc.ConstTypePtr(&recovery_as_collision) , gdc.ConstTypePtr(&max_collisions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewKinematicCollision3D()
  pinner.Pin(&test_only)
  pinner.Pin(&safe_margin)
  pinner.Pin(&recovery_as_collision)
  pinner.Pin(&max_collisions)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsBody3D) TestMove(from Transform3D, motion Vector3, collision KinematicCollision3D, safe_margin float64, recovery_as_collision bool, max_collisions int64, ) bool {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("test_move")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2481691619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), motion.AsCTypePtr(), collision.AsCTypePtr(), gdc.ConstTypePtr(&safe_margin) , gdc.ConstTypePtr(&recovery_as_collision) , gdc.ConstTypePtr(&max_collisions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&safe_margin)
  pinner.Pin(&recovery_as_collision)
  pinner.Pin(&max_collisions)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsBody3D) SetAxisLock(axis PhysicsServer3DBodyAxis, lock bool, )  {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_axis_lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1787895195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis) , gdc.ConstTypePtr(&lock) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsBody3D) GetAxisLock(axis PhysicsServer3DBodyAxis, ) bool {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_axis_lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2264617709) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&axis) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&axis)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsBody3D) GetCollisionExceptions() []PhysicsBody3D {
  classNameV := StringNameFromStr("PhysicsBody3D")
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
  return ConvertArrayToSlice[PhysicsBody3D](ret)
}

func  (me *PhysicsBody3D) AddCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsBody3D) RemoveCollisionExceptionWith(body Node, )  {
  classNameV := StringNameFromStr("PhysicsBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_collision_exception_with")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
