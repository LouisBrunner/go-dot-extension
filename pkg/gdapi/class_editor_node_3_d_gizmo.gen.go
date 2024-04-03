// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorNode3DGizmo struct {
  Node3DGizmo
}

func (me *EditorNode3DGizmo) BaseClass() string {
  return "EditorNode3DGizmo"
}

func NewEditorNode3DGizmo() *EditorNode3DGizmo {
  str := StringNameFromStr("EditorNode3DGizmo") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorNode3DGizmo{}
  obj.SetBaseObject(objPtr)
  return obj
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2910971437) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(lines.AsCTypePtr()), gdc.ConstTypePtr(material.AsCTypePtr()), gdc.ConstTypePtr(&billboard), gdc.ConstTypePtr(modulate.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) AddMesh(mesh Mesh, material Material, transform Transform3D, skeleton SkinReference, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1579955111) // FIXME: should cache?
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

func  (me *EditorNode3DGizmo) AddUnscaledBillboard(material Material, default_scale float64, modulate Color, )  {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_unscaled_billboard")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 520007164) // FIXME: should cache?
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewNode3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorNode3DGizmo) GetPlugin() EditorNode3DGizmoPlugin {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4250544552) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewEditorNode3DGizmoPlugin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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

func  (me *EditorNode3DGizmo) IsSubgizmoSelected(id int64, ) bool {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_subgizmo_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorNode3DGizmo) GetSubgizmoSelection() PackedInt32Array {
  classNameV := StringNameFromStr("EditorNode3DGizmo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_subgizmo_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
