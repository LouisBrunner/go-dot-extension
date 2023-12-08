// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  PerformanceTimeFps PerformanceMonitor = 0
  PerformanceTimeProcess PerformanceMonitor = 1
  PerformanceTimePhysicsProcess PerformanceMonitor = 2
  PerformanceTimeNavigationProcess PerformanceMonitor = 3
  PerformanceMemoryStatic PerformanceMonitor = 4
  PerformanceMemoryStaticMax PerformanceMonitor = 5
  PerformanceMemoryMessageBufferMax PerformanceMonitor = 6
  PerformanceObjectCount PerformanceMonitor = 7
  PerformanceObjectResourceCount PerformanceMonitor = 8
  PerformanceObjectNodeCount PerformanceMonitor = 9
  PerformanceObjectOrphanNodeCount PerformanceMonitor = 10
  PerformanceRenderTotalObjectsInFrame PerformanceMonitor = 11
  PerformanceRenderTotalPrimitivesInFrame PerformanceMonitor = 12
  PerformanceRenderTotalDrawCallsInFrame PerformanceMonitor = 13
  PerformanceRenderVideoMemUsed PerformanceMonitor = 14
  PerformanceRenderTextureMemUsed PerformanceMonitor = 15
  PerformanceRenderBufferMemUsed PerformanceMonitor = 16
  PerformancePhysics2DActiveObjects PerformanceMonitor = 17
  PerformancePhysics2DCollisionPairs PerformanceMonitor = 18
  PerformancePhysics2DIslandCount PerformanceMonitor = 19
  PerformancePhysics3DActiveObjects PerformanceMonitor = 20
  PerformancePhysics3DCollisionPairs PerformanceMonitor = 21
  PerformancePhysics3DIslandCount PerformanceMonitor = 22
  PerformanceAudioOutputLatency PerformanceMonitor = 23
  PerformanceNavigationActiveMaps PerformanceMonitor = 24
  PerformanceNavigationRegionCount PerformanceMonitor = 25
  PerformanceNavigationAgentCount PerformanceMonitor = 26
  PerformanceNavigationLinkCount PerformanceMonitor = 27
  PerformanceNavigationPolygonCount PerformanceMonitor = 28
  PerformanceNavigationEdgeCount PerformanceMonitor = 29
  PerformanceNavigationEdgeMergeCount PerformanceMonitor = 30
  PerformanceNavigationEdgeConnectionCount PerformanceMonitor = 31
  PerformanceNavigationEdgeFreeCount PerformanceMonitor = 32
  PerformanceMonitorMax PerformanceMonitor = 33
)
