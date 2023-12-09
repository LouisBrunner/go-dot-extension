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



// Enums

func (me *AStar3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AStar3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AStar3D) XEstimateCost(from_id int, to_id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) XComputeCost(from_id int, to_id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetAvailablePointId()  {
  panic("TODO: implement")
}

func  (me *AStar3D) AddPoint(id int, position Vector3, weight_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetPointPosition(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) SetPointPosition(id int, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetPointWeightScale(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) SetPointWeightScale(id int, weight_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) RemovePoint(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) HasPoint(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetPointConnections(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetPointIds()  {
  panic("TODO: implement")
}

func  (me *AStar3D) SetPointDisabled(id int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) IsPointDisabled(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) ConnectPoints(id int, to_id int, bidirectional bool, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) DisconnectPoints(id int, to_id int, bidirectional bool, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) ArePointsConnected(id int, to_id int, bidirectional bool, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetPointCount()  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetPointCapacity()  {
  panic("TODO: implement")
}

func  (me *AStar3D) ReserveSpace(num_nodes int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) Clear()  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetClosestPoint(to_position Vector3, include_disabled bool, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetClosestPositionInSegment(to_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetPointPath(from_id int, to_id int, )  {
  panic("TODO: implement")
}

func  (me *AStar3D) GetIdPath(from_id int, to_id int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
