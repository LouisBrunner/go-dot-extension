// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorNode3DGizmoPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorNode3DGizmoPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorNode3DGizmoPlugin) BaseClass() string {
  return "EditorNode3DGizmoPlugin"
}



// Enums

func (me *EditorNode3DGizmoPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorNode3DGizmoPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorNode3DGizmoPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorNode3DGizmoPlugin) CreateMaterial(name String, color Color, billboard bool, on_top bool, use_vertex_color bool, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3486012546) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&billboard), gdc.ConstTypePtr(&on_top), gdc.ConstTypePtr(&use_vertex_color), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmoPlugin) CreateIconMaterial(name String, texture Texture2D, on_top bool, color Color, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_icon_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3804976916) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(&on_top), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmoPlugin) CreateHandleMaterial(name String, billboard bool, texture Texture2D, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_handle_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2486475223) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&billboard), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmoPlugin) AddMaterial(name String, material StandardMaterial3D, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1374068695) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmoPlugin) GetMaterial(name String, gizmo EditorNode3DGizmo, ) StandardMaterial3D {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 974464017) // FIXME: should cache?
  var ret StandardMaterial3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(gizmo.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
