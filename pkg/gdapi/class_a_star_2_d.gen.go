// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *AStar2D) GetAvailablePointId() int {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_available_point_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) AddPoint(id int, position Vector2, weight_scale float32, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4074201818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&weight_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) GetPointPosition(id int, ) Vector2 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) SetPointPosition(id int, position Vector2, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) GetPointWeightScale(id int, ) float32 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_weight_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) SetPointWeightScale(id int, weight_scale float32, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_weight_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&weight_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) RemovePoint(id int, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) HasPoint(id int, ) bool {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) GetPointConnections(id int, ) PackedInt64Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2865087369) // FIXME: should cache?
  var ret PackedInt64Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) GetPointIds() PackedInt64Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3851388692) // FIXME: should cache?
  var ret PackedInt64Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) SetPointDisabled(id int, disabled bool, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 972357352) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) IsPointDisabled(id int, ) bool {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_point_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) ConnectPoints(id int, to_id int, bidirectional bool, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3710494224) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&bidirectional), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) DisconnectPoints(id int, to_id int, bidirectional bool, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3710494224) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&bidirectional), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) ArePointsConnected(id int, to_id int, bidirectional bool, ) bool {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_points_connected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2288175859) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&bidirectional), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) GetPointCount() int {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) GetPointCapacity() int {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_capacity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) ReserveSpace(num_nodes int, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reserve_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&num_nodes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) Clear()  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AStar2D) GetClosestPoint(to_position Vector2, include_disabled bool, ) int {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2300324924) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to_position.AsCTypePtr()), gdc.ConstTypePtr(&include_disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) GetClosestPositionInSegment(to_position Vector2, ) Vector2 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_position_in_segment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to_position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) GetPointPath(from_id int, to_id int, ) PackedVector2Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 281625055) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id), gdc.ConstTypePtr(&to_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AStar2D) GetIdPath(from_id int, to_id int, ) PackedInt64Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3404614526) // FIXME: should cache?
  var ret PackedInt64Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id), gdc.ConstTypePtr(&to_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
