// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationMeshSourceGeometryData3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationMeshSourceGeometryData3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationMeshSourceGeometryData3D) BaseClass() string {
  return "NavigationMeshSourceGeometryData3D"
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
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) GetVertices()  {
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) SetIndices(indices PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) GetIndices()  {
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) Clear()  {
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) HasData()  {
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) AddMesh(mesh Mesh, xform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) AddMeshArray(mesh_array Array, xform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *NavigationMeshSourceGeometryData3D) AddFaces(faces PackedVector3Array, xform Transform3D, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
