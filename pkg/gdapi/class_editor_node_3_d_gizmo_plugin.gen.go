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

func  (me *EditorNode3DGizmoPlugin) XHasGizmo(for_node_3d Node3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XCreateGizmo(for_node_3d Node3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XGetGizmoName() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XGetPriority() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XCanBeHidden() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XIsSelectableWhenHidden() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XRedraw(gizmo EditorNode3DGizmo, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XGetHandleName(gizmo EditorNode3DGizmo, handle_id int, secondary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XIsHandleHighlighted(gizmo EditorNode3DGizmo, handle_id int, secondary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XGetHandleValue(gizmo EditorNode3DGizmo, handle_id int, secondary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XSetHandle(gizmo EditorNode3DGizmo, handle_id int, secondary bool, camera Camera3D, screen_pos Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XCommitHandle(gizmo EditorNode3DGizmo, handle_id int, secondary bool, restore Variant, cancel bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XSubgizmosIntersectRay(gizmo EditorNode3DGizmo, camera Camera3D, screen_pos Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XSubgizmosIntersectFrustum(gizmo EditorNode3DGizmo, camera Camera3D, frustum_planes Plane, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XGetSubgizmoTransform(gizmo EditorNode3DGizmo, subgizmo_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XSetSubgizmoTransform(gizmo EditorNode3DGizmo, subgizmo_id int, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) XCommitSubgizmos(gizmo EditorNode3DGizmo, ids PackedInt32Array, restores Transform3D, cancel bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) CreateMaterial(name String, color Color, billboard bool, on_top bool, use_vertex_color bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) CreateIconMaterial(name String, texture Texture2D, on_top bool, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) CreateHandleMaterial(name String, billboard bool, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) AddMaterial(name String, material StandardMaterial3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmoPlugin) GetMaterial(name String, gizmo EditorNode3DGizmo, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
