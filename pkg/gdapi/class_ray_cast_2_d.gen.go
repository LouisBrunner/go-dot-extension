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

type ptrsForRayCast2DList struct {
  fnSetEnabled gdc.MethodBindPtr
  fnIsEnabled gdc.MethodBindPtr
  fnSetTargetPosition gdc.MethodBindPtr
  fnGetTargetPosition gdc.MethodBindPtr
  fnIsColliding gdc.MethodBindPtr
  fnForceRaycastUpdate gdc.MethodBindPtr
  fnGetCollider gdc.MethodBindPtr
  fnGetColliderRid gdc.MethodBindPtr
  fnGetColliderShape gdc.MethodBindPtr
  fnGetCollisionPoint gdc.MethodBindPtr
  fnGetCollisionNormal gdc.MethodBindPtr
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
  fnSetHitFromInside gdc.MethodBindPtr
  fnIsHitFromInsideEnabled gdc.MethodBindPtr
}

var ptrsForRayCast2D ptrsForRayCast2DList

func initRayCast2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RayCast2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_enabled")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_enabled")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_target_position")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_target_position")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("is_colliding")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnIsColliding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("force_raycast_update")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnForceRaycastUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_collider")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1981248198))
  }
  {
    methodName := StringNameFromStr("get_collider_rid")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("get_collider_shape")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_collision_point")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetCollisionPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("get_collision_normal")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetCollisionNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("add_exception_rid")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnAddExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("add_exception")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnAddException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3090941106))
  }
  {
    methodName := StringNameFromStr("remove_exception_rid")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnRemoveExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("remove_exception")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnRemoveException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3090941106))
  }
  {
    methodName := StringNameFromStr("clear_exceptions")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnClearExceptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_collision_mask")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_collision_mask")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_collision_mask_value")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_collision_mask_value")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_exclude_parent_body")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_exclude_parent_body")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnGetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collide_with_areas")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_collide_with_areas_enabled")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collide_with_bodies")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_collide_with_bodies_enabled")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_hit_from_inside")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnSetHitFromInside = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_hit_from_inside_enabled")
    defer methodName.Destroy()
    ptrsForRayCast2D.fnIsHitFromInsideEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type RayCast2D struct {
  Node2D
}

func (me *RayCast2D) BaseClass() string {
  return "RayCast2D"
}

func NewRayCast2D() *RayCast2D {
  str := StringNameFromStr("RayCast2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RayCast2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RayCast2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RayCast2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RayCast2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RayCast2D) SetEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) IsEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) SetTargetPosition(local_point Vector2, )  {
  cargs := []gdc.ConstTypePtr{local_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) GetTargetPosition() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RayCast2D) IsColliding() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnIsColliding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) ForceRaycastUpdate()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnForceRaycastUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) GetCollider() Object {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RayCast2D) GetColliderRid() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RayCast2D) GetColliderShape() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) GetCollisionPoint() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetCollisionPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RayCast2D) GetCollisionNormal() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetCollisionNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RayCast2D) AddExceptionRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnAddExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) AddException(node CollisionObject2D, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnAddException), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) RemoveExceptionRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnRemoveExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) RemoveException(node CollisionObject2D, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnRemoveException), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) ClearExceptions()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnClearExceptions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) SetCollisionMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) GetCollisionMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) SetCollisionMaskValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) GetCollisionMaskValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) SetExcludeParentBody(mask bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetExcludeParentBody), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) GetExcludeParentBody() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnGetExcludeParentBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) SetCollideWithAreas(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) IsCollideWithAreasEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) SetCollideWithBodies(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) IsCollideWithBodiesEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RayCast2D) SetHitFromInside(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnSetHitFromInside), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RayCast2D) IsHitFromInsideEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRayCast2D.fnIsHitFromInsideEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
