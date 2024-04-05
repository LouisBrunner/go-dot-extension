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

type ptrsForNavigationPolygonList struct {
  fnSetVertices gdc.MethodBindPtr
  fnGetVertices gdc.MethodBindPtr
  fnAddPolygon gdc.MethodBindPtr
  fnGetPolygonCount gdc.MethodBindPtr
  fnGetPolygon gdc.MethodBindPtr
  fnClearPolygons gdc.MethodBindPtr
  fnGetNavigationMesh gdc.MethodBindPtr
  fnAddOutline gdc.MethodBindPtr
  fnAddOutlineAtIndex gdc.MethodBindPtr
  fnGetOutlineCount gdc.MethodBindPtr
  fnSetOutline gdc.MethodBindPtr
  fnGetOutline gdc.MethodBindPtr
  fnRemoveOutline gdc.MethodBindPtr
  fnClearOutlines gdc.MethodBindPtr
  fnMakePolygonsFromOutlines gdc.MethodBindPtr
  fnSetCellSize gdc.MethodBindPtr
  fnGetCellSize gdc.MethodBindPtr
  fnSetParsedGeometryType gdc.MethodBindPtr
  fnGetParsedGeometryType gdc.MethodBindPtr
  fnSetParsedCollisionMask gdc.MethodBindPtr
  fnGetParsedCollisionMask gdc.MethodBindPtr
  fnSetParsedCollisionMaskValue gdc.MethodBindPtr
  fnGetParsedCollisionMaskValue gdc.MethodBindPtr
  fnSetSourceGeometryMode gdc.MethodBindPtr
  fnGetSourceGeometryMode gdc.MethodBindPtr
  fnSetSourceGeometryGroupName gdc.MethodBindPtr
  fnGetSourceGeometryGroupName gdc.MethodBindPtr
  fnSetAgentRadius gdc.MethodBindPtr
  fnGetAgentRadius gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
}

var ptrsForNavigationPolygon ptrsForNavigationPolygonList

func initNavigationPolygonPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NavigationPolygon")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_vertices")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
  }
  {
    methodName := StringNameFromStr("get_vertices")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
  }
  {
    methodName := StringNameFromStr("add_polygon")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnAddPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_polygon_count")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetPolygonCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_polygon")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3668444399))
  }
  {
    methodName := StringNameFromStr("clear_polygons")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnClearPolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_navigation_mesh")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 330232164))
  }
  {
    methodName := StringNameFromStr("add_outline")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnAddOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
  }
  {
    methodName := StringNameFromStr("add_outline_at_index")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnAddOutlineAtIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1569738947))
  }
  {
    methodName := StringNameFromStr("get_outline_count")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetOutlineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_outline")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1201971903))
  }
  {
    methodName := StringNameFromStr("get_outline")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3946907486))
  }
  {
    methodName := StringNameFromStr("remove_outline")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnRemoveOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("clear_outlines")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnClearOutlines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("make_polygons_from_outlines")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnMakePolygonsFromOutlines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_cell_size")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_cell_size")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_parsed_geometry_type")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetParsedGeometryType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2507971764))
  }
  {
    methodName := StringNameFromStr("get_parsed_geometry_type")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetParsedGeometryType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1073219508))
  }
  {
    methodName := StringNameFromStr("set_parsed_collision_mask")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetParsedCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_parsed_collision_mask")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetParsedCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_parsed_collision_mask_value")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetParsedCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_parsed_collision_mask_value")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetParsedCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_source_geometry_mode")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetSourceGeometryMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4002316705))
  }
  {
    methodName := StringNameFromStr("get_source_geometry_mode")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetSourceGeometryMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 459686762))
  }
  {
    methodName := StringNameFromStr("set_source_geometry_group_name")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetSourceGeometryGroupName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_source_geometry_group_name")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetSourceGeometryGroupName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_agent_radius")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnSetAgentRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_agent_radius")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnGetAgentRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForNavigationPolygon.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

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
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetVertices() PackedVector2Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetVertices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) AddPolygon(polygon PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnAddPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetPolygonCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetPolygonCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) GetPolygon(idx int64, ) PackedInt32Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) ClearPolygons()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnClearPolygons), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetNavigationMesh() NavigationMesh {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNavigationMesh()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetNavigationMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) AddOutline(outline PackedVector2Array, )  {
  cargs := []gdc.ConstTypePtr{outline.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnAddOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) AddOutlineAtIndex(outline PackedVector2Array, index int64, )  {
  cargs := []gdc.ConstTypePtr{outline.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnAddOutlineAtIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetOutlineCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetOutlineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetOutline(idx int64, outline PackedVector2Array, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , outline.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetOutline(idx int64, ) PackedVector2Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetOutline), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) RemoveOutline(idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnRemoveOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) ClearOutlines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnClearOutlines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) MakePolygonsFromOutlines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnMakePolygonsFromOutlines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) SetCellSize(cell_size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cell_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetCellSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetParsedGeometryType(geometry_type NavigationPolygonParsedGeometryType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&geometry_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetParsedGeometryType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetParsedGeometryType() NavigationPolygonParsedGeometryType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPolygonParsedGeometryType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetParsedGeometryType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationPolygon) SetParsedCollisionMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetParsedCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetParsedCollisionMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetParsedCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetParsedCollisionMaskValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetParsedCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetParsedCollisionMaskValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetParsedCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) SetSourceGeometryMode(geometry_mode NavigationPolygonSourceGeometryMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&geometry_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetSourceGeometryMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetSourceGeometryMode() NavigationPolygonSourceGeometryMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPolygonSourceGeometryMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetSourceGeometryMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationPolygon) SetSourceGeometryGroupName(group_name StringName, )  {
  cargs := []gdc.ConstTypePtr{group_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetSourceGeometryGroupName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetSourceGeometryGroupName() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetSourceGeometryGroupName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationPolygon) SetAgentRadius(agent_radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnSetAgentRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationPolygon) GetAgentRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnGetAgentRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationPolygon) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationPolygon.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
