// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFSkeleton struct {
  obj gdc.ObjectPtr
}

func (me *GLTFSkeleton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFSkeleton) BaseClass() string {
  return "GLTFSkeleton"
}



// Enums

func (me *GLTFSkeleton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFSkeleton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFSkeleton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFSkeleton) GetJoints() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joints")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkeleton) SetJoints(joints PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joints")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joints.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkeleton) GetRoots() PackedInt32Array {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_roots")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkeleton) SetRoots(roots PackedInt32Array, )  {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_roots")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(roots.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkeleton) GetGodotSkeleton() Skeleton3D {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_godot_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1814733083) // FIXME: should cache?
  var ret Skeleton3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkeleton) GetUniqueNames() String {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unique_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkeleton) SetUniqueNames(unique_names String, )  {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unique_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(unique_names.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkeleton) GetGodotBoneNode() Dictionary {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_godot_bone_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkeleton) SetGodotBoneNode(godot_bone_node Dictionary, )  {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_godot_bone_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(godot_bone_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFSkeleton) GetBoneAttachmentCount() int {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_attachment_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFSkeleton) GetBoneAttachment(idx int, ) BoneAttachment3D {
  classNameV := StringNameFromStr("GLTFSkeleton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_attachment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 945440495) // FIXME: should cache?
  var ret BoneAttachment3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
