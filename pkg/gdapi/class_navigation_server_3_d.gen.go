// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationServer3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationServer3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationServer3D) BaseClass() string {
  return "NavigationServer3D"
}

type NavigationServer3DProcessInfo int
const (
  NavigationServer3DInfoActiveMaps NavigationServer3DProcessInfo = 0
  NavigationServer3DInfoRegionCount NavigationServer3DProcessInfo = 1
  NavigationServer3DInfoAgentCount NavigationServer3DProcessInfo = 2
  NavigationServer3DInfoLinkCount NavigationServer3DProcessInfo = 3
  NavigationServer3DInfoPolygonCount NavigationServer3DProcessInfo = 4
  NavigationServer3DInfoEdgeCount NavigationServer3DProcessInfo = 5
  NavigationServer3DInfoEdgeMergeCount NavigationServer3DProcessInfo = 6
  NavigationServer3DInfoEdgeConnectionCount NavigationServer3DProcessInfo = 7
  NavigationServer3DInfoEdgeFreeCount NavigationServer3DProcessInfo = 8
)
