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



// Enums

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

func (me *NavigationPathQueryParameters2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPathQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationPathQueryParameters2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) GetPathfindingAlgorithm()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) GetPathPostprocessing()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) SetMap(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) GetMap()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) SetStartPosition(start_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) GetStartPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) SetTargetPosition(target_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) SetNavigationLayers(navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) GetNavigationLayers()  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) SetMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags, )  {
  panic("TODO: implement")
}

func  (me *NavigationPathQueryParameters2D) GetMetadataFlags()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
