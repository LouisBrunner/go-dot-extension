// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationRegion2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationRegion2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationRegion2D) BaseClass() string {
  return "NavigationRegion2D"
}



// Enums

func (me *NavigationRegion2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationRegion2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationRegion2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationRegion2D) SetNavigationPolygon(navigation_polygon NavigationPolygon, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1515040758) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetNavigationPolygon() NavigationPolygon {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1046532237) // FIXME: should cache?
  var ret NavigationPolygon
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) IsEnabled() bool {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetNavigationMap(navigation_map RID, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_map.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetNavigationMap() RID {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetUseEdgeConnections(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetUseEdgeConnections() bool {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetNavigationLayers(navigation_layers int, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetNavigationLayers() int {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetNavigationLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetNavigationLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetConstrainAvoidance(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constrain_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetConstrainAvoidance() bool {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constrain_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetAvoidanceLayers(layers int, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetAvoidanceLayers() int {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetAvoidanceLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetAvoidanceLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) GetRegionRid() RID {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetEnterCost(enter_cost float32, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetEnterCost() float32 {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) SetTravelCost(travel_cost float32, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&travel_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationRegion2D) GetTravelCost() float32 {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationRegion2D) BakeNavigationPolygon(on_thread bool, )  {
  classNameV := StringNameFromStr("NavigationRegion2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_navigation_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_thread), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NavigationRegion2DNavigationPolygonChangedSignalFn func()

func (me *NavigationRegion2D) ConnectNavigationPolygonChanged(subs SignalSubscribers, fn NavigationRegion2DNavigationPolygonChangedSignalFn) {
  sig := StringNameFromStr("navigation_polygon_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *NavigationRegion2D) DisconnectNavigationPolygonChanged(subs SignalSubscribers, fn NavigationRegion2DNavigationPolygonChangedSignalFn) {
  sig := StringNameFromStr("navigation_polygon_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type NavigationRegion2DBakeFinishedSignalFn func()

func (me *NavigationRegion2D) ConnectBakeFinished(subs SignalSubscribers, fn NavigationRegion2DBakeFinishedSignalFn) {
  sig := StringNameFromStr("bake_finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *NavigationRegion2D) DisconnectBakeFinished(subs SignalSubscribers, fn NavigationRegion2DBakeFinishedSignalFn) {
  sig := StringNameFromStr("bake_finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
