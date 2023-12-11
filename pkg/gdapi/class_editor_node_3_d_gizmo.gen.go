// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorNode3DGizmo struct {
  obj gdc.ObjectPtr
}

func (me *EditorNode3DGizmo) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorNode3DGizmo) BaseClass() string {
  return "EditorNode3DGizmo"
}



// Enums

func (me *EditorNode3DGizmo) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorNode3DGizmo) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorNode3DGizmo) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorNode3DGizmo) AddLines(lines PackedVector3Array, material Material, billboard bool, modulate Color, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 302451090) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(lines.AsCTypePtr()), gdc.ConstTypePtr(material.AsCTypePtr()), gdc.ConstTypePtr(&billboard), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) AddMesh(mesh Mesh, material Material, transform Transform3D, skeleton SkinReference, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1868867708) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), gdc.ConstTypePtr(material.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(skeleton.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) AddCollisionSegments(segments PackedVector3Array, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_segments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(segments.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) AddCollisionTriangles(triangles TriangleMesh, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_collision_triangles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 54901064) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(triangles.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) AddUnscaledBillboard(material Material, default_scale float32, modulate Color, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_unscaled_billboard")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3719733075) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), gdc.ConstTypePtr(&default_scale), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) AddHandles(handles PackedVector3Array, material Material, ids PackedInt32Array, billboard bool, secondary bool, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_handles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2254560097) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(handles.AsCTypePtr()), gdc.ConstTypePtr(material.AsCTypePtr()), gdc.ConstTypePtr(ids.AsCTypePtr()), gdc.ConstTypePtr(&billboard), gdc.ConstTypePtr(&secondary), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) SetNode3D(node Node, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) GetNode3D() Node3D {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_3d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 151077316) // FIXME: should cache?
  var ret Node3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorNode3DGizmo) GetPlugin() EditorNode3DGizmoPlugin {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4250544552) // FIXME: should cache?
  var ret EditorNode3DGizmoPlugin
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorNode3DGizmo) Clear()  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) SetHidden(hidden bool, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorNode3DGizmo) IsSubgizmoSelected(id int, ) bool {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_subgizmo_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorNode3DGizmo) GetSubgizmoSelection() PackedInt32Array {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subgizmo_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
