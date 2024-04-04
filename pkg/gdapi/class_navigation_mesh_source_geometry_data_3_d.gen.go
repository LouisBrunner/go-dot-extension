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

type NavigationMeshSourceGeometryData3D struct {
  Resource
}

func (me *NavigationMeshSourceGeometryData3D) BaseClass() string {
  return "NavigationMeshSourceGeometryData3D"
}

func NewNavigationMeshSourceGeometryData3D() *NavigationMeshSourceGeometryData3D {
  str := StringNameFromStr("NavigationMeshSourceGeometryData3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationMeshSourceGeometryData3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NavigationMeshSourceGeometryData3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationMeshSourceGeometryData3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationMeshSourceGeometryData3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationMeshSourceGeometryData3D) SetVertices(vertices PackedFloat32Array, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshSourceGeometryData3D) GetVertices() PackedFloat32Array {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationMeshSourceGeometryData3D) SetIndices(indices PackedInt32Array, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_indices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{indices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshSourceGeometryData3D) GetIndices() PackedInt32Array {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_indices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationMeshSourceGeometryData3D) Clear()  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshSourceGeometryData3D) HasData() bool {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMeshSourceGeometryData3D) AddMesh(mesh Mesh, xform Transform3D, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 975462459) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshSourceGeometryData3D) AddMeshArray(mesh_array Array, xform Transform3D, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_mesh_array")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4235710913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh_array.AsCTypePtr(), xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshSourceGeometryData3D) AddFaces(faces PackedVector3Array, xform Transform3D, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1440358797) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{faces.AsCTypePtr(), xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
