// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStar2D struct {
  obj gdc.ObjectPtr
}

func (me *AStar2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AStar2D) BaseClass() string {
  return "AStar2D"
}

func  (me *AStar2D) XEstimateCost(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) XComputeCost(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetAvailablePointId() { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) AddPoint(id int, position Vector2, weight_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetPointPosition(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) SetPointPosition(id int, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetPointWeightScale(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) SetPointWeightScale(id int, weight_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) RemovePoint(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) HasPoint(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetPointConnections(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetPointIds() { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) SetPointDisabled(id int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) IsPointDisabled(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) ConnectPoints(id int, to_id int, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) DisconnectPoints(id int, to_id int, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) ArePointsConnected(id int, to_id int, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetPointCapacity() { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) ReserveSpace(num_nodes int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetClosestPoint(to_position Vector2, include_disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetClosestPositionInSegment(to_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetPointPath(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar2D) GetIdPath(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
