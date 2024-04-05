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

type ptrsForNode3DList struct {
  fnSetTransform gdc.MethodBindPtr
  fnGetTransform gdc.MethodBindPtr
  fnSetPosition gdc.MethodBindPtr
  fnGetPosition gdc.MethodBindPtr
  fnSetRotation gdc.MethodBindPtr
  fnGetRotation gdc.MethodBindPtr
  fnSetRotationDegrees gdc.MethodBindPtr
  fnGetRotationDegrees gdc.MethodBindPtr
  fnSetRotationOrder gdc.MethodBindPtr
  fnGetRotationOrder gdc.MethodBindPtr
  fnSetRotationEditMode gdc.MethodBindPtr
  fnGetRotationEditMode gdc.MethodBindPtr
  fnSetScale gdc.MethodBindPtr
  fnGetScale gdc.MethodBindPtr
  fnSetQuaternion gdc.MethodBindPtr
  fnGetQuaternion gdc.MethodBindPtr
  fnSetBasis gdc.MethodBindPtr
  fnGetBasis gdc.MethodBindPtr
  fnSetGlobalTransform gdc.MethodBindPtr
  fnGetGlobalTransform gdc.MethodBindPtr
  fnSetGlobalPosition gdc.MethodBindPtr
  fnGetGlobalPosition gdc.MethodBindPtr
  fnSetGlobalBasis gdc.MethodBindPtr
  fnGetGlobalBasis gdc.MethodBindPtr
  fnSetGlobalRotation gdc.MethodBindPtr
  fnGetGlobalRotation gdc.MethodBindPtr
  fnSetGlobalRotationDegrees gdc.MethodBindPtr
  fnGetGlobalRotationDegrees gdc.MethodBindPtr
  fnGetParentNode3D gdc.MethodBindPtr
  fnSetIgnoreTransformNotification gdc.MethodBindPtr
  fnSetAsTopLevel gdc.MethodBindPtr
  fnIsSetAsTopLevel gdc.MethodBindPtr
  fnSetDisableScale gdc.MethodBindPtr
  fnIsScaleDisabled gdc.MethodBindPtr
  fnGetWorld3D gdc.MethodBindPtr
  fnForceUpdateTransform gdc.MethodBindPtr
  fnSetVisibilityParent gdc.MethodBindPtr
  fnGetVisibilityParent gdc.MethodBindPtr
  fnUpdateGizmos gdc.MethodBindPtr
  fnAddGizmo gdc.MethodBindPtr
  fnGetGizmos gdc.MethodBindPtr
  fnClearGizmos gdc.MethodBindPtr
  fnSetSubgizmoSelection gdc.MethodBindPtr
  fnClearSubgizmoSelection gdc.MethodBindPtr
  fnSetVisible gdc.MethodBindPtr
  fnIsVisible gdc.MethodBindPtr
  fnIsVisibleInTree gdc.MethodBindPtr
  fnShow gdc.MethodBindPtr
  fnHide gdc.MethodBindPtr
  fnSetNotifyLocalTransform gdc.MethodBindPtr
  fnIsLocalTransformNotificationEnabled gdc.MethodBindPtr
  fnSetNotifyTransform gdc.MethodBindPtr
  fnIsTransformNotificationEnabled gdc.MethodBindPtr
  fnRotate gdc.MethodBindPtr
  fnGlobalRotate gdc.MethodBindPtr
  fnGlobalScale gdc.MethodBindPtr
  fnGlobalTranslate gdc.MethodBindPtr
  fnRotateObjectLocal gdc.MethodBindPtr
  fnScaleObjectLocal gdc.MethodBindPtr
  fnTranslateObjectLocal gdc.MethodBindPtr
  fnRotateX gdc.MethodBindPtr
  fnRotateY gdc.MethodBindPtr
  fnRotateZ gdc.MethodBindPtr
  fnTranslate gdc.MethodBindPtr
  fnOrthonormalize gdc.MethodBindPtr
  fnSetIdentity gdc.MethodBindPtr
  fnLookAt gdc.MethodBindPtr
  fnLookAtFromPosition gdc.MethodBindPtr
  fnToLocal gdc.MethodBindPtr
  fnToGlobal gdc.MethodBindPtr
}

