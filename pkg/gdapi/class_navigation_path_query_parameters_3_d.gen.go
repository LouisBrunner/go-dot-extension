// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPathQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPathQueryParameters3D) BaseClass() string {
  return "NavigationPathQueryParameters3D"
}

type NavigationPathQueryParameters3DPathfindingAlgorithm int
const (
  NavigationPathQueryParameters3DPathfindingAlgorithmAstar NavigationPathQueryParameters3DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters3DPathPostProcessing int
const (
  NavigationPathQueryParameters3DPathPostprocessingCorridorfunnel NavigationPathQueryParameters3DPathPostProcessing = 0
  NavigationPathQueryParameters3DPathPostprocessingEdgecentered NavigationPathQueryParameters3DPathPostProcessing = 1
)

type NavigationPathQueryParameters3DPathMetadataFlags int
const (
  NavigationPathQueryParameters3DPathMetadataIncludeNone NavigationPathQueryParameters3DPathMetadataFlags = 0
  NavigationPathQueryParameters3DPathMetadataIncludeTypes NavigationPathQueryParameters3DPathMetadataFlags = 1
  NavigationPathQueryParameters3DPathMetadataIncludeRids NavigationPathQueryParameters3DPathMetadataFlags = 2
  NavigationPathQueryParameters3DPathMetadataIncludeOwners NavigationPathQueryParameters3DPathMetadataFlags = 4
  NavigationPathQueryParameters3DPathMetadataIncludeAll NavigationPathQueryParameters3DPathMetadataFlags = 7
)
