// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  NavigationPathQueryParameters3DPathfindingAlgorithmPathfindingAlgorithmAstar NavigationPathQueryParameters3DPathfindingAlgorithm = 0
)

type NavigationPathQueryParameters3DPathPostProcessing int
const (
  NavigationPathQueryParameters3DPathPostProcessingPathPostprocessingCorridorfunnel NavigationPathQueryParameters3DPathPostProcessing = 0
  NavigationPathQueryParameters3DPathPostProcessingPathPostprocessingEdgecentered NavigationPathQueryParameters3DPathPostProcessing = 1
)

type NavigationPathQueryParameters3DPathMetadataFlags int
const (
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeNone NavigationPathQueryParameters3DPathMetadataFlags = 0
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeTypes NavigationPathQueryParameters3DPathMetadataFlags = 1
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeRids NavigationPathQueryParameters3DPathMetadataFlags = 2
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeOwners NavigationPathQueryParameters3DPathMetadataFlags = 4
  NavigationPathQueryParameters3DPathMetadataFlagsPathMetadataIncludeAll NavigationPathQueryParameters3DPathMetadataFlags = 7
)

func  (me *NavigationPathQueryParameters3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) GetPathfindingAlgorithm() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) GetPathPostprocessing() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) SetMap(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) GetMap() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) SetStartPosition(start_position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) GetStartPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) SetTargetPosition(target_position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) GetTargetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) SetNavigationLayers(navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) GetNavigationLayers() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) SetMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationPathQueryParameters3D) GetMetadataFlags() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
