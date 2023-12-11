// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Performance struct {
  obj gdc.ObjectPtr
}

func (me *Performance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Performance) BaseClass() string {
  return "Performance"
}



// Enums

type PerformanceMonitor int
const (
  PerformanceMonitorTimeFps PerformanceMonitor = 0
  PerformanceMonitorTimeProcess PerformanceMonitor = 1
  PerformanceMonitorTimePhysicsProcess PerformanceMonitor = 2
  PerformanceMonitorTimeNavigationProcess PerformanceMonitor = 3
  PerformanceMonitorMemoryStatic PerformanceMonitor = 4
  PerformanceMonitorMemoryStaticMax PerformanceMonitor = 5
  PerformanceMonitorMemoryMessageBufferMax PerformanceMonitor = 6
  PerformanceMonitorObjectCount PerformanceMonitor = 7
  PerformanceMonitorObjectResourceCount PerformanceMonitor = 8
  PerformanceMonitorObjectNodeCount PerformanceMonitor = 9
  PerformanceMonitorObjectOrphanNodeCount PerformanceMonitor = 10
  PerformanceMonitorRenderTotalObjectsInFrame PerformanceMonitor = 11
  PerformanceMonitorRenderTotalPrimitivesInFrame PerformanceMonitor = 12
  PerformanceMonitorRenderTotalDrawCallsInFrame PerformanceMonitor = 13
  PerformanceMonitorRenderVideoMemUsed PerformanceMonitor = 14
  PerformanceMonitorRenderTextureMemUsed PerformanceMonitor = 15
  PerformanceMonitorRenderBufferMemUsed PerformanceMonitor = 16
  PerformanceMonitorPhysics2DActiveObjects PerformanceMonitor = 17
  PerformanceMonitorPhysics2DCollisionPairs PerformanceMonitor = 18
  PerformanceMonitorPhysics2DIslandCount PerformanceMonitor = 19
  PerformanceMonitorPhysics3DActiveObjects PerformanceMonitor = 20
  PerformanceMonitorPhysics3DCollisionPairs PerformanceMonitor = 21
  PerformanceMonitorPhysics3DIslandCount PerformanceMonitor = 22
  PerformanceMonitorAudioOutputLatency PerformanceMonitor = 23
  PerformanceMonitorNavigationActiveMaps PerformanceMonitor = 24
  PerformanceMonitorNavigationRegionCount PerformanceMonitor = 25
  PerformanceMonitorNavigationAgentCount PerformanceMonitor = 26
  PerformanceMonitorNavigationLinkCount PerformanceMonitor = 27
  PerformanceMonitorNavigationPolygonCount PerformanceMonitor = 28
  PerformanceMonitorNavigationEdgeCount PerformanceMonitor = 29
  PerformanceMonitorNavigationEdgeMergeCount PerformanceMonitor = 30
  PerformanceMonitorNavigationEdgeConnectionCount PerformanceMonitor = 31
  PerformanceMonitorNavigationEdgeFreeCount PerformanceMonitor = 32
  PerformanceMonitorMonitorMax PerformanceMonitor = 33
)

func (me *Performance) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Performance) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Performance) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Performance) GetMonitor(monitor PerformanceMonitor, ) float32 {
  classNameV := StringNameFromStr("Performance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_monitor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1943275655) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&monitor), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Performance) AddCustomMonitor(id StringName, callable Callable, arguments Array, )  {
  classNameV := StringNameFromStr("Performance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_custom_monitor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2865980031) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), gdc.ConstTypePtr(callable.AsCTypePtr()), gdc.ConstTypePtr(arguments.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Performance) RemoveCustomMonitor(id StringName, )  {
  classNameV := StringNameFromStr("Performance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_custom_monitor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Performance) HasCustomMonitor(id StringName, ) bool {
  classNameV := StringNameFromStr("Performance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_custom_monitor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2041966384) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Performance) GetCustomMonitor(id StringName, ) Variant {
  classNameV := StringNameFromStr("Performance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_monitor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2138907829) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Performance) GetMonitorModificationTime() int {
  classNameV := StringNameFromStr("Performance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_monitor_modification_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Performance) GetCustomMonitorNames() StringName {
  classNameV := StringNameFromStr("Performance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_monitor_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
