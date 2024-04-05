// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForEditorNode3DGizmoList struct {
  fnXRedraw gdc.MethodBindPtr
  fnXGetHandleName gdc.MethodBindPtr
  fnXIsHandleHighlighted gdc.MethodBindPtr
  fnXGetHandleValue gdc.MethodBindPtr
  fnXSetHandle gdc.MethodBindPtr
  fnXCommitHandle gdc.MethodBindPtr
  fnXSubgizmosIntersectRay gdc.MethodBindPtr
  fnXSubgizmosIntersectFrustum gdc.MethodBindPtr
  fnXSetSubgizmoTransform gdc.MethodBindPtr
  fnXGetSubgizmoTransform gdc.MethodBindPtr
  fnXCommitSubgizmos gdc.MethodBindPtr
  fnAddLines gdc.MethodBindPtr
  fnAddMesh gdc.MethodBindPtr
  fnAddCollisionSegments gdc.MethodBindPtr
  fnAddCollisionTriangles gdc.MethodBindPtr
  fnAddUnscaledBillboard gdc.MethodBindPtr
  fnAddHandles gdc.MethodBindPtr
  fnSetNode3D gdc.MethodBindPtr
  fnGetNode3D gdc.MethodBindPtr
  fnGetPlugin gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnSetHidden gdc.MethodBindPtr
  fnIsSubgizmoSelected gdc.MethodBindPtr
  fnGetSubgizmoSelection gdc.MethodBindPtr
}

var ptrsForEditorNode3DGizmo ptrsForEditorNode3DGizmoList

func initEditorNode3DGizmoPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorNode3DGizmo")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_lines")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnAddLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2910971437))
  }
  {
    methodName := StringNameFromStr("add_mesh")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnAddMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1579955111))
  }
  {
    methodName := StringNameFromStr("add_collision_segments")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnAddCollisionSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
  }
  {
    methodName := StringNameFromStr("add_collision_triangles")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnAddCollisionTriangles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 54901064))
  }
  {
    methodName := StringNameFromStr("add_unscaled_billboard")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnAddUnscaledBillboard = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 520007164))
  }
  {
    methodName := StringNameFromStr("add_handles")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnAddHandles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2254560097))
  }
  {
    methodName := StringNameFromStr("set_node_3d")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnSetNode3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
  }
  {
    methodName := StringNameFromStr("get_node_3d")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnGetNode3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 151077316))
  }
  {
    methodName := StringNameFromStr("get_plugin")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnGetPlugin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4250544552))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_hidden")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnSetHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_subgizmo_selected")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnIsSubgizmoSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("get_subgizmo_selection")
    defer methodName.Destroy()
    ptrsForEditorNode3DGizmo.fnGetSubgizmoSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
}

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
  cargs := []gdc.ConstTypePtr{lines.AsCTypePtr(), material.AsCTypePtr(), gdc.ConstTypePtr(&billboard) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnAddLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) AddMesh(mesh Mesh, material Material, transform Transform3D, skeleton SkinReference, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), material.AsCTypePtr(), transform.AsCTypePtr(), skeleton.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnAddMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) AddCollisionSegments(segments PackedVector3Array, )  {
  cargs := []gdc.ConstTypePtr{segments.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnAddCollisionSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) AddCollisionTriangles(triangles TriangleMesh, )  {
  cargs := []gdc.ConstTypePtr{triangles.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnAddCollisionTriangles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) AddUnscaledBillboard(material Material, default_scale float64, modulate Color, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), gdc.ConstTypePtr(&default_scale) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnAddUnscaledBillboard), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) AddHandles(handles PackedVector3Array, material Material, ids PackedInt32Array, billboard bool, secondary bool, )  {
  cargs := []gdc.ConstTypePtr{handles.AsCTypePtr(), material.AsCTypePtr(), ids.AsCTypePtr(), gdc.ConstTypePtr(&billboard) , gdc.ConstTypePtr(&secondary) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnAddHandles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) SetNode3D(node Node, )  {
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnSetNode3D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) GetNode3D() Node3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnGetNode3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorNode3DGizmo) GetPlugin() EditorNode3DGizmoPlugin {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorNode3DGizmoPlugin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnGetPlugin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorNode3DGizmo) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) SetHidden(hidden bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnSetHidden), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorNode3DGizmo) IsSubgizmoSelected(id int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnIsSubgizmoSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorNode3DGizmo) GetSubgizmoSelection() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorNode3DGizmo.fnGetSubgizmoSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