var ptrsForNode3D ptrsForNode3DList

func initNode3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Node3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_transform")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
  }
  {
    methodName := StringNameFromStr("get_transform")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
  }
  {
    methodName := StringNameFromStr("set_position")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_position")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_rotation")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_rotation")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_rotation_degrees")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_rotation_degrees")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_rotation_order")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetRotationOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1820889989))
  }
  {
    methodName := StringNameFromStr("get_rotation_order")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetRotationOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 916939469))
  }
  {
    methodName := StringNameFromStr("set_rotation_edit_mode")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetRotationEditMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 141483330))
  }
  {
    methodName := StringNameFromStr("get_rotation_edit_mode")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetRotationEditMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1572188370))
  }
  {
    methodName := StringNameFromStr("set_scale")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_scale")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_quaternion")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetQuaternion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1727505552))
  }
  {
    methodName := StringNameFromStr("get_quaternion")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetQuaternion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1222331677))
  }
  {
    methodName := StringNameFromStr("set_basis")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1055510324))
  }
  {
    methodName := StringNameFromStr("get_basis")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2716978435))
  }
  {
    methodName := StringNameFromStr("set_global_transform")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetGlobalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
  }
  {
    methodName := StringNameFromStr("get_global_transform")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetGlobalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
  }
  {
    methodName := StringNameFromStr("set_global_position")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_global_position")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_global_basis")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetGlobalBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1055510324))
  }
  {
    methodName := StringNameFromStr("get_global_basis")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetGlobalBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2716978435))
  }
  {
    methodName := StringNameFromStr("set_global_rotation")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetGlobalRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_global_rotation")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetGlobalRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_global_rotation_degrees")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetGlobalRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_global_rotation_degrees")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetGlobalRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("get_parent_node_3d")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetParentNode3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 151077316))
  }
  {
    methodName := StringNameFromStr("set_ignore_transform_notification")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetIgnoreTransformNotification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_as_top_level")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetAsTopLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_set_as_top_level")
    defer methodName.Destroy()
    ptrsForNode3D.fnIsSetAsTopLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_disable_scale")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetDisableScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_scale_disabled")
    defer methodName.Destroy()
    ptrsForNode3D.fnIsScaleDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_world_3d")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetWorld3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 317588385))
  }
  {
    methodName := StringNameFromStr("force_update_transform")
    defer methodName.Destroy()
    ptrsForNode3D.fnForceUpdateTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_visibility_parent")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetVisibilityParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_visibility_parent")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetVisibilityParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("update_gizmos")
    defer methodName.Destroy()
    ptrsForNode3D.fnUpdateGizmos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("add_gizmo")
    defer methodName.Destroy()
    ptrsForNode3D.fnAddGizmo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1544533845))
  }
  {
    methodName := StringNameFromStr("get_gizmos")
    defer methodName.Destroy()
    ptrsForNode3D.fnGetGizmos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("clear_gizmos")
    defer methodName.Destroy()
    ptrsForNode3D.fnClearGizmos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_subgizmo_selection")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetSubgizmoSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3317607635))
  }
  {
    methodName := StringNameFromStr("clear_subgizmo_selection")
    defer methodName.Destroy()
    ptrsForNode3D.fnClearSubgizmoSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_visible")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_visible")
    defer methodName.Destroy()
    ptrsForNode3D.fnIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("is_visible_in_tree")
    defer methodName.Destroy()
    ptrsForNode3D.fnIsVisibleInTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("show")
    defer methodName.Destroy()
    ptrsForNode3D.fnShow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("hide")
    defer methodName.Destroy()
    ptrsForNode3D.fnHide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_notify_local_transform")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetNotifyLocalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_local_transform_notification_enabled")
    defer methodName.Destroy()
    ptrsForNode3D.fnIsLocalTransformNotificationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_notify_transform")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetNotifyTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_transform_notification_enabled")
    defer methodName.Destroy()
    ptrsForNode3D.fnIsTransformNotificationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("rotate")
    defer methodName.Destroy()
    ptrsForNode3D.fnRotate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3436291937))
  }
  {
    methodName := StringNameFromStr("global_rotate")
    defer methodName.Destroy()
    ptrsForNode3D.fnGlobalRotate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3436291937))
  }
  {
    methodName := StringNameFromStr("global_scale")
    defer methodName.Destroy()
    ptrsForNode3D.fnGlobalScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("global_translate")
    defer methodName.Destroy()
    ptrsForNode3D.fnGlobalTranslate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("rotate_object_local")
    defer methodName.Destroy()
    ptrsForNode3D.fnRotateObjectLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3436291937))
  }
  {
    methodName := StringNameFromStr("scale_object_local")
    defer methodName.Destroy()
    ptrsForNode3D.fnScaleObjectLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("translate_object_local")
    defer methodName.Destroy()
    ptrsForNode3D.fnTranslateObjectLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("rotate_x")
    defer methodName.Destroy()
    ptrsForNode3D.fnRotateX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("rotate_y")
    defer methodName.Destroy()
    ptrsForNode3D.fnRotateY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("rotate_z")
    defer methodName.Destroy()
    ptrsForNode3D.fnRotateZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("translate")
    defer methodName.Destroy()
    ptrsForNode3D.fnTranslate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("orthonormalize")
    defer methodName.Destroy()
    ptrsForNode3D.fnOrthonormalize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_identity")
    defer methodName.Destroy()
    ptrsForNode3D.fnSetIdentity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("look_at")
    defer methodName.Destroy()
    ptrsForNode3D.fnLookAt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2882425029))
  }
  {
    methodName := StringNameFromStr("look_at_from_position")
    defer methodName.Destroy()
    ptrsForNode3D.fnLookAtFromPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2086826090))
  }
  {
    methodName := StringNameFromStr("to_local")
    defer methodName.Destroy()
    ptrsForNode3D.fnToLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 192990374))
  }
  {
    methodName := StringNameFromStr("to_global")
    defer methodName.Destroy()
    ptrsForNode3D.fnToGlobal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 192990374))
  }
}

