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



// Enums

func (me *AStar2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AStar2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AStar2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AStar2D) XEstimateCost(from_id int, to_id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) XComputeCost(from_id int, to_id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetAvailablePointId()  {
  panic("TODO: implement")
}

func  (me *AStar2D) AddPoint(id int, position Vector2, weight_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetPointPosition(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) SetPointPosition(id int, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetPointWeightScale(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) SetPointWeightScale(id int, weight_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) RemovePoint(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) HasPoint(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetPointConnections(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetPointIds()  {
  panic("TODO: implement")
}

func  (me *AStar2D) SetPointDisabled(id int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) IsPointDisabled(id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) ConnectPoints(id int, to_id int, bidirectional bool, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) DisconnectPoints(id int, to_id int, bidirectional bool, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) ArePointsConnected(id int, to_id int, bidirectional bool, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetPointCount()  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetPointCapacity()  {
  panic("TODO: implement")
}

func  (me *AStar2D) ReserveSpace(num_nodes int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) Clear()  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetClosestPoint(to_position Vector2, include_disabled bool, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetClosestPositionInSegment(to_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetPointPath(from_id int, to_id int, )  {
  panic("TODO: implement")
}

func  (me *AStar2D) GetIdPath(from_id int, to_id int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
