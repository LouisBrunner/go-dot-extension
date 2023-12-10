// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationPolygon struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPolygon) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPolygon) BaseClass() string {
  return "NavigationPolygon"
}



// Enums

func (me *NavigationPolygon) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationPolygon) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPolygon) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationPolygon) SetVertices(vertices PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vertices.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) GetVertices() PackedVector2Array {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPolygon) AddPolygon(polygon PackedInt32Array, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) GetPolygonCount() int {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPolygon) GetPolygon(idx int, ) PackedInt32Array {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3668444399) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPolygon) ClearPolygons()  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) GetNavigationMesh() NavigationMesh {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 330232164) // FIXME: should cache?
  var ret NavigationMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPolygon) AddOutline(outline PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(outline.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) AddOutlineAtIndex(outline PackedVector2Array, index int, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_outline_at_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1569738947) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(outline.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) GetOutlineCount() int {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPolygon) SetOutline(idx int, outline PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1201971903) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(outline.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) GetOutline(idx int, ) PackedVector2Array {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3946907486) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationPolygon) RemoveOutline(idx int, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) ClearOutlines()  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) MakePolygonsFromOutlines()  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_polygons_from_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) SetCellSize(cell_size float32, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cell_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationPolygon) GetCellSize() float32 {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *NavigationPolygon) GetPropVertices() PackedVector2Array {
  panic("TODO: implement")
}

func (me *NavigationPolygon) SetPropVertices(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *NavigationPolygon) GetPropPolygons() Array {
  panic("TODO: implement")
}

func (me *NavigationPolygon) SetPropPolygons(value Array) {
  panic("TODO: implement")
}

func (me *NavigationPolygon) GetPropOutlines() Array {
  panic("TODO: implement")
}

func (me *NavigationPolygon) SetPropOutlines(value Array) {
  panic("TODO: implement")
}

func (me *NavigationPolygon) GetPropCellSize() float32 {
  panic("TODO: implement")
}

func (me *NavigationPolygon) SetPropCellSize(value float32) {
  panic("TODO: implement")
}