// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPathQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPathQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPathQueryParameters2D) BaseClass() string {
  return "NavigationPathQueryParameters2D"
}

type NavigationPathQueryParameters2DPathfindingAlgorithm int
const (
  NavigationPathQueryParameters2DPathfindingAlgorithmAstar NavigationPathQueryParameters2DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters2DPathPostProcessing int
const (
  NavigationPathQueryParameters2DPathPostprocessingCorridorfunnel NavigationPathQueryParameters2DPathPostProcessing = 0
  NavigationPathQueryParameters2DPathPostprocessingEdgecentered NavigationPathQueryParameters2DPathPostProcessing = 1
)

type NavigationPathQueryParameters2DPathMetadataFlags int
const (
  NavigationPathQueryParameters2DPathMetadataIncludeNone NavigationPathQueryParameters2DPathMetadataFlags = 0
  NavigationPathQueryParameters2DPathMetadataIncludeTypes NavigationPathQueryParameters2DPathMetadataFlags = 1
  NavigationPathQueryParameters2DPathMetadataIncludeRids NavigationPathQueryParameters2DPathMetadataFlags = 2
  NavigationPathQueryParameters2DPathMetadataIncludeOwners NavigationPathQueryParameters2DPathMetadataFlags = 4
  NavigationPathQueryParameters2DPathMetadataIncludeAll NavigationPathQueryParameters2DPathMetadataFlags = 7
)
