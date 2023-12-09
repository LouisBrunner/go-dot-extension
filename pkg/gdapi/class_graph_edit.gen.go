// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *GraphEdit) XIsInInputHotzone(in_node Object, in_port int, mouse_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) XIsInOutputHotzone(in_node Object, in_port int, mouse_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) XGetConnectionLine(from_position Vector2, to_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) XIsNodeHoverValid(from_node StringName, from_port int, to_node StringName, to_port int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) ConnectNode(from_node StringName, from_port int, to_node StringName, to_port int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsNodeConnected(from_node StringName, from_port int, to_node StringName, to_port int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) DisconnectNode(from_node StringName, from_port int, to_node StringName, to_port int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetConnectionActivity(from_node StringName, from_port int, to_node StringName, to_port int, amount float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetConnectionList()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) ClearConnections()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) ForceConnectionDragEnd()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetScrollOfs()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetScrollOfs(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) AddValidRightDisconnectType(type_ int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) RemoveValidRightDisconnectType(type_ int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) AddValidLeftDisconnectType(type_ int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) RemoveValidLeftDisconnectType(type_ int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) AddValidConnectionType(from_type int, to_type int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) RemoveValidConnectionType(from_type int, to_type int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsValidConnectionType(from_type int, to_type int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetConnectionLine(from_node Vector2, to_node Vector2, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetPanningScheme(scheme GraphEditPanningScheme, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetPanningScheme()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetZoom(zoom float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetZoom()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetZoomMin(zoom_min float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetZoomMin()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetZoomMax(zoom_max float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetZoomMax()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetZoomStep(zoom_step float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetZoomStep()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetShowZoomLabel(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsShowingZoomLabel()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetSnap(pixels int, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetSnap()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetUseSnap(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsUsingSnap()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetConnectionLinesCurvature(curvature float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetConnectionLinesCurvature()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetConnectionLinesThickness(pixels float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetConnectionLinesThickness()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetConnectionLinesAntialiased(pixels bool, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsConnectionLinesAntialiased()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetMinimapSize(size Vector2, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetMinimapSize()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetMinimapOpacity(opacity float32, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetMinimapOpacity()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetMinimapEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsMinimapEnabled()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetArrangeNodesButtonHidden(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsArrangeNodesButtonHidden()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetRightDisconnects(enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphEdit) IsRightDisconnectsEnabled()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) GetZoomHbox()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) ArrangeNodes()  {
  panic("TODO: implement")
}

func  (me *GraphEdit) SetSelected(node Node, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
