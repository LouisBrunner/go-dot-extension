// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  NavigationPathQueryParameters2DPathfindingAlgorithmPathfindingAlgorithmAstar NavigationPathQueryParameters2DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters2DPathPostProcessing int
const (
  NavigationPathQueryParameters2DPathPostProcessingPathPostprocessingCorridorfunnel NavigationPathQueryParameters2DPathPostProcessing = 0
  NavigationPathQueryParameters2DPathPostProcessingPathPostprocessingEdgecentered NavigationPathQueryParameters2DPathPostProcessing = 1
)

type NavigationPathQueryParameters2DPathMetadataFlags int
const (
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeNone NavigationPathQueryParameters2DPathMetadataFlags = 0
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeTypes NavigationPathQueryParameters2DPathMetadataFlags = 1
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeRids NavigationPathQueryParameters2DPathMetadataFlags = 2
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeOwners NavigationPathQueryParameters2DPathMetadataFlags = 4
  NavigationPathQueryParameters2DPathMetadataFlagsPathMetadataIncludeAll NavigationPathQueryParameters2DPathMetadataFlags = 7
)

func  (me *NavigationPathQueryParameters2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) GetPathfindingAlgorithm() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) GetPathPostprocessing() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) SetMap(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) GetMap() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) SetStartPosition(start_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) GetStartPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) SetTargetPosition(target_position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) GetTargetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) SetNavigationLayers(navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) GetNavigationLayers() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) SetMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters2D) GetMetadataFlags() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
