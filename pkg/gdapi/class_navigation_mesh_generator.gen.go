// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationMeshGenerator struct {
  obj gdc.ObjectPtr
}

func (me *NavigationMeshGenerator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationMeshGenerator) BaseClass() string {
  return "NavigationMeshGenerator"
}



// Enums

func (me *NavigationMeshGenerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationMeshGenerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationMeshGenerator) Bake(navigation_mesh NavigationMesh, root_node Node, )  {
  panic("TODO: implement")
}

func  (me *NavigationMeshGenerator) Clear(navigation_mesh NavigationMesh, )  {
  panic("TODO: implement")
}

func  (me *NavigationMeshGenerator) ParseSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, root_node Node, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *NavigationMeshGenerator) BakeFromSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
