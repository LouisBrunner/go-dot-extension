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

type GraphEdit struct {
  Control
}

func (me *GraphEdit) BaseClass() string {
  return "GraphEdit"
}

func NewGraphEdit() *GraphEdit {
  str := StringNameFromStr("GraphEdit") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GraphEdit{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type GraphEditPanningScheme int
const (
  GraphEditPanningSchemeScrollZooms GraphEditPanningScheme = 0
  GraphEditPanningSchemeScrollPans GraphEditPanningScheme = 1
)

func (me *GraphEdit) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GraphEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GraphEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GraphEdit) ConnectNode(from_node StringName, from_port int64, to_node StringName, to_port int64, ) Error {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 195065850) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port) , to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&from_port)
  pinner.Pin(&to_port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GraphEdit) IsNodeConnected(from_node StringName, from_port int64, to_node StringName, to_port int64, ) bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_node_connected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4216241294) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port) , to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&from_port)
  pinner.Pin(&to_port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) DisconnectNode(from_node StringName, from_port int64, to_node StringName, to_port int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1933654315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port) , to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) SetConnectionActivity(from_node StringName, from_port int64, to_node StringName, to_port int64, amount float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_activity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1141899943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port) , to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port) , gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetConnectionList() []Dictionary {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GraphEdit) ClearConnections()  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) ForceConnectionDragEnd()  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_connection_drag_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetScrollOffset() Vector2 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphEdit) SetScrollOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) AddValidRightDisconnectType(type_ int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_valid_right_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) RemoveValidRightDisconnectType(type_ int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_valid_right_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) AddValidLeftDisconnectType(type_ int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_valid_left_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) RemoveValidLeftDisconnectType(type_ int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_valid_left_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) AddValidConnectionType(from_type int64, to_type int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_valid_connection_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type) , gdc.ConstTypePtr(&to_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) RemoveValidConnectionType(from_type int64, to_type int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_valid_connection_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type) , gdc.ConstTypePtr(&to_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsValidConnectionType(from_type int64, to_type int64, ) bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid_connection_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type) , gdc.ConstTypePtr(&to_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&from_type)
  pinner.Pin(&to_type)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) GetConnectionLine(from_node Vector2, to_node Vector2, ) PackedVector2Array {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1562168077) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), to_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphEdit) SetPanningScheme(scheme GraphEditPanningScheme, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_panning_scheme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 18893313) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scheme) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetPanningScheme() GraphEditPanningScheme {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_panning_scheme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 549924446) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GraphEditPanningScheme

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GraphEdit) SetZoom(zoom float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetZoom() float64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetZoomMin(zoom_min float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_min) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetZoomMin() float64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetZoomMax(zoom_max float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_max) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetZoomMax() float64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetZoomStep(zoom_step float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_step) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetZoomStep() float64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetShowGrid(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_grid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsShowingGrid() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_grid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetSnappingEnabled(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snapping_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsSnappingEnabled() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_snapping_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetSnappingDistance(pixels int64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snapping_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetSnappingDistance() int64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_snapping_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetConnectionLinesCurvature(curvature float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_lines_curvature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curvature) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetConnectionLinesCurvature() float64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_lines_curvature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetConnectionLinesThickness(pixels float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_lines_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetConnectionLinesThickness() float64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_lines_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetConnectionLinesAntialiased(pixels bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_lines_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsConnectionLinesAntialiased() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_connection_lines_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetMinimapSize(size Vector2, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minimap_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetMinimapSize() Vector2 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimap_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphEdit) SetMinimapOpacity(opacity float64, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minimap_opacity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&opacity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) GetMinimapOpacity() float64 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimap_opacity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetMinimapEnabled(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minimap_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsMinimapEnabled() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_minimap_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetShowMenu(hidden bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsShowingMenu() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetShowZoomLabel(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_zoom_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsShowingZoomLabel() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_zoom_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetShowGridButtons(hidden bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_grid_buttons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsShowingGridButtons() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_grid_buttons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetShowZoomButtons(hidden bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_zoom_buttons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsShowingZoomButtons() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_zoom_buttons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetShowMinimapButton(hidden bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_minimap_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsShowingMinimapButton() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_minimap_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetShowArrangeButton(hidden bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_arrange_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsShowingArrangeButton() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_arrange_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) SetRightDisconnects(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_right_disconnects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) IsRightDisconnectsEnabled() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_right_disconnects_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphEdit) GetMenuHbox() HBoxContainer {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu_hbox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3590609951) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewHBoxContainer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphEdit) ArrangeNodes()  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("arrange_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphEdit) SetSelected(node Node, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphEditConnectionRequestSignalFn func(from_node StringName, from_port int, to_node StringName, to_port int, )

func (me *GraphEdit) ConnectConnectionRequest(subs SignalSubscribers, fn GraphEditConnectionRequestSignalFn) {
  sig := StringNameFromStr("connection_request")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionRequest(subs SignalSubscribers, fn GraphEditConnectionRequestSignalFn) {
  sig := StringNameFromStr("connection_request")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditDisconnectionRequestSignalFn func(from_node StringName, from_port int, to_node StringName, to_port int, )

func (me *GraphEdit) ConnectDisconnectionRequest(subs SignalSubscribers, fn GraphEditDisconnectionRequestSignalFn) {
  sig := StringNameFromStr("disconnection_request")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectDisconnectionRequest(subs SignalSubscribers, fn GraphEditDisconnectionRequestSignalFn) {
  sig := StringNameFromStr("disconnection_request")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionToEmptySignalFn func(from_node StringName, from_port int, release_position Vector2, )

func (me *GraphEdit) ConnectConnectionToEmpty(subs SignalSubscribers, fn GraphEditConnectionToEmptySignalFn) {
  sig := StringNameFromStr("connection_to_empty")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionToEmpty(subs SignalSubscribers, fn GraphEditConnectionToEmptySignalFn) {
  sig := StringNameFromStr("connection_to_empty")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionFromEmptySignalFn func(to_node StringName, to_port int, release_position Vector2, )

func (me *GraphEdit) ConnectConnectionFromEmpty(subs SignalSubscribers, fn GraphEditConnectionFromEmptySignalFn) {
  sig := StringNameFromStr("connection_from_empty")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionFromEmpty(subs SignalSubscribers, fn GraphEditConnectionFromEmptySignalFn) {
  sig := StringNameFromStr("connection_from_empty")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionDragStartedSignalFn func(from_node StringName, from_port int, is_output bool, )

func (me *GraphEdit) ConnectConnectionDragStarted(subs SignalSubscribers, fn GraphEditConnectionDragStartedSignalFn) {
  sig := StringNameFromStr("connection_drag_started")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionDragStarted(subs SignalSubscribers, fn GraphEditConnectionDragStartedSignalFn) {
  sig := StringNameFromStr("connection_drag_started")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionDragEndedSignalFn func()

func (me *GraphEdit) ConnectConnectionDragEnded(subs SignalSubscribers, fn GraphEditConnectionDragEndedSignalFn) {
  sig := StringNameFromStr("connection_drag_ended")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionDragEnded(subs SignalSubscribers, fn GraphEditConnectionDragEndedSignalFn) {
  sig := StringNameFromStr("connection_drag_ended")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditCopyNodesRequestSignalFn func()

func (me *GraphEdit) ConnectCopyNodesRequest(subs SignalSubscribers, fn GraphEditCopyNodesRequestSignalFn) {
  sig := StringNameFromStr("copy_nodes_request")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectCopyNodesRequest(subs SignalSubscribers, fn GraphEditCopyNodesRequestSignalFn) {
  sig := StringNameFromStr("copy_nodes_request")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditPasteNodesRequestSignalFn func()

func (me *GraphEdit) ConnectPasteNodesRequest(subs SignalSubscribers, fn GraphEditPasteNodesRequestSignalFn) {
  sig := StringNameFromStr("paste_nodes_request")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectPasteNodesRequest(subs SignalSubscribers, fn GraphEditPasteNodesRequestSignalFn) {
  sig := StringNameFromStr("paste_nodes_request")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditDuplicateNodesRequestSignalFn func()

func (me *GraphEdit) ConnectDuplicateNodesRequest(subs SignalSubscribers, fn GraphEditDuplicateNodesRequestSignalFn) {
  sig := StringNameFromStr("duplicate_nodes_request")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectDuplicateNodesRequest(subs SignalSubscribers, fn GraphEditDuplicateNodesRequestSignalFn) {
  sig := StringNameFromStr("duplicate_nodes_request")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditDeleteNodesRequestSignalFn func(nodes []StringName, )

func (me *GraphEdit) ConnectDeleteNodesRequest(subs SignalSubscribers, fn GraphEditDeleteNodesRequestSignalFn) {
  sig := StringNameFromStr("delete_nodes_request")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectDeleteNodesRequest(subs SignalSubscribers, fn GraphEditDeleteNodesRequestSignalFn) {
  sig := StringNameFromStr("delete_nodes_request")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditNodeSelectedSignalFn func(node Node, )

func (me *GraphEdit) ConnectNodeSelected(subs SignalSubscribers, fn GraphEditNodeSelectedSignalFn) {
  sig := StringNameFromStr("node_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectNodeSelected(subs SignalSubscribers, fn GraphEditNodeSelectedSignalFn) {
  sig := StringNameFromStr("node_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditNodeDeselectedSignalFn func(node Node, )

func (me *GraphEdit) ConnectNodeDeselected(subs SignalSubscribers, fn GraphEditNodeDeselectedSignalFn) {
  sig := StringNameFromStr("node_deselected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectNodeDeselected(subs SignalSubscribers, fn GraphEditNodeDeselectedSignalFn) {
  sig := StringNameFromStr("node_deselected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditPopupRequestSignalFn func(position Vector2, )

func (me *GraphEdit) ConnectPopupRequest(subs SignalSubscribers, fn GraphEditPopupRequestSignalFn) {
  sig := StringNameFromStr("popup_request")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectPopupRequest(subs SignalSubscribers, fn GraphEditPopupRequestSignalFn) {
  sig := StringNameFromStr("popup_request")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditBeginNodeMoveSignalFn func()

func (me *GraphEdit) ConnectBeginNodeMove(subs SignalSubscribers, fn GraphEditBeginNodeMoveSignalFn) {
  sig := StringNameFromStr("begin_node_move")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectBeginNodeMove(subs SignalSubscribers, fn GraphEditBeginNodeMoveSignalFn) {
  sig := StringNameFromStr("begin_node_move")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditEndNodeMoveSignalFn func()

func (me *GraphEdit) ConnectEndNodeMove(subs SignalSubscribers, fn GraphEditEndNodeMoveSignalFn) {
  sig := StringNameFromStr("end_node_move")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectEndNodeMove(subs SignalSubscribers, fn GraphEditEndNodeMoveSignalFn) {
  sig := StringNameFromStr("end_node_move")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditScrollOffsetChangedSignalFn func(offset Vector2, )

func (me *GraphEdit) ConnectScrollOffsetChanged(subs SignalSubscribers, fn GraphEditScrollOffsetChangedSignalFn) {
  sig := StringNameFromStr("scroll_offset_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectScrollOffsetChanged(subs SignalSubscribers, fn GraphEditScrollOffsetChangedSignalFn) {
  sig := StringNameFromStr("scroll_offset_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
