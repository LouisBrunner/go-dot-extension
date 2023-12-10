// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GraphEdit struct {
  obj gdc.ObjectPtr
}

func (me *GraphEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GraphEdit) BaseClass() string {
  return "GraphEdit"
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

func  (me *GraphEdit) ConnectNode(from_node StringName, from_port int, to_node StringName, to_port int, ) Error {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 195065850) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(to_node.AsCTypePtr()), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) IsNodeConnected(from_node StringName, from_port int, to_node StringName, to_port int, ) bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_node_connected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4216241294) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(to_node.AsCTypePtr()), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) DisconnectNode(from_node StringName, from_port int, to_node StringName, to_port int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1933654315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(to_node.AsCTypePtr()), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) SetConnectionActivity(from_node StringName, from_port int, to_node StringName, to_port int, amount float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_activity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1141899943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(to_node.AsCTypePtr()), gdc.ConstTypePtr(&to_port), gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetConnectionList() Dictionary {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) ClearConnections()  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) ForceConnectionDragEnd()  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_connection_drag_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetScrollOfs() Vector2 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_ofs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetScrollOfs(offset Vector2, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_ofs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) AddValidRightDisconnectType(type_ int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_valid_right_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) RemoveValidRightDisconnectType(type_ int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_valid_right_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) AddValidLeftDisconnectType(type_ int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_valid_left_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) RemoveValidLeftDisconnectType(type_ int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_valid_left_disconnect_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) AddValidConnectionType(from_type int, to_type int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_valid_connection_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type), gdc.ConstTypePtr(&to_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) RemoveValidConnectionType(from_type int, to_type int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_valid_connection_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type), gdc.ConstTypePtr(&to_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) IsValidConnectionType(from_type int, to_type int, ) bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid_connection_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type), gdc.ConstTypePtr(&to_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) GetConnectionLine(from_node Vector2, to_node Vector2, ) PackedVector2Array {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1562168077) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from_node.AsCTypePtr()), gdc.ConstTypePtr(to_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetPanningScheme(scheme GraphEditPanningScheme, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_panning_scheme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 18893313) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scheme), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetPanningScheme() GraphEditPanningScheme {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_panning_scheme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 549924446) // FIXME: should cache?
  var ret GraphEditPanningScheme
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetZoom(zoom float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetZoom() float32 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetZoomMin(zoom_min float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_min), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetZoomMin() float32 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetZoomMax(zoom_max float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_max), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetZoomMax() float32 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetZoomStep(zoom_step float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zoom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_step), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetZoomStep() float32 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetShowZoomLabel(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_zoom_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) IsShowingZoomLabel() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_showing_zoom_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetSnap(pixels int, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetSnap() int {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetUseSnap(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) IsUsingSnap() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetConnectionLinesCurvature(curvature float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_lines_curvature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curvature), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetConnectionLinesCurvature() float32 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_lines_curvature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetConnectionLinesThickness(pixels float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_lines_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetConnectionLinesThickness() float32 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_lines_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetConnectionLinesAntialiased(pixels bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection_lines_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) IsConnectionLinesAntialiased() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_connection_lines_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetMinimapSize(size Vector2, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minimap_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetMinimapSize() Vector2 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimap_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetMinimapOpacity(opacity float32, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minimap_opacity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&opacity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) GetMinimapOpacity() float32 {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimap_opacity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetMinimapEnabled(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_minimap_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) IsMinimapEnabled() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_minimap_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetArrangeNodesButtonHidden(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_arrange_nodes_button_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) IsArrangeNodesButtonHidden() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_arrange_nodes_button_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) SetRightDisconnects(enable bool, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_right_disconnects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) IsRightDisconnectsEnabled() bool {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_right_disconnects_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) GetZoomHbox() HBoxContainer {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zoom_hbox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3590609951) // FIXME: should cache?
  var ret HBoxContainer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphEdit) ArrangeNodes()  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("arrange_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphEdit) SetSelected(node Node, )  {
  classNameV := StringNameFromStr("GraphEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *GraphEdit) GetPropRightDisconnects() bool {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropRightDisconnects(value bool) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropScrollOffset() Vector2 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropScrollOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropSnapDistance() int {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropSnapDistance(value int) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropUseSnap() bool {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropUseSnap(value bool) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropPanningScheme() int {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropPanningScheme(value int) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropConnectionLinesCurvature() float32 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropConnectionLinesCurvature(value float32) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropConnectionLinesThickness() float32 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropConnectionLinesThickness(value float32) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropConnectionLinesAntialiased() bool {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropConnectionLinesAntialiased(value bool) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropZoom() float32 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropZoom(value float32) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropZoomMin() float32 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropZoomMin(value float32) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropZoomMax() float32 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropZoomMax(value float32) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropZoomStep() float32 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropZoomStep(value float32) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropShowZoomLabel() bool {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropShowZoomLabel(value bool) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropMinimapEnabled() bool {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropMinimapEnabled(value bool) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropMinimapSize() Vector2 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropMinimapSize(value Vector2) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropMinimapOpacity() float32 {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropMinimapOpacity(value float32) {
  panic("TODO: implement")
}

func (me *GraphEdit) GetPropArrangeNodesButtonHidden() bool {
  panic("TODO: implement")
}

func (me *GraphEdit) SetPropArrangeNodesButtonHidden(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API