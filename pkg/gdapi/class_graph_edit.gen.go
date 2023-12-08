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

type GraphEditPanningScheme int
const (
  GraphEditPanningSchemeScrollZooms GraphEditPanningScheme = 0
  GraphEditPanningSchemeScrollPans GraphEditPanningScheme = 1
)

func  (me *GraphEdit) XIsInInputHotzone(in_node Object, in_port int, mouse_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) XIsInOutputHotzone(in_node Object, in_port int, mouse_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) XGetConnectionLine(from_position Vector2, to_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) XIsNodeHoverValid(from_node StringName, from_port int, to_node StringName, to_port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) ConnectNode(from_node StringName, from_port int, to_node StringName, to_port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsNodeConnected(from_node StringName, from_port int, to_node StringName, to_port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) DisconnectNode(from_node StringName, from_port int, to_node StringName, to_port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetConnectionActivity(from_node StringName, from_port int, to_node StringName, to_port int, amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetConnectionList() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) ClearConnections() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) ForceConnectionDragEnd() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetScrollOfs() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetScrollOfs(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) AddValidRightDisconnectType(type_ int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) RemoveValidRightDisconnectType(type_ int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) AddValidLeftDisconnectType(type_ int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) RemoveValidLeftDisconnectType(type_ int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) AddValidConnectionType(from_type int, to_type int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) RemoveValidConnectionType(from_type int, to_type int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsValidConnectionType(from_type int, to_type int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetConnectionLine(from_node Vector2, to_node Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetPanningScheme(scheme GraphEditPanningScheme, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetPanningScheme() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetZoom(zoom float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetZoom() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetZoomMin(zoom_min float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetZoomMin() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetZoomMax(zoom_max float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetZoomMax() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetZoomStep(zoom_step float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetZoomStep() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetShowZoomLabel(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsShowingZoomLabel() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetSnap(pixels int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetSnap() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetUseSnap(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsUsingSnap() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetConnectionLinesCurvature(curvature float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetConnectionLinesCurvature() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetConnectionLinesThickness(pixels float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetConnectionLinesThickness() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetConnectionLinesAntialiased(pixels bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsConnectionLinesAntialiased() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetMinimapSize(size Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetMinimapSize() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetMinimapOpacity(opacity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetMinimapOpacity() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetMinimapEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsMinimapEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetArrangeNodesButtonHidden(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsArrangeNodesButtonHidden() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetRightDisconnects(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) IsRightDisconnectsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) GetZoomHbox() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) ArrangeNodes() { // TODO: return value
  // TODO: implement
}

func  (me *GraphEdit) SetSelected(node Node, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
