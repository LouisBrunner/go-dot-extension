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

type ptrsForAStar2DList struct {
  fnXEstimateCost gdc.MethodBindPtr
  fnXComputeCost gdc.MethodBindPtr
  fnGetAvailablePointId gdc.MethodBindPtr
  fnAddPoint gdc.MethodBindPtr
  fnGetPointPosition gdc.MethodBindPtr
  fnSetPointPosition gdc.MethodBindPtr
  fnGetPointWeightScale gdc.MethodBindPtr
  fnSetPointWeightScale gdc.MethodBindPtr
  fnRemovePoint gdc.MethodBindPtr
  fnHasPoint gdc.MethodBindPtr
  fnGetPointConnections gdc.MethodBindPtr
  fnGetPointIds gdc.MethodBindPtr
  fnSetPointDisabled gdc.MethodBindPtr
  fnIsPointDisabled gdc.MethodBindPtr
  fnConnectPoints gdc.MethodBindPtr
  fnDisconnectPoints gdc.MethodBindPtr
  fnArePointsConnected gdc.MethodBindPtr
  fnGetPointCount gdc.MethodBindPtr
  fnGetPointCapacity gdc.MethodBindPtr
  fnReserveSpace gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnGetClosestPoint gdc.MethodBindPtr
  fnGetClosestPositionInSegment gdc.MethodBindPtr
  fnGetPointPath gdc.MethodBindPtr
  fnGetIdPath gdc.MethodBindPtr
}

var ptrsForAStar2D ptrsForAStar2DList

func initAStar2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AStar2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_available_point_id")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetAvailablePointId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("add_point")
    defer methodName.Destroy()
    ptrsForAStar2D.fnAddPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4074201818))
  }
  {
    methodName := StringNameFromStr("get_point_position")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
  }
  {
    methodName := StringNameFromStr("set_point_position")
    defer methodName.Destroy()
    ptrsForAStar2D.fnSetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
  }
  {
    methodName := StringNameFromStr("get_point_weight_scale")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetPointWeightScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("set_point_weight_scale")
    defer methodName.Destroy()
    ptrsForAStar2D.fnSetPointWeightScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("remove_point")
    defer methodName.Destroy()
    ptrsForAStar2D.fnRemovePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("has_point")
    defer methodName.Destroy()
    ptrsForAStar2D.fnHasPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("get_point_connections")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetPointConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2865087369))
  }
  {
    methodName := StringNameFromStr("get_point_ids")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetPointIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3851388692))
  }
  {
    methodName := StringNameFromStr("set_point_disabled")
    defer methodName.Destroy()
    ptrsForAStar2D.fnSetPointDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 972357352))
  }
  {
    methodName := StringNameFromStr("is_point_disabled")
    defer methodName.Destroy()
    ptrsForAStar2D.fnIsPointDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("connect_points")
    defer methodName.Destroy()
    ptrsForAStar2D.fnConnectPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3710494224))
  }
  {
    methodName := StringNameFromStr("disconnect_points")
    defer methodName.Destroy()
    ptrsForAStar2D.fnDisconnectPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3710494224))
  }
  {
    methodName := StringNameFromStr("are_points_connected")
    defer methodName.Destroy()
    ptrsForAStar2D.fnArePointsConnected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2288175859))
  }
  {
    methodName := StringNameFromStr("get_point_count")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_point_capacity")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetPointCapacity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("reserve_space")
    defer methodName.Destroy()
    ptrsForAStar2D.fnReserveSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForAStar2D.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_closest_point")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetClosestPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2300324924))
  }
  {
    methodName := StringNameFromStr("get_closest_position_in_segment")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetClosestPositionInSegment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2656412154))
  }
  {
    methodName := StringNameFromStr("get_point_path")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetPointPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 281625055))
  }
  {
    methodName := StringNameFromStr("get_id_path")
    defer methodName.Destroy()
    ptrsForAStar2D.fnGetIdPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3404614526))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetAvailablePointId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) AddPoint(id int64, position Vector2, weight_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , position.AsCTypePtr(), gdc.ConstTypePtr(&weight_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnAddPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) GetPointPosition(id int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) SetPointPosition(id int64, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnSetPointPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) GetPointWeightScale(id int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetPointWeightScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) SetPointWeightScale(id int64, weight_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&weight_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnSetPointWeightScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) RemovePoint(id int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnRemovePoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) HasPoint(id int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnHasPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetPointConnections(id int64, ) PackedInt64Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetPointConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) GetPointIds() PackedInt64Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetPointIds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) SetPointDisabled(id int64, disabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnSetPointDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) IsPointDisabled(id int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnIsPointDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) ConnectPoints(id int64, to_id int64, bidirectional bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&to_id) , gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnConnectPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) DisconnectPoints(id int64, to_id int64, bidirectional bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&to_id) , gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnDisconnectPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) ArePointsConnected(id int64, to_id int64, bidirectional bool, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&to_id) , gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)
  pinner.Pin(&to_id)
  pinner.Pin(&bidirectional)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnArePointsConnected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetPointCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetPointCapacity() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetPointCapacity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) ReserveSpace(num_nodes int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&num_nodes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnReserveSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AStar2D) GetClosestPoint(to_position Vector2, include_disabled bool, ) int64 {
  cargs := []gdc.ConstTypePtr{to_position.AsCTypePtr(), gdc.ConstTypePtr(&include_disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&include_disabled)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetClosestPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AStar2D) GetClosestPositionInSegment(to_position Vector2, ) Vector2 {
  cargs := []gdc.ConstTypePtr{to_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetClosestPositionInSegment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) GetPointPath(from_id int64, to_id int64, ) PackedVector2Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id) , gdc.ConstTypePtr(&to_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()
  pinner.Pin(&from_id)
  pinner.Pin(&to_id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetPointPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AStar2D) GetIdPath(from_id int64, to_id int64, ) PackedInt64Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id) , gdc.ConstTypePtr(&to_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()
  pinner.Pin(&from_id)
  pinner.Pin(&to_id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar2D.fnGetIdPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
