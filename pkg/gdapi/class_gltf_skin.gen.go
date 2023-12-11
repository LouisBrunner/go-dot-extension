// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFSkin struct {
  obj gdc.ObjectPtr
}

func (me *GLTFSkin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFSkin) BaseClass() string {
  return "GLTFSkin"
}



// Enums

func (me *GLTFSkin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFSkin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFSkin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFSkin) GetSkinRoot() int {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skin_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetSkinRoot(skin_root int, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skin_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skin_root), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetJointsOriginal() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joints_original")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetJointsOriginal(joints_original PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joints_original")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joints_original.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetInverseBinds() Transform3D {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inverse_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetInverseBinds(inverse_binds Transform3D, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inverse_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(inverse_binds.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetJoints() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joints")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetJoints(joints PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joints")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joints.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetNonJoints() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_non_joints")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetNonJoints(non_joints PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_non_joints")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(non_joints.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetRoots() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_roots")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetRoots(roots PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_roots")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(roots.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetSkeleton() int {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetSkeleton(skeleton int, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skeleton), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetJointIToBoneI() Dictionary {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_i_to_bone_i")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetJointIToBoneI(joint_i_to_bone_i Dictionary, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_i_to_bone_i")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint_i_to_bone_i.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetJointIToName() Dictionary {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_i_to_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetJointIToName(joint_i_to_name Dictionary, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_i_to_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint_i_to_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkin) GetGodotSkin() Skin {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_godot_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1032037385) // FIXME: should cache?
  var ret Skin
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkin) SetGodotSkin(godot_skin Skin, )  {
  classNameV := StringNameFromStr("GLTFSkin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_godot_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3971435618) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(godot_skin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
