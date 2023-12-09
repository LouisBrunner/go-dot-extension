// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *GLTFSkeleton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFSkeleton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GLTFSkeleton) GetJoints()  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) SetJoints(joints PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) GetRoots()  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) SetRoots(roots PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) GetGodotSkeleton()  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) GetUniqueNames()  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) SetUniqueNames(unique_names String, )  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) GetGodotBoneNode()  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) SetGodotBoneNode(godot_bone_node Dictionary, )  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) GetBoneAttachmentCount()  {
  panic("TODO: implement")
}

func  (me *GLTFSkeleton) GetBoneAttachment(idx int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
