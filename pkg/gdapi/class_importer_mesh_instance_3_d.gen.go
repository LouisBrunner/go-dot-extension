// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ImporterMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *ImporterMeshInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImporterMeshInstance3D) BaseClass() string {
  return "ImporterMeshInstance3D"
}



// Enums

func (me *ImporterMeshInstance3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImporterMeshInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImporterMeshInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ImporterMeshInstance3D) SetMesh(mesh ImporterMesh, )  {
  classNameV := StringNameFromStr("ImporterMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2255166972) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImporterMeshInstance3D) GetMesh() ImporterMesh {
  classNameV := StringNameFromStr("ImporterMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3161779525) // FIXME: should cache?
  var ret ImporterMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ImporterMeshInstance3D) SetSkin(skin Skin, )  {
  classNameV := StringNameFromStr("ImporterMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3971435618) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImporterMeshInstance3D) GetSkin() Skin {
  classNameV := StringNameFromStr("ImporterMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2074563878) // FIXME: should cache?
  var ret Skin
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ImporterMeshInstance3D) SetSkeletonPath(skeleton_path NodePath, )  {
  classNameV := StringNameFromStr("ImporterMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeleton_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skeleton_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ImporterMeshInstance3D) GetSkeletonPath() NodePath {
  classNameV := StringNameFromStr("ImporterMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
