// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationMesh struct {
  obj gdc.ObjectPtr
}

func (me *NavigationMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationMesh) BaseClass() string {
  return "NavigationMesh"
}

type NavigationMeshSamplePartitionType int
const (
  NavigationMeshSamplePartitionWatershed NavigationMeshSamplePartitionType = 0
  NavigationMeshSamplePartitionMonotone NavigationMeshSamplePartitionType = 1
  NavigationMeshSamplePartitionLayers NavigationMeshSamplePartitionType = 2
  NavigationMeshSamplePartitionMax NavigationMeshSamplePartitionType = 3
)

type NavigationMeshParsedGeometryType int
const (
  NavigationMeshParsedGeometryMeshInstances NavigationMeshParsedGeometryType = 0
  NavigationMeshParsedGeometryStaticColliders NavigationMeshParsedGeometryType = 1
  NavigationMeshParsedGeometryBoth NavigationMeshParsedGeometryType = 2
  NavigationMeshParsedGeometryMax NavigationMeshParsedGeometryType = 3
)

type NavigationMeshSourceGeometryMode int
const (
  NavigationMeshSourceGeometryRootNodeChildren NavigationMeshSourceGeometryMode = 0
  NavigationMeshSourceGeometryGroupsWithChildren NavigationMeshSourceGeometryMode = 1
  NavigationMeshSourceGeometryGroupsExplicit NavigationMeshSourceGeometryMode = 2
  NavigationMeshSourceGeometryMax NavigationMeshSourceGeometryMode = 3
)
