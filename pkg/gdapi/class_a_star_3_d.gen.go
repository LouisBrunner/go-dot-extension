// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AStar3D struct {
  obj gdc.ObjectPtr
}

func (me *AStar3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AStar3D) BaseClass() string {
  return "AStar3D"
}

func  (me *AStar3D) XEstimateCost(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) XComputeCost(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetAvailablePointId() { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) AddPoint(id int, position Vector3, weight_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetPointPosition(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) SetPointPosition(id int, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetPointWeightScale(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) SetPointWeightScale(id int, weight_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) RemovePoint(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) HasPoint(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetPointConnections(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetPointIds() { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) SetPointDisabled(id int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) IsPointDisabled(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) ConnectPoints(id int, to_id int, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) DisconnectPoints(id int, to_id int, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) ArePointsConnected(id int, to_id int, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetPointCapacity() { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) ReserveSpace(num_nodes int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetClosestPoint(to_position Vector3, include_disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetClosestPositionInSegment(to_position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetPointPath(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AStar3D) GetIdPath(from_id int, to_id int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
