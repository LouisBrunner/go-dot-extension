// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *EditorNode3DGizmoPlugin) XHasGizmo(for_node_3d Node3D, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XCreateGizmo(for_node_3d Node3D, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XGetGizmoName()  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XGetPriority()  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XCanBeHidden()  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XIsSelectableWhenHidden()  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XRedraw(gizmo EditorNode3DGizmo, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XGetHandleName(gizmo EditorNode3DGizmo, handle_id int, secondary bool, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XIsHandleHighlighted(gizmo EditorNode3DGizmo, handle_id int, secondary bool, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XGetHandleValue(gizmo EditorNode3DGizmo, handle_id int, secondary bool, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XSetHandle(gizmo EditorNode3DGizmo, handle_id int, secondary bool, camera Camera3D, screen_pos Vector2, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XCommitHandle(gizmo EditorNode3DGizmo, handle_id int, secondary bool, restore Variant, cancel bool, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XSubgizmosIntersectRay(gizmo EditorNode3DGizmo, camera Camera3D, screen_pos Vector2, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XSubgizmosIntersectFrustum(gizmo EditorNode3DGizmo, camera Camera3D, frustum_planes Plane, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XGetSubgizmoTransform(gizmo EditorNode3DGizmo, subgizmo_id int, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XSetSubgizmoTransform(gizmo EditorNode3DGizmo, subgizmo_id int, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) XCommitSubgizmos(gizmo EditorNode3DGizmo, ids PackedInt32Array, restores Transform3D, cancel bool, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) CreateMaterial(name String, color Color, billboard bool, on_top bool, use_vertex_color bool, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) CreateIconMaterial(name String, texture Texture2D, on_top bool, color Color, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) CreateHandleMaterial(name String, billboard bool, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) AddMaterial(name String, material StandardMaterial3D, )  {
  panic("TODO: implement")
}

func  (me *EditorNode3DGizmoPlugin) GetMaterial(name String, gizmo EditorNode3DGizmo, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