type Node3D struct {
  Node
}

func (me *Node3D) BaseClass() string {
  return "Node3D"
}

func NewNode3D() *Node3D {
  str := StringNameFromStr("Node3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Node3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  Node3DNotificationTransformChanged = 2000
  Node3DNotificationEnterWorld = 41
  Node3DNotificationExitWorld = 42
  Node3DNotificationVisibilityChanged = 43
  Node3DNotificationLocalTransformChanged = 44
)

// Enums

type Node3DRotationEditMode int
const (
  Node3DRotationEditModeRotationEditModeEuler Node3DRotationEditMode = 0
  Node3DRotationEditModeRotationEditModeQuaternion Node3DRotationEditMode = 1
  Node3DRotationEditModeRotationEditModeBasis Node3DRotationEditMode = 2
)

func (me *Node3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Node3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Node3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Node3D) SetTransform(local Transform3D, )  {
  cargs := []gdc.ConstTypePtr{local.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetTransform() Transform3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetPosition(position Vector3, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetRotation(euler_radians Vector3, )  {
  cargs := []gdc.ConstTypePtr{euler_radians.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotation() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetRotationDegrees(euler_degrees Vector3, )  {
  cargs := []gdc.ConstTypePtr{euler_degrees.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetRotationDegrees), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotationDegrees() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetRotationDegrees), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetRotationOrder(order EulerOrder, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetRotationOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotationOrder() EulerOrder {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret EulerOrder

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetRotationOrder), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Node3D) SetRotationEditMode(edit_mode Node3DRotationEditMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edit_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetRotationEditMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotationEditMode() Node3DRotationEditMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Node3DRotationEditMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetRotationEditMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Node3D) SetScale(scale Vector3, )  {
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetScale() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetQuaternion(quaternion Quaternion, )  {
  cargs := []gdc.ConstTypePtr{quaternion.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetQuaternion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetQuaternion() Quaternion {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewQuaternion()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetQuaternion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetBasis(basis Basis, )  {
  cargs := []gdc.ConstTypePtr{basis.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetBasis), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetBasis() Basis {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBasis()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetBasis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalTransform(global Transform3D, )  {
  cargs := []gdc.ConstTypePtr{global.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetGlobalTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalTransform() Transform3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetGlobalTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalPosition(position Vector3, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetGlobalPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetGlobalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalBasis(basis Basis, )  {
  cargs := []gdc.ConstTypePtr{basis.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetGlobalBasis), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalBasis() Basis {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBasis()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetGlobalBasis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalRotation(euler_radians Vector3, )  {
  cargs := []gdc.ConstTypePtr{euler_radians.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetGlobalRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalRotation() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetGlobalRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalRotationDegrees(euler_degrees Vector3, )  {
  cargs := []gdc.ConstTypePtr{euler_degrees.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetGlobalRotationDegrees), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalRotationDegrees() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetGlobalRotationDegrees), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) GetParentNode3D() Node3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetParentNode3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetIgnoreTransformNotification(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetIgnoreTransformNotification), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetAsTopLevel(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetAsTopLevel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsSetAsTopLevel() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnIsSetAsTopLevel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) SetDisableScale(disable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetDisableScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsScaleDisabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnIsScaleDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) GetWorld3D() World3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWorld3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetWorld3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) ForceUpdateTransform()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnForceUpdateTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetVisibilityParent(path NodePath, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetVisibilityParent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetVisibilityParent() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetVisibilityParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) UpdateGizmos()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnUpdateGizmos), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) AddGizmo(gizmo Node3DGizmo, )  {
  cargs := []gdc.ConstTypePtr{gizmo.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnAddGizmo), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGizmos() []Node3DGizmo {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGetGizmos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Node3DGizmo](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Node3D) ClearGizmos()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnClearGizmos), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetSubgizmoSelection(gizmo Node3DGizmo, id int64, transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{gizmo.AsCTypePtr(), gdc.ConstTypePtr(&id) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetSubgizmoSelection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) ClearSubgizmoSelection()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnClearSubgizmoSelection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetVisible(visible bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsVisible() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnIsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) IsVisibleInTree() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnIsVisibleInTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) Show()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnShow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) Hide()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnHide), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetNotifyLocalTransform(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetNotifyLocalTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsLocalTransformNotificationEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnIsLocalTransformNotificationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) SetNotifyTransform(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetNotifyTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsTransformNotificationEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnIsTransformNotificationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) Rotate(axis Vector3, angle float64, )  {
  cargs := []gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnRotate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GlobalRotate(axis Vector3, angle float64, )  {
  cargs := []gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGlobalRotate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GlobalScale(scale Vector3, )  {
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGlobalScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GlobalTranslate(offset Vector3, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnGlobalTranslate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateObjectLocal(axis Vector3, angle float64, )  {
  cargs := []gdc.ConstTypePtr{axis.AsCTypePtr(), gdc.ConstTypePtr(&angle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnRotateObjectLocal), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) ScaleObjectLocal(scale Vector3, )  {
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnScaleObjectLocal), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) TranslateObjectLocal(offset Vector3, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnTranslateObjectLocal), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateX(angle float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnRotateX), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateY(angle float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnRotateY), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateZ(angle float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnRotateZ), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) Translate(offset Vector3, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnTranslate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) Orthonormalize()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnOrthonormalize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetIdentity()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnSetIdentity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) LookAt(target Vector3, up Vector3, use_model_front bool, )  {
  cargs := []gdc.ConstTypePtr{target.AsCTypePtr(), up.AsCTypePtr(), gdc.ConstTypePtr(&use_model_front) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnLookAt), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) LookAtFromPosition(position Vector3, target Vector3, up Vector3, use_model_front bool, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), target.AsCTypePtr(), up.AsCTypePtr(), gdc.ConstTypePtr(&use_model_front) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnLookAtFromPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) ToLocal(global_point Vector3, ) Vector3 {
  cargs := []gdc.ConstTypePtr{global_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnToLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) ToGlobal(local_point Vector3, ) Vector3 {
  cargs := []gdc.ConstTypePtr{local_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode3D.fnToGlobal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type Node3DVisibilityChangedSignalFn func()

func (me *Node3D) ConnectVisibilityChanged(subs SignalSubscribers, fn Node3DVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node3D) DisconnectVisibilityChanged(subs SignalSubscribers, fn Node3DVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
