// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Performance struct {
  obj gdc.ObjectPtr
}

func (me *Performance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Performance) BaseClass() string {
  return "Performance"
}

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

func  (me *Performance) GetMonitor(monitor PerformanceMonitor, ) { // TODO: return value
  // TODO: implement
}

func  (me *Performance) AddCustomMonitor(id StringName, callable Callable, arguments Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Performance) RemoveCustomMonitor(id StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Performance) HasCustomMonitor(id StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Performance) GetCustomMonitor(id StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Performance) GetMonitorModificationTime() { // TODO: return value
  // TODO: implement
}

func  (me *Performance) GetCustomMonitorNames() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
