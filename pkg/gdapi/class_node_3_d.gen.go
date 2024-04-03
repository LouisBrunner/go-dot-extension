// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetTransform() Transform3D {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetPosition(position Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetPosition() Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetRotation(euler_radians Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(euler_radians.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotation() Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetRotationDegrees(euler_degrees Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(euler_degrees.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotationDegrees() Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetRotationOrder(order EulerOrder, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1820889989) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotationOrder() EulerOrder {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 916939469) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret EulerOrder

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Node3D) SetRotationEditMode(edit_mode Node3DRotationEditMode, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_edit_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 141483330) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edit_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetRotationEditMode() Node3DRotationEditMode {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_edit_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1572188370) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret Node3DRotationEditMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Node3D) SetScale(scale Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetScale() Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetQuaternion(quaternion Quaternion, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_quaternion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1727505552) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(quaternion.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetQuaternion() Quaternion {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_quaternion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1222331677) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewQuaternion()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetBasis(basis Basis, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_basis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1055510324) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(basis.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetBasis() Basis {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_basis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2716978435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBasis()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalTransform(global Transform3D, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(global.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalTransform() Transform3D {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalPosition(position Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalPosition() Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalBasis(basis Basis, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_basis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1055510324) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(basis.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalBasis() Basis {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_basis")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2716978435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBasis()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalRotation(euler_radians Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(euler_radians.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalRotation() Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetGlobalRotationDegrees(euler_degrees Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(euler_degrees.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGlobalRotationDegrees() Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) GetParentNode3D() Node3D {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent_node_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 151077316) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNode3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) SetIgnoreTransformNotification(enabled bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ignore_transform_notification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetAsTopLevel(enable bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_top_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsSetAsTopLevel() bool {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_set_as_top_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) SetDisableScale(disable bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disable_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsScaleDisabled() bool {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_scale_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) GetWorld3D() World3D {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 317588385) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewWorld3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) ForceUpdateTransform()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetVisibilityParent(path NodePath, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetVisibilityParent() NodePath {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) UpdateGizmos()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_gizmos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) AddGizmo(gizmo Node3DGizmo, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_gizmo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1544533845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(gizmo.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GetGizmos() []Node3DGizmo {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gizmos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Node3DGizmo](ret)
}

func  (me *Node3D) ClearGizmos()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_gizmos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetSubgizmoSelection(gizmo Node3DGizmo, id int64, transform Transform3D, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_subgizmo_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3317607635) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(gizmo.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(transform.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) ClearSubgizmoSelection()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_subgizmo_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetVisible(visible bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsVisible() bool {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) IsVisibleInTree() bool {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible_in_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) Show()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("show")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) Hide()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetNotifyLocalTransform(enable bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_notify_local_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsLocalTransformNotificationEnabled() bool {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_local_transform_notification_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) SetNotifyTransform(enable bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_notify_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) IsTransformNotificationEnabled() bool {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_transform_notification_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node3D) Rotate(axis Vector3, angle float64, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3436291937) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(axis.AsCTypePtr()), gdc.ConstTypePtr(&angle), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GlobalRotate(axis Vector3, angle float64, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_rotate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3436291937) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(axis.AsCTypePtr()), gdc.ConstTypePtr(&angle), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GlobalScale(scale Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) GlobalTranslate(offset Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_translate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateObjectLocal(axis Vector3, angle float64, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate_object_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3436291937) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(axis.AsCTypePtr()), gdc.ConstTypePtr(&angle), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) ScaleObjectLocal(scale Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scale_object_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) TranslateObjectLocal(offset Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("translate_object_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateX(angle float64, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateY(angle float64, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) RotateZ(angle float64, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate_z")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) Translate(offset Vector3, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("translate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) Orthonormalize()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("orthonormalize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) SetIdentity()  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_identity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) LookAt(target Vector3, up Vector3, use_model_front bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("look_at")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2882425029) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(target.AsCTypePtr()), gdc.ConstTypePtr(up.AsCTypePtr()), gdc.ConstTypePtr(&use_model_front), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) LookAtFromPosition(position Vector3, target Vector3, up Vector3, use_model_front bool, )  {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("look_at_from_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2086826090) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(target.AsCTypePtr()), gdc.ConstTypePtr(up.AsCTypePtr()), gdc.ConstTypePtr(&use_model_front), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node3D) ToLocal(global_point Vector3, ) Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 192990374) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(global_point.AsCTypePtr()), }
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node3D) ToGlobal(local_point Vector3, ) Vector3 {
  classNameV := StringNameFromStr("Node3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_global")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 192990374) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_point.AsCTypePtr()), }
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
