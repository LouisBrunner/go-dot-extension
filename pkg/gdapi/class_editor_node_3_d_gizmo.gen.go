// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorNode3DGizmo struct {
  obj gdc.ObjectPtr
}

func (me *EditorNode3DGizmo) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorNode3DGizmo) BaseClass() string {
  return "EditorNode3DGizmo"
}

func  (me *EditorNode3DGizmo) XRedraw() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XGetHandleName(id int, secondary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XIsHandleHighlighted(id int, secondary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XGetHandleValue(id int, secondary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XSetHandle(id int, secondary bool, camera Camera3D, point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XCommitHandle(id int, secondary bool, restore Variant, cancel bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XSubgizmosIntersectRay(camera Camera3D, point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XSubgizmosIntersectFrustum(camera Camera3D, frustum Plane, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XSetSubgizmoTransform(id int, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XGetSubgizmoTransform(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) XCommitSubgizmos(ids PackedInt32Array, restores Transform3D, cancel bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) AddLines(lines PackedVector3Array, material Material, billboard bool, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) AddMesh(mesh Mesh, material Material, transform Transform3D, skeleton SkinReference, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) AddCollisionSegments(segments PackedVector3Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) AddCollisionTriangles(triangles TriangleMesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) AddUnscaledBillboard(material Material, default_scale float32, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) AddHandles(handles PackedVector3Array, material Material, ids PackedInt32Array, billboard bool, secondary bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) SetNode3D(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) GetNode3D() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) GetPlugin() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) SetHidden(hidden bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) IsSubgizmoSelected(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorNode3DGizmo) GetSubgizmoSelection() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
