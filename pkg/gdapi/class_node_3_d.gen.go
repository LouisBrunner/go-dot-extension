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

// TODO: needed?
// const (
// // )

var (
  Node3DNotificationTransformChanged = "2000" // TODO: construct correctly
  Node3DNotificationEnterWorld = "41" // TODO: construct correctly
  Node3DNotificationExitWorld = "42" // TODO: construct correctly
  Node3DNotificationVisibilityChanged = "43" // TODO: construct correctly
  Node3DNotificationLocalTransformChanged = "44" // TODO: construct correctly
)

type Node3DRotationEditMode int
const (
  Node3DRotationEditModeRotationEditModeEuler Node3DRotationEditMode = 0
  Node3DRotationEditModeRotationEditModeQuaternion Node3DRotationEditMode = 1
  Node3DRotationEditModeRotationEditModeBasis Node3DRotationEditMode = 2
)

func  (me *Node3D) SetTransform(local Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetTransform() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetPosition(position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetRotation(euler_radians Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetRotation() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetRotationDegrees(euler_degrees Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetRotationDegrees() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetRotationOrder(order EulerOrder, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetRotationOrder() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetRotationEditMode(edit_mode Node3DRotationEditMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetRotationEditMode() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetScale(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetScale() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetQuaternion(quaternion Quaternion, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetQuaternion() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetBasis(basis Basis, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetBasis() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetGlobalTransform(global Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetGlobalTransform() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetGlobalPosition(position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetGlobalPosition() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetGlobalRotation(euler_radians Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetGlobalRotation() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetGlobalRotationDegrees(euler_degrees Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetGlobalRotationDegrees() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetParentNode3D() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetIgnoreTransformNotification(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetAsTopLevel(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) IsSetAsTopLevel() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetDisableScale(disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) IsScaleDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetWorld3D() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) ForceUpdateTransform() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetVisibilityParent(path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetVisibilityParent() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) UpdateGizmos() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) AddGizmo(gizmo Node3DGizmo, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GetGizmos() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) ClearGizmos() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetSubgizmoSelection(gizmo Node3DGizmo, id int, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) ClearSubgizmoSelection() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) IsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) IsVisibleInTree() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) Show() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) Hide() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetNotifyLocalTransform(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) IsLocalTransformNotificationEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetNotifyTransform(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) IsTransformNotificationEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) Rotate(axis Vector3, angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GlobalRotate(axis Vector3, angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GlobalScale(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) GlobalTranslate(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) RotateObjectLocal(axis Vector3, angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) ScaleObjectLocal(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) TranslateObjectLocal(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) RotateX(angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) RotateY(angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) RotateZ(angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) Translate(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) Orthonormalize() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) SetIdentity() { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) LookAt(target Vector3, up Vector3, use_model_front bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) LookAtFromPosition(position Vector3, target Vector3, up Vector3, use_model_front bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) ToLocal(global_point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Node3D) ToGlobal(local_point Vector3, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
