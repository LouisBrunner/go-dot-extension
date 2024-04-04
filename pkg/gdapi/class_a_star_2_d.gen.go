// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type AStar2D struct {
  RefCounted
}

func (me *AStar2D) BaseClass() string {
  return "AStar2D"
}

func NewAStar2D() *AStar2D {
  str := StringNameFromStr("AStar2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AStar2D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AStar2D) GetAvailablePointId() int64 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_available_point_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) AddPoint(id int64, position Vector2, weight_scale float64, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4074201818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , position.AsCTypePtr(), gdc.ConstTypePtr(&weight_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) GetPointPosition(id int64, ) Vector2 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) SetPointPosition(id int64, position Vector2, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) GetPointWeightScale(id int64, ) float64 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_weight_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) SetPointWeightScale(id int64, weight_scale float64, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_weight_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&weight_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) RemovePoint(id int64, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) HasPoint(id int64, ) bool {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetPointConnections(id int64, ) PackedInt64Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2865087369) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) GetPointIds() PackedInt64Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3851388692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) SetPointDisabled(id int64, disabled bool, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 972357352) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) IsPointDisabled(id int64, ) bool {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_point_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) ConnectPoints(id int64, to_id int64, bidirectional bool, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3710494224) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&to_id) , gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) DisconnectPoints(id int64, to_id int64, bidirectional bool, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3710494224) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&to_id) , gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) ArePointsConnected(id int64, to_id int64, bidirectional bool, ) bool {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_points_connected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2288175859) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&to_id) , gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)
  pinner.Pin(&to_id)
  pinner.Pin(&bidirectional)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetPointCount() int64 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetPointCapacity() int64 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_capacity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) ReserveSpace(num_nodes int64, )  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reserve_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&num_nodes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) Clear()  {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) GetClosestPoint(to_position Vector2, include_disabled bool, ) int64 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2300324924) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{to_position.AsCTypePtr(), gdc.ConstTypePtr(&include_disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&include_disabled)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetClosestPositionInSegment(to_position Vector2, ) Vector2 {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_position_in_segment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{to_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) GetPointPath(from_id int64, to_id int64, ) PackedVector2Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 281625055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id) , gdc.ConstTypePtr(&to_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()
  pinner.Pin(&from_id)
  pinner.Pin(&to_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) GetIdPath(from_id int64, to_id int64, ) PackedInt64Array {
  classNameV := StringNameFromStr("AStar2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3404614526) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id) , gdc.ConstTypePtr(&to_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()
  pinner.Pin(&from_id)
  pinner.Pin(&to_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
