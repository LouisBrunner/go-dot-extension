// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForShapeCast2DList struct {
  fnSetEnabled gdc.MethodBindPtr
  fnIsEnabled gdc.MethodBindPtr
  fnSetShape gdc.MethodBindPtr
  fnGetShape gdc.MethodBindPtr
  fnSetTargetPosition gdc.MethodBindPtr
  fnGetTargetPosition gdc.MethodBindPtr
  fnSetMargin gdc.MethodBindPtr
  fnGetMargin gdc.MethodBindPtr
  fnSetMaxResults gdc.MethodBindPtr
  fnGetMaxResults gdc.MethodBindPtr
  fnIsColliding gdc.MethodBindPtr
  fnGetCollisionCount gdc.MethodBindPtr
  fnForceShapecastUpdate gdc.MethodBindPtr
  fnGetCollider gdc.MethodBindPtr
  fnGetColliderRid gdc.MethodBindPtr
  fnGetColliderShape gdc.MethodBindPtr
  fnGetCollisionPoint gdc.MethodBindPtr
  fnGetCollisionNormal gdc.MethodBindPtr
  fnGetClosestCollisionSafeFraction gdc.MethodBindPtr
  fnGetClosestCollisionUnsafeFraction gdc.MethodBindPtr
  fnAddExceptionRid gdc.MethodBindPtr
  fnAddException gdc.MethodBindPtr
  fnRemoveExceptionRid gdc.MethodBindPtr
  fnRemoveException gdc.MethodBindPtr
  fnClearExceptions gdc.MethodBindPtr
  fnSetCollisionMask gdc.MethodBindPtr
  fnGetCollisionMask gdc.MethodBindPtr
  fnSetCollisionMaskValue gdc.MethodBindPtr
  fnGetCollisionMaskValue gdc.MethodBindPtr
  fnSetExcludeParentBody gdc.MethodBindPtr
  fnGetExcludeParentBody gdc.MethodBindPtr
  fnSetCollideWithAreas gdc.MethodBindPtr
  fnIsCollideWithAreasEnabled gdc.MethodBindPtr
  fnSetCollideWithBodies gdc.MethodBindPtr
  fnIsCollideWithBodiesEnabled gdc.MethodBindPtr
}

var ptrsForShapeCast2D ptrsForShapeCast2DList

func initShapeCast2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ShapeCast2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_shape")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 771364740))
  }
  {
    methodName := StringNameFromStr("get_shape")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 522005891))
  }
  {
    methodName := StringNameFromStr("set_target_position")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_target_position")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_margin")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_margin")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_results")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetMaxResults = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_results")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetMaxResults = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("is_colliding")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnIsColliding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_collision_count")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetCollisionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("force_shapecast_update")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnForceShapecastUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_collider")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3332903315))
  }
  {
    methodName := StringNameFromStr("get_collider_rid")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 495598643))
  }
  {
    methodName := StringNameFromStr("get_collider_shape")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("get_collision_point")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetCollisionPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
  }
  {
    methodName := StringNameFromStr("get_collision_normal")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetCollisionNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
  }
  {
    methodName := StringNameFromStr("get_closest_collision_safe_fraction")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetClosestCollisionSafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_closest_collision_unsafe_fraction")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetClosestCollisionUnsafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("add_exception_rid")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnAddExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("add_exception")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnAddException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3090941106))
  }
  {
    methodName := StringNameFromStr("remove_exception_rid")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnRemoveExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("remove_exception")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnRemoveException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3090941106))
  }
  {
    methodName := StringNameFromStr("clear_exceptions")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnClearExceptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_collision_mask")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_collision_mask")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_collision_mask_value")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_collision_mask_value")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_exclude_parent_body")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_exclude_parent_body")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnGetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collide_with_areas")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_collide_with_areas_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collide_with_bodies")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_collide_with_bodies_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast2D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type ShapeCast2D struct {
  Node2D
}

func (me *ShapeCast2D) BaseClass() string {
  return "ShapeCast2D"
}

func NewShapeCast2D() *ShapeCast2D {
  str := StringNameFromStr("ShapeCast2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ShapeCast2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ShapeCast2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ShapeCast2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShapeCast2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ShapeCast2D) SetEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) IsEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) SetShape(shape Shape2D, )  {
  cargs := []gdc.ConstTypePtr{shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetShape() Shape2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShape2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast2D) SetTargetPosition(local_point Vector2, )  {
  cargs := []gdc.ConstTypePtr{local_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetTargetPosition() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast2D) SetMargin(margin float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetMargin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) SetMaxResults(max_results int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetMaxResults), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetMaxResults() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetMaxResults), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) IsColliding() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnIsColliding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) GetCollisionCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetCollisionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) ForceShapecastUpdate()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnForceShapecastUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetCollider(index int64, ) Object {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast2D) GetColliderRid(index int64, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast2D) GetColliderShape(index int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) GetCollisionPoint(index int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetCollisionPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast2D) GetCollisionNormal(index int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetCollisionNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast2D) GetClosestCollisionSafeFraction() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetClosestCollisionSafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) GetClosestCollisionUnsafeFraction() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetClosestCollisionUnsafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) AddExceptionRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnAddExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) AddException(node CollisionObject2D, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnAddException), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) RemoveExceptionRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnRemoveExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) RemoveException(node CollisionObject2D, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnRemoveException), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) ClearExceptions()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnClearExceptions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) SetCollisionMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetCollisionMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) SetCollisionMaskValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetCollisionMaskValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) SetExcludeParentBody(mask bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetExcludeParentBody), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) GetExcludeParentBody() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnGetExcludeParentBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) SetCollideWithAreas(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) IsCollideWithAreasEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast2D) SetCollideWithBodies(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast2D) IsCollideWithBodiesEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast2D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
