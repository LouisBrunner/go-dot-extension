// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type EditorNode3DGizmoPlugin struct {
  Resource
}

func (me *EditorNode3DGizmoPlugin) BaseClass() string {
  return "EditorNode3DGizmoPlugin"
}

func NewEditorNode3DGizmoPlugin() *EditorNode3DGizmoPlugin {
  str := StringNameFromStr("EditorNode3DGizmoPlugin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorNode3DGizmoPlugin{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&billboard) , gdc.ConstTypePtr(&on_top) , gdc.ConstTypePtr(&use_vertex_color) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmoPlugin) CreateIconMaterial(name String, texture Texture2D, on_top bool, color Color, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_icon_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3804976916) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&on_top) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmoPlugin) CreateHandleMaterial(name String, billboard bool, texture Texture2D, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_handle_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2486475223) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&billboard) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmoPlugin) AddMaterial(name String, material StandardMaterial3D, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1374068695) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmoPlugin) GetMaterial(name String, gizmo EditorNode3DGizmo, ) StandardMaterial3D {
  classNameV := StringNameFromStr("EditorNode3DGizmoPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 974464017) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gizmo.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStandardMaterial3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
