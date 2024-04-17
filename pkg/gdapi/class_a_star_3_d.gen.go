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

type ptrsForAStar3DList struct {
	fnXEstimateCost               gdc.MethodBindPtr
	fnXComputeCost                gdc.MethodBindPtr
	fnGetAvailablePointId         gdc.MethodBindPtr
	fnAddPoint                    gdc.MethodBindPtr
	fnGetPointPosition            gdc.MethodBindPtr
	fnSetPointPosition            gdc.MethodBindPtr
	fnGetPointWeightScale         gdc.MethodBindPtr
	fnSetPointWeightScale         gdc.MethodBindPtr
	fnRemovePoint                 gdc.MethodBindPtr
	fnHasPoint                    gdc.MethodBindPtr
	fnGetPointConnections         gdc.MethodBindPtr
	fnGetPointIds                 gdc.MethodBindPtr
	fnSetPointDisabled            gdc.MethodBindPtr
	fnIsPointDisabled             gdc.MethodBindPtr
	fnConnectPoints               gdc.MethodBindPtr
	fnDisconnectPoints            gdc.MethodBindPtr
	fnArePointsConnected          gdc.MethodBindPtr
	fnGetPointCount               gdc.MethodBindPtr
	fnGetPointCapacity            gdc.MethodBindPtr
	fnReserveSpace                gdc.MethodBindPtr
	fnClear                       gdc.MethodBindPtr
	fnGetClosestPoint             gdc.MethodBindPtr
	fnGetClosestPositionInSegment gdc.MethodBindPtr
	fnGetPointPath                gdc.MethodBindPtr
	fnGetIdPath                   gdc.MethodBindPtr
}

var ptrsForAStar3D ptrsForAStar3DList

func initAStar3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AStar3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_available_point_id")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetAvailablePointId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_point")
		defer methodName.Destroy()
		ptrsForAStar3D.fnAddPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1038703438))
	}
	{
		methodName := StringNameFromStr("get_point_position")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("set_point_position")
		defer methodName.Destroy()
		ptrsForAStar3D.fnSetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
	}
	{
		methodName := StringNameFromStr("get_point_weight_scale")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetPointWeightScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_point_weight_scale")
		defer methodName.Destroy()
		ptrsForAStar3D.fnSetPointWeightScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("remove_point")
		defer methodName.Destroy()
		ptrsForAStar3D.fnRemovePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("has_point")
		defer methodName.Destroy()
		ptrsForAStar3D.fnHasPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_point_connections")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetPointConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2865087369))
	}
	{
		methodName := StringNameFromStr("get_point_ids")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetPointIds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3851388692))
	}
	{
		methodName := StringNameFromStr("set_point_disabled")
		defer methodName.Destroy()
		ptrsForAStar3D.fnSetPointDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 972357352))
	}
	{
		methodName := StringNameFromStr("is_point_disabled")
		defer methodName.Destroy()
		ptrsForAStar3D.fnIsPointDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("connect_points")
		defer methodName.Destroy()
		ptrsForAStar3D.fnConnectPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3710494224))
	}
	{
		methodName := StringNameFromStr("disconnect_points")
		defer methodName.Destroy()
		ptrsForAStar3D.fnDisconnectPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3710494224))
	}
	{
		methodName := StringNameFromStr("are_points_connected")
		defer methodName.Destroy()
		ptrsForAStar3D.fnArePointsConnected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2288175859))
	}
	{
		methodName := StringNameFromStr("get_point_count")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_point_capacity")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetPointCapacity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("reserve_space")
		defer methodName.Destroy()
		ptrsForAStar3D.fnReserveSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForAStar3D.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_closest_point")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetClosestPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3241074317))
	}
	{
		methodName := StringNameFromStr("get_closest_position_in_segment")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetClosestPositionInSegment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 192990374))
	}
	{
		methodName := StringNameFromStr("get_point_path")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetPointPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1562654675))
	}
	{
		methodName := StringNameFromStr("get_id_path")
		defer methodName.Destroy()
		ptrsForAStar3D.fnGetIdPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3136199648))
	}

}

type AStar3D struct {
	RefCounted
}

func (me *AStar3D) BaseClass() string {
	return "AStar3D"
}

func NewAStar3D() *AStar3D {
	str := StringNameFromStr("AStar3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AStar3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AStar3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AStar3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AStar3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AStar3D) GetAvailablePointId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetAvailablePointId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) AddPoint(id int64, position Vector3, weight_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), position.AsCTypePtr(), gdc.ConstTypePtr(&weight_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnAddPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) GetPointPosition(id int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStar3D) SetPointPosition(id int64, position Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnSetPointPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) GetPointWeightScale(id int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetPointWeightScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) SetPointWeightScale(id int64, weight_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&weight_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnSetPointWeightScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) RemovePoint(id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnRemovePoint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) HasPoint(id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnHasPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) GetPointConnections(id int64) PackedInt64Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt64Array()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetPointConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStar3D) GetPointIds() PackedInt64Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt64Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetPointIds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStar3D) SetPointDisabled(id int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnSetPointDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) IsPointDisabled(id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnIsPointDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) ConnectPoints(id int64, to_id int64, bidirectional bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&bidirectional)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnConnectPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) DisconnectPoints(id int64, to_id int64, bidirectional bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&bidirectional)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnDisconnectPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) ArePointsConnected(id int64, to_id int64, bidirectional bool) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&bidirectional)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&id)
	pinner.Pin(&to_id)
	pinner.Pin(&bidirectional)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnArePointsConnected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) GetPointCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) GetPointCapacity() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetPointCapacity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) ReserveSpace(num_nodes int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&num_nodes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnReserveSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AStar3D) GetClosestPoint(to_position Vector3, include_disabled bool) int64 {
	cargs := []gdc.ConstTypePtr{to_position.AsCTypePtr(), gdc.ConstTypePtr(&include_disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&include_disabled)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetClosestPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AStar3D) GetClosestPositionInSegment(to_position Vector3) Vector3 {
	cargs := []gdc.ConstTypePtr{to_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetClosestPositionInSegment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStar3D) GetPointPath(from_id int64, to_id int64, allow_partial_path bool) PackedVector3Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&allow_partial_path)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()
	pinner.Pin(&from_id)
	pinner.Pin(&to_id)
	pinner.Pin(&allow_partial_path)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetPointPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AStar3D) GetIdPath(from_id int64, to_id int64, allow_partial_path bool) PackedInt64Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_id), gdc.ConstTypePtr(&to_id), gdc.ConstTypePtr(&allow_partial_path)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt64Array()
	pinner.Pin(&from_id)
	pinner.Pin(&to_id)
	pinner.Pin(&allow_partial_path)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAStar3D.fnGetIdPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
