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

type NavigationPolygon struct {
  Resource
}

func (me *NavigationPolygon) BaseClass() string {
  return "NavigationPolygon"
}

func NewNavigationPolygon() *NavigationPolygon {
  str := StringNameFromStr("NavigationPolygon") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationPolygon{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type NavigationPolygonParsedGeometryType int
const (
  NavigationPolygonParsedGeometryTypeParsedGeometryMeshInstances NavigationPolygonParsedGeometryType = 0
  NavigationPolygonParsedGeometryTypeParsedGeometryStaticColliders NavigationPolygonParsedGeometryType = 1
  NavigationPolygonParsedGeometryTypeParsedGeometryBoth NavigationPolygonParsedGeometryType = 2
  NavigationPolygonParsedGeometryTypeParsedGeometryMax NavigationPolygonParsedGeometryType = 3
)

type NavigationPolygonSourceGeometryMode int
const (
  NavigationPolygonSourceGeometryModeSourceGeometryRootNodeChildren NavigationPolygonSourceGeometryMode = 0
  NavigationPolygonSourceGeometryModeSourceGeometryGroupsWithChildren NavigationPolygonSourceGeometryMode = 1
  NavigationPolygonSourceGeometryModeSourceGeometryGroupsExplicit NavigationPolygonSourceGeometryMode = 2
  NavigationPolygonSourceGeometryModeSourceGeometryMax NavigationPolygonSourceGeometryMode = 3
)

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
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetVertices() PackedVector2Array {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) AddPolygon(polygon PackedInt32Array, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetPolygonCount() int64 {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) GetPolygon(idx int64, ) PackedInt32Array {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3668444399) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) ClearPolygons()  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetNavigationMesh() NavigationMesh {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 330232164) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNavigationMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) AddOutline(outline PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{outline.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) AddOutlineAtIndex(outline PackedVector2Array, index int64, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_outline_at_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1569738947) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{outline.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetOutlineCount() int64 {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetOutline(idx int64, outline PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1201971903) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , outline.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetOutline(idx int64, ) PackedVector2Array {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3946907486) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) RemoveOutline(idx int64, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) ClearOutlines()  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) MakePolygonsFromOutlines()  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_polygons_from_outlines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) SetCellSize(cell_size float64, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cell_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetCellSize() float64 {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetParsedGeometryType(geometry_type NavigationPolygonParsedGeometryType, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parsed_geometry_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2507971764) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&geometry_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetParsedGeometryType() NavigationPolygonParsedGeometryType {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parsed_geometry_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1073219508) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPolygonParsedGeometryType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationPolygon) SetParsedCollisionMask(mask int64, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parsed_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetParsedCollisionMask() int64 {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parsed_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetParsedCollisionMaskValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parsed_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetParsedCollisionMaskValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parsed_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetSourceGeometryMode(geometry_mode NavigationPolygonSourceGeometryMode, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source_geometry_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4002316705) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&geometry_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetSourceGeometryMode() NavigationPolygonSourceGeometryMode {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_geometry_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 459686762) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPolygonSourceGeometryMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationPolygon) SetSourceGeometryGroupName(group_name StringName, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source_geometry_group_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{group_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetSourceGeometryGroupName() StringName {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_geometry_group_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) SetAgentRadius(agent_radius float64, )  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_agent_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetAgentRadius() float64 {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_agent_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) Clear()  {
  classNameV := StringNameFromStr("NavigationPolygon")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
