// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *NavigationMeshGenerator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationMeshGenerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationMeshGenerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationMeshGenerator) Bake(navigation_mesh NavigationMesh, root_node Node, )  {
  classNameV := StringNameFromStr("NavigationMeshGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1401173477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_mesh.AsCTypePtr()), gdc.ConstTypePtr(root_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMeshGenerator) Clear(navigation_mesh NavigationMesh, )  {
  classNameV := StringNameFromStr("NavigationMeshGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2923361153) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMeshGenerator) ParseSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, root_node Node, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationMeshGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 685862123) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_mesh.AsCTypePtr()), gdc.ConstTypePtr(source_geometry_data.AsCTypePtr()), gdc.ConstTypePtr(root_node.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMeshGenerator) BakeFromSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationMeshGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_from_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2469318639) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_mesh.AsCTypePtr()), gdc.ConstTypePtr(source_geometry_data.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
