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

type ptrsForShapeCast3DList struct {
  fnResourceChanged gdc.MethodBindPtr
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
  fnSetDebugShapeCustomColor gdc.MethodBindPtr
  fnGetDebugShapeCustomColor gdc.MethodBindPtr
}

var ptrsForShapeCast3D ptrsForShapeCast3DList

func initShapeCast3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ShapeCast3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("resource_changed")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnResourceChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 968641751))
  }
  {
    methodName := StringNameFromStr("set_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_shape")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549710052))
  }
  {
    methodName := StringNameFromStr("get_shape")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3214262478))
  }
  {
    methodName := StringNameFromStr("set_target_position")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_target_position")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_margin")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_margin")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_results")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetMaxResults = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_results")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetMaxResults = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("is_colliding")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnIsColliding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_collision_count")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetCollisionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("force_shapecast_update")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnForceShapecastUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_collider")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3332903315))
  }
  {
    methodName := StringNameFromStr("get_collider_rid")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 495598643))
  }
  {
    methodName := StringNameFromStr("get_collider_shape")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("get_collision_point")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetCollisionPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
  }
  {
    methodName := StringNameFromStr("get_collision_normal")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetCollisionNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
  }
  {
    methodName := StringNameFromStr("get_closest_collision_safe_fraction")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetClosestCollisionSafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_closest_collision_unsafe_fraction")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetClosestCollisionUnsafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("add_exception_rid")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnAddExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("add_exception")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnAddException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1976431078))
  }
  {
    methodName := StringNameFromStr("remove_exception_rid")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnRemoveExceptionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("remove_exception")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnRemoveException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1976431078))
  }
  {
    methodName := StringNameFromStr("clear_exceptions")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnClearExceptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_collision_mask")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_collision_mask")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_collision_mask_value")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_collision_mask_value")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_exclude_parent_body")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_exclude_parent_body")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetExcludeParentBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collide_with_areas")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_collide_with_areas_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collide_with_bodies")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_collide_with_bodies_enabled")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_debug_shape_custom_color")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnSetDebugShapeCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_debug_shape_custom_color")
    defer methodName.Destroy()
    ptrsForShapeCast3D.fnGetDebugShapeCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
}

type ShapeCast3D struct {
  Node3D
}

func (me *ShapeCast3D) BaseClass() string {
  return "ShapeCast3D"
}

func NewShapeCast3D() *ShapeCast3D {
  str := StringNameFromStr("ShapeCast3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ShapeCast3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ShapeCast3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ShapeCast3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShapeCast3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ShapeCast3D) ResourceChanged(resource Resource, )  {
  cargs := []gdc.ConstTypePtr{resource.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnResourceChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) SetEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) IsEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) SetShape(shape Shape3D, )  {
  cargs := []gdc.ConstTypePtr{shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetShape() Shape3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShape3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast3D) SetTargetPosition(local_point Vector3, )  {
  cargs := []gdc.ConstTypePtr{local_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetTargetPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast3D) SetMargin(margin float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetMargin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) SetMaxResults(max_results int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetMaxResults), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetMaxResults() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetMaxResults), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) IsColliding() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnIsColliding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) GetCollisionCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetCollisionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) ForceShapecastUpdate()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnForceShapecastUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetCollider(index int64, ) Object {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast3D) GetColliderRid(index int64, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast3D) GetColliderShape(index int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) GetCollisionPoint(index int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetCollisionPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast3D) GetCollisionNormal(index int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetCollisionNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShapeCast3D) GetClosestCollisionSafeFraction() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetClosestCollisionSafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) GetClosestCollisionUnsafeFraction() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetClosestCollisionUnsafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) AddExceptionRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnAddExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) AddException(node CollisionObject3D, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnAddException), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) RemoveExceptionRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnRemoveExceptionRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) RemoveException(node CollisionObject3D, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnRemoveException), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) ClearExceptions()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnClearExceptions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) SetCollisionMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetCollisionMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) SetCollisionMaskValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetCollisionMaskValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) SetExcludeParentBody(mask bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetExcludeParentBody), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetExcludeParentBody() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetExcludeParentBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) SetCollideWithAreas(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) IsCollideWithAreasEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) SetCollideWithBodies(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) IsCollideWithBodiesEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ShapeCast3D) SetDebugShapeCustomColor(debug_shape_custom_color Color, )  {
  cargs := []gdc.ConstTypePtr{debug_shape_custom_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnSetDebugShapeCustomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShapeCast3D) GetDebugShapeCustomColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShapeCast3D.fnGetDebugShapeCustomColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
