// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node3D struct {
  obj gdc.ObjectPtr
}

func (me *Node3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Node3D) BaseClass() string {
  return "Node3D"
}



// Constants

var (
  Node3DNotificationTransformChanged = "2000" // TODO: construct correctly
  Node3DNotificationEnterWorld = "41" // TODO: construct correctly
  Node3DNotificationExitWorld = "42" // TODO: construct correctly
  Node3DNotificationVisibilityChanged = "43" // TODO: construct correctly
  Node3DNotificationLocalTransformChanged = "44" // TODO: construct correctly
)

// Enums

type Node3DRotationEditMode int
const (
  Node3DRotationEditModeRotationEditModeEuler Node3DRotationEditMode = 0
  Node3DRotationEditModeRotationEditModeQuaternion Node3DRotationEditMode = 1
  Node3DRotationEditModeRotationEditModeBasis Node3DRotationEditMode = 2
)

func (me *Node3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Node3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Node3D) SetTransform(local Transform3D, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetTransform()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetPosition(position Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetPosition()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetRotation(euler_radians Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetRotation()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetRotationDegrees(euler_degrees Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetRotationDegrees()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetRotationOrder(order EulerOrder, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetRotationOrder()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetRotationEditMode(edit_mode Node3DRotationEditMode, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetRotationEditMode()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetScale(scale Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetScale()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetQuaternion(quaternion Quaternion, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetQuaternion()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetBasis(basis Basis, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetBasis()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetGlobalTransform(global Transform3D, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetGlobalTransform()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetGlobalPosition(position Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetGlobalPosition()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetGlobalRotation(euler_radians Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetGlobalRotation()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetGlobalRotationDegrees(euler_degrees Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetGlobalRotationDegrees()  {
  panic("TODO: implement")
}

func  (me *Node3D) GetParentNode3D()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetIgnoreTransformNotification(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) SetAsTopLevel(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) IsSetAsTopLevel()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetDisableScale(disable bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) IsScaleDisabled()  {
  panic("TODO: implement")
}

func  (me *Node3D) GetWorld3D()  {
  panic("TODO: implement")
}

func  (me *Node3D) ForceUpdateTransform()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetVisibilityParent(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetVisibilityParent()  {
  panic("TODO: implement")
}

func  (me *Node3D) UpdateGizmos()  {
  panic("TODO: implement")
}

func  (me *Node3D) AddGizmo(gizmo Node3DGizmo, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GetGizmos()  {
  panic("TODO: implement")
}

func  (me *Node3D) ClearGizmos()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetSubgizmoSelection(gizmo Node3DGizmo, id int, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *Node3D) ClearSubgizmoSelection()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) IsVisible()  {
  panic("TODO: implement")
}

func  (me *Node3D) IsVisibleInTree()  {
  panic("TODO: implement")
}

func  (me *Node3D) Show()  {
  panic("TODO: implement")
}

func  (me *Node3D) Hide()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetNotifyLocalTransform(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) IsLocalTransformNotificationEnabled()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetNotifyTransform(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) IsTransformNotificationEnabled()  {
  panic("TODO: implement")
}

func  (me *Node3D) Rotate(axis Vector3, angle float32, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GlobalRotate(axis Vector3, angle float32, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GlobalScale(scale Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) GlobalTranslate(offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) RotateObjectLocal(axis Vector3, angle float32, )  {
  panic("TODO: implement")
}

func  (me *Node3D) ScaleObjectLocal(scale Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) TranslateObjectLocal(offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) RotateX(angle float32, )  {
  panic("TODO: implement")
}

func  (me *Node3D) RotateY(angle float32, )  {
  panic("TODO: implement")
}

func  (me *Node3D) RotateZ(angle float32, )  {
  panic("TODO: implement")
}

func  (me *Node3D) Translate(offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) Orthonormalize()  {
  panic("TODO: implement")
}

func  (me *Node3D) SetIdentity()  {
  panic("TODO: implement")
}

func  (me *Node3D) LookAt(target Vector3, up Vector3, use_model_front bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) LookAtFromPosition(position Vector3, target Vector3, up Vector3, use_model_front bool, )  {
  panic("TODO: implement")
}

func  (me *Node3D) ToLocal(global_point Vector3, )  {
  panic("TODO: implement")
}

func  (me *Node3D) ToGlobal(local_point Vector3, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
