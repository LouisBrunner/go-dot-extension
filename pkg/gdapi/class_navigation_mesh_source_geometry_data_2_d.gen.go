// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationMeshSourceGeometryData2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationMeshSourceGeometryData2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationMeshSourceGeometryData2D) BaseClass() string {
  return "NavigationMeshSourceGeometryData2D"
}



// Enums

func (me *NavigationMeshSourceGeometryData2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationMeshSourceGeometryData2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationMeshSourceGeometryData2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationMeshSourceGeometryData2D) Clear()  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMeshSourceGeometryData2D) HasData() bool {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMeshSourceGeometryData2D) SetTraversableOutlines(traversable_outlines PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_traversable_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(traversable_outlines.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMeshSourceGeometryData2D) GetTraversableOutlines() PackedVector2Array {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_traversable_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMeshSourceGeometryData2D) SetObstructionOutlines(obstruction_outlines PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_obstruction_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstruction_outlines.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMeshSourceGeometryData2D) GetObstructionOutlines() PackedVector2Array {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_obstruction_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMeshSourceGeometryData2D) AddTraversableOutline(shape_outline PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_traversable_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape_outline.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMeshSourceGeometryData2D) AddObstructionOutline(shape_outline PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationMeshSourceGeometryData2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_obstruction_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape_outline.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
