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

type NavigationMeshGenerator struct {
  Object
}

func (me *NavigationMeshGenerator) BaseClass() string {
  return "NavigationMeshGenerator"
}

func NewNavigationMeshGenerator() *NavigationMeshGenerator {
  str := StringNameFromStr("NavigationMeshGenerator") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationMeshGenerator{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), root_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshGenerator) Clear(navigation_mesh NavigationMesh, )  {
  classNameV := StringNameFromStr("NavigationMeshGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2923361153) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshGenerator) ParseSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, root_node Node, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationMeshGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 685862123) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), root_node.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMeshGenerator) BakeFromSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationMeshGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_from_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2469318639) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
